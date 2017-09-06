package models

type Film struct {
	F_id int64 `gorm:"primary_key"`
}

func (f *Film) Save() (res bool) {
	db := Connect()
	defer db.Close()
	res = db.NewRecord(f)
	db.Create(f)
	return
}
