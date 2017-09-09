package douban

import (
	"fmt"
	"github.com/gladmo/film-info/models"
)

const (
	BTTIANTANGS = "www.bttiantangs.com"
)

var Thread = 20

var ThreadCount = 0

var ChanBuff = 5

func Run() {

	tMovie := new(models.T_movie)

	for {
		data := tMovie.GetData(Thread)

		fmt.Println("Get data count: ", len(data))
		if len(data) == 0 {
			fmt.Println("Scrapy end!")
			return
		}

		ch := make(chan int, ChanBuff)
		ThreadCount += len(data)

		for _, v := range data {
			// Scrapy by source
			switch v.Source {
			case BTTIANTANGS: // www.bttiantangs.com
				go BttiantangsSpider(v.Douban, v.Id, ch)
			default:
				go DefaultSpider(ch)
			}
		}

		for i := 0; i < len(data); i++ {
			c := <-ch
			ThreadCount--
			fmt.Println("chan over, status is: ", c, " Thread count is: ", ThreadCount)
		}

		fmt.Println("end of chan")
	}
}
