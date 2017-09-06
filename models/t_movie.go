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
	Uri         string
	Create_time time.Time
	Status      int64
}

func (t *T_movie) GetData() (m []T_movie) {
	db := Connect()
	defer db.Close()

	db.Limit(1).Select("*").Where("status = 0").Find(&m)

	return
}
