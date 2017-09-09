package douban

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gladmo/film-info/models"
	"github.com/gladmo/film-info/proxy"
	// "github.com/gladmo/film-info/tools"
)

const (
	HOST = "http://api.douban.com/v2/movie/subject/%s"
)

/**
 * scrapy by douban id
 * @param  id  string douban id
 * @param  tm_id int64  t_movie id
 * @param  ch   chan   int
 */
func (api *Api) ScrapyById(id string, tm_id int64, ch chan int) {

	// check repeat
	if new(models.Film).FindById(id) {
		ch <- -2
	}

	// get film info by douban id
	info, err_code := api.douban_api_v2(id)

	// 0 success, 2 not find, 3 other err
	if err_code != 0 {
		err := models.Error_log{
			CreateAt: time.Now(),
			Tm_id:    tm_id,
		}
		switch err_code {
		case 2:
			err.Msg = "movie_not_found"
		case 3:
			err.Msg = "repeat_many_times"
		}

		new(models.T_movie).CompleteById(tm_id, 0, int64(err_code))
		err.Save()

		ch <- 0
		return
	}

	// arr to json
	Rating_ext, _ := json.Marshal(info.Rating)
	Images_ext, _ := json.Marshal(info.Images)
	Casts, _ := json.Marshal(info.Casts)

	film := &models.Film{
		Id:             info.Id,
		Title:          info.Title,
		Original_title: info.Original_title,
		Aka:            strings.Join(info.Aka, ","),
		Alt:            info.Alt,
		Mobile_url:     info.Mobile_url,

		Rating:     info.Rating.Average, // change
		Rating_ext: string(Rating_ext),  // change

		Ratings_count: info.Ratings_count,
		Wish_count:    info.Wish_count,
		Collect_count: info.Collect_count,
		Do_count:      info.Do_count,

		Images:     info.Images.Large,  // change
		Images_ext: string(Images_ext), // change

		Subtype: info.Subtype,

		Casts: string(Casts), // change

		Year: info.Year,

		Genres:    strings.Join(info.Genres, ","),    // change
		Countries: strings.Join(info.Countries, ","), // change

		Summary:        info.Summary,
		Comments_count: info.Comments_count,
		Reviews_count:  info.Reviews_count,
		Seasons_count:  info.Seasons_count,
		Current_season: info.Current_season,
		Episodes_count: info.Episodes_count,
	}

	// save film info
	film.Save()

	// update t_movie status to 1
	new(models.T_movie).CompleteById(tm_id, film.F_id, 1)

	// when succ, update t_resource
	new(models.T_resource).UpdateRelation(tm_id, film.F_id)

	ch <- 1
}

type Api struct {
	UseProxy        bool
	Dbv2RepeatCount int
}

var api = Api{
	UseProxy:        false,
	Dbv2RepeatCount: 10,
}

var RepeatMap = make(map[string]int)

/**
 * [douban_api_v2 description]
 * @param  {[type]} id string)       (s DoubanStruct, err_code int [description]
 * @return filmInfo    DoubanStruct
 * @return err_code    int				0 success, 2 not find, 3 other err.
 */
func (api *Api) douban_api_v2(id string) (filmInfo DoubanStruct, err_code int) {

	err_code = 0

	apiHost := fmt.Sprintf(HOST, id)

	client := http.Client{
		Timeout: time.Second * 8, // set timeout
	}

	_proxy := proxy.Proxy{}
	if api.UseProxy {
		// get proxy ip
		urli := url.URL{}
		_proxy := _proxy.GetProxy()
		urlproxy, _ := urli.Parse("http://" + _proxy.Ip)

		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(urlproxy), // set proxy
		}
	}

	req, err := http.NewRequest("GET", apiHost, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("timeout...")

		if api.UseProxy {
			// timeout delete proxy
			_proxy.DeleteOne(_proxy.Ip)
		}

		return api.douban_api_v2(id)
	}
	defer res.Body.Close()

	result, _ := ioutil.ReadAll(res.Body)

	var r ErrorRes
	json.Unmarshal(result, &r)

	// rate limit
	if r.Code == 112 {
		fmt.Println("repeat...", r.Code)
		return api.douban_api_v2(id)
	}

	if r.Msg == "movie_not_found" {
		err_code = 2
		return
	}

	json.Unmarshal(result, &filmInfo)

	// if not parse result, repeat it
	if filmInfo.Id == "" {
		if RepeatMap[id] >= api.Dbv2RepeatCount {
			err_code = 3
			return
		}
		RepeatMap[id]++
		return api.douban_api_v2(id)
	}

	if _, ok := RepeatMap[id]; ok {
		// key not set can delete too, no err
		delete(RepeatMap, id)
	}

	return
}
