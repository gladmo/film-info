package douban

import (
	// "github.com/gladmo/film-info/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	HOST = "http://api.douban.com/v2/movie/subject/%s"
)

func Scrapy(url string) {

	idRe := regexp.MustCompile(`\d+$`)

	id := idRe.FindString(url)

	douban_api_v2(id)
}

func douban_api_v2(id string) {
	apiHost := fmt.Sprintf(HOST, id)

	client := http.Client{}

	req, err := http.NewRequest("GET", apiHost, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36")

	res, _ := client.Do(req)
	defer res.Body.Close()

	result, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(result))

	var s DoubanStruct

	json.Unmarshal(result, &s)

	fmt.Print(s)
}
