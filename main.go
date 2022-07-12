package main

import (
	"Day4_5/Config"
	"Day4_5/Models"
	"Day4_5/Routes"
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

	Config.DB.AutoMigrate(&Models.Product{}, &Models.Order{})
	r := Routes.SetupRouter()
	//running
	r.Run()

}
