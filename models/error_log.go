package models

import (
	"time"
)

type Error_log struct {
	E_id     int64 `gorm:"primary_key"`
	Tm_id    int64
	Msg      string `gorm:"size:255"`
	CreateAt time.Time
}

func (e *Error_log) Save() (res bool) {
	db := Connect()
	defer db.Close()
	res = db.NewRecord(e)
	db.Create(e)
	return
}
