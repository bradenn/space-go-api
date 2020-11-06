package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
	"space-api/src/Config"
	"space-api/src/Models"
	"space-api/src/Router"
)

func main() {

	if os.Getenv("GIN_MODE") == "" {
		envErr := godotenv.Load()
		if envErr != nil {
			fmt.Println("Failed to load .env file.")
		}
	}

	var err interface{}
	Config.DB, err = gorm.Open("postgres", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("status: ", err)
	}
	defer Config.DB.Close()

	Config.DB.AutoMigrate(&Models.System{})

	r := Router.SetupRouter()
	r.Run()

}
