package models

type Film struct {
	F_id           int64  `gorm:"primary_key"`
	Id             string `gorm:"size:100;unique_index"`
	Title          string `gorm:"size:255"`
	Original_title string `gorm:"size:255"`
	Aka            string `gorm:"size:255"`
	Alt            string `gorm:"size:255"`
	Mobile_url     string `gorm:"size:255"`

	Rating     float64
	Rating_ext string `gorm:"size:255"`

	Ratings_count int64
	Wish_count    int64
	Collect_count int64
	Do_count      int64

	Images     string `gorm:"size:255"`
	Images_ext string `gorm:"size:255"`

	Subtype string `gorm:"size:255"`

	Casts string `gorm:"size:255"`

	Year           string `gorm:"size:255"`
	Genres         string `gorm:"size:255"`
	Countries      string `gorm:"size:255"`
	Summary        string `gorm:"type:text"`
	Comments_count int64
	Reviews_count  int64
	Seasons_count  int64
	Current_season int64
	Episodes_count int64
}

func (f *Film) Save() (res bool) {
	db := Connect()
	defer db.Close()
	res = db.NewRecord(f)
	db.Create(f)
	return
}

// id is douban api's id not f_id
func (f *Film) FindById(id string) (f_id int64, ok bool) {
	db := Connect()
	defer db.Close()

	db.Select("id,f_id").Where("`id` = ?", id).First(&f)
	if f.Id == "" {
		return 0, false
	}
	return f.F_id, true

}
