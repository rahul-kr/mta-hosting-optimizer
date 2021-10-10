package mtahosting

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rahul-kr/mta-hosting-optimizer/datasources/mysql/mta_db"
)

func GetMappings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var hostmapping []mta_db.HostsMap
	mta_db.Client.Preload("HostsMap").Find(&hostmapping)
	json.NewEncoder(w).Encode(hostmapping)
}

func GetThreshold(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	inputMapCount := params["X"]
	var hostmap []mta_db.HostsMap
	mta_db.Client.Raw("SELECT host_name, COUNT(active) AS active_status FROM hosts_maps WHERE active=true GROUP BY host_name HAVING count(active_status)= ?;", inputMapCount).Scan(&hostmap)
	json.NewEncoder(w).Encode(hostmap)
}
