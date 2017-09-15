package douban

import (
	"github.com/gladmo/film-info/settings"
	"regexp"
)

// exec data, get douban id and transfer ScrapyById()

/**
 * 	ch   chan   int channel status
 * 	// data layout
 * 	-4 Dytt8 not match
 * 	-3 BttiantangsSpider not find
 *  -2 repeat
 *  -1 default
 *
 *	// succ
 *  0 succ
 *
 *  // system layout
 *  1 not find
 *  2 repeat to many times
 */

/**
 * default spider
 * @param ch chan int [description]
 */
func DefaultSpider(ch chan int) {
	ch <- -1
}

/**
 * scrapy www.bttiantangs.com
 * @param url   string douban url, can get douban id
 * @param tm_id int64  t_movie id
 * @param ch    chan   int   chanel
 */
func BttiantangsSpider(url string, tm_id int64, ch chan int) {

	idRe := regexp.MustCompile(`\d+$`)

	id := idRe.FindString(url)

	// if not find id
	if id == "" {
		ch <- 1
		return
	}

	api := Api{
		UseProxy:        settings.GetBool("proxy.useproxy"),
		Dbv2RepeatCount: 10,
	}

	api.ScrapyById(id, tm_id, ch)
}

/**
 * Scrapy www.dytt8.net
 * @param tm_id			int64		t_movie id
 * @param  title		string		save title
 * @param  year			string		movie year
 */
func Dytt8(tm_id int64, title, year string, ch chan int) {
	api := Api{
		UseProxy:        settings.GetBool("proxy.useproxy"),
		Dbv2RepeatCount: 10,
	}

	nameRe := regexp.MustCompile("《.*》")
	name := nameRe.FindString(title)
	if name == "" {
		// if not match 《》, set the keyword is title
		name = title
	}

	id := api.Douban_search(tm_id, name, year)

	// not match this movie
	if id == "" {
		ch <- -4
		return
	}

	api.ScrapyById(id, tm_id, ch)
}
