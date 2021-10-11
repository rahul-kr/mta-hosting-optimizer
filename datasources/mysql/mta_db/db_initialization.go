package mta_db

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var Client *gorm.DB

type HostsMap struct {
	Id       int64  `json:"id,omitempty" gorm:"primary_key"`
	Ip       string `json:"ip,omitempty" gorm:"type:varchar(255);not null"`
	HostName string `json:"hostname,omitempty" gorm:"type:varchar(255);not null"`
	Active   bool   `json:"active,omitempty" gorm:"type:bool;not null default:false"`
}

func init() {
	var err error

	envFileErr := godotenv.Load("docker/.env")
	if envFileErr != nil {
		log.Fatalf("Error loading .env file")
	}
	//These values derived from docker/.env file and Dockerfile ENV var
	dataSourceName := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":3306)/" + os.Getenv("MYSQL_DATABASE") + "?parseTime=True"
	Client, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	Client.Exec("CREATE DATABASE IF NOT EXISTS mta_hosting")
	Client.AutoMigrate(&HostsMap{})
}
