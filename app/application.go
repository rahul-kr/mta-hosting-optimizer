package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var router = mux.NewRouter()

func StartApplication() {
	mapUrls()
	log.Println("about to start the application...")
	initDB()
	log.Fatal(http.ListenAndServe(":8080", router))
}

type HostsMap struct {
	Id       int64  `json:"id,omitempty" gorm:"primary_key"`
	Ip       string `json:"ip,omitempty" gorm:"type:varchar(255);not null"`
	HostName string `json:"hostname,omitempty" gorm:"type:varchar(255);not null"`
	Active   bool   `json:"active,omitempty" gorm:"type:bool;not null default:false"`
}

var db *gorm.DB

func initDB() {
	var err error
	dataSourceName := "root:@tcp(localhost:3306)/?parseTime=True"
	db, err = gorm.Open("mysql", dataSourceName)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	// Create the database. This is a one-time step.
	// Comment out if running multiple times - You may see an error otherwise
	db.Exec("CREATE DATABASE IF NOT EXISTS mta_hosting")
	db.Exec("USE mta_hosting")
	db.AutoMigrate(&HostsMap{})
}

func getMappings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var hostmapping []HostsMap
	db.Preload("HostsMap").Find(&hostmapping)
	json.NewEncoder(w).Encode(hostmapping)
}

func getThreshold(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputMapCount := params["X"]
	var hostmap []HostsMap
	db.Raw("SELECT host_name, COUNT(active) AS active_status FROM hosts_maps WHERE active=true GROUP BY host_name HAVING count(active_status)= ?;", inputMapCount).Scan(&hostmap)
	json.NewEncoder(w).Encode(hostmap)
}
