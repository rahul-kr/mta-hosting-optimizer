package app

func mapUrls() {
	router.HandleFunc("/threshold/{X}", getThreshold).Methods("GET")
	// Read-all
	router.HandleFunc("/mappings", getMappings).Methods("GET")

}
