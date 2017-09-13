package models

import (
	"time"
)

func (T_movie) TableName() string {
	return "t_movie"
}

type T_movie struct {
	Id          int64 `gorm:"primary_key"`
	Name        string
	Douban      string
	Year        string
	Url         string
	Source      string
	F_id        int64
	Status      int64
	Create_time time.Time
}

func (t *T_movie) GetData(limit int) (m []T_movie) {
	db := Connect()
	defer db.Close()

	db.Limit(limit).Select("*").Where("status = 0").Find(&m)

	return
}

/**
 * Complete t_movie
 * status :
 * # Go heavy
 * -1 aralay scrapy
 * -99 not match, Need human intervention
 * # default
 * 0 not scrapy
 * # succ
 * 1 success
 * # system layout
 * 2 not find
 * 3 repeat many times
 *
 * @param id int64
 * @param f_id int64
 * @param status int64
 */
func (t *T_movie) CompleteById(id int64, f_id int64, status int64) {
	db := Connect()
	defer db.Close()

	db.Model(t).Where("id = ?", id).Update(T_movie{F_id: f_id, Status: status})
}
