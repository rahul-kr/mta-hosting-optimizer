package mta_db

import (
	"fmt"

	"github.com/jinzhu/gorm"
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
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	Client, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	Client.Exec("CREATE DATABASE IF NOT EXISTS mta_hosting")
	Client.Exec("USE mta_hosting")
	Client.AutoMigrate(&HostsMap{})
}
