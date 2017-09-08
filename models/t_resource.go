package models

import (
	"time"
)

type T_resource struct {
	Id          int64 `gorm:"primary_key"`
	Name        string
	Url         string
	Movie_id    int64
	Create_time time.Time
	F_id        int64
}

func (_ *T_resource) UpdateRelation(m_id int64, f_id int64) {
	db := Connect()
	defer db.Close()

	db.Model(t).Where("movie_id = ?", m_id).Update("f_id", f_id)
}
