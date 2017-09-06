package douban

import (
// "github.com/gladmo/film-info/models"
)

func Spider() {
	Scrapy("http://movie.douban.com/subject/20451290")
	// data := new(models.T_movie).GetData()

	// for _, v := range data {
	// 	Scrapy(v.Douban)
	// }
}
