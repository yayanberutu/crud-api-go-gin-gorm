package main

import (
	"crud-api/Config"
	"crud-api/Models"
	"crud-api/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Admin{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}