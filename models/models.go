package models

import (
	"github.com/gladmo/film-info/settings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {

	username := settings.Get("mysql.username")
	pass := settings.Get("mysql.password")
	host := settings.Get("mysql.host")
	port := settings.Get("mysql.port")
	schema := settings.Get("mysql.schema")
	printsql := settings.GetBool("mysql.printsql")

	dsn := username + ":" + pass + "@tcp(" + host + ":" + port + ")/" + schema + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.LogMode(printsql)
	return db
}

func InitTables() {
	db := Connect()
	db.AutoMigrate(&Film{})
	db.AutoMigrate(&Error_log{})

	defer db.Close()
}
