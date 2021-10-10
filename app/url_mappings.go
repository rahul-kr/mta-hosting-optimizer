package app

import (
	"github.com/rahul-kr/mta-hosting-optimizer/controllers/mtahosting"
	"github.com/rahul-kr/mta-hosting-optimizer/controllers/ping"
)

func mapUrls() {
	// filter result based on query parameters
	router.HandleFunc("/ping", ping.Ping).Methods("GET")
	// filter result based on query parameters
	router.HandleFunc("/threshold/{X}", mtahosting.GetThreshold).Methods("GET")
	// Read-all
	router.HandleFunc("/mappings", mtahosting.GetMappings).Methods("GET")

}
