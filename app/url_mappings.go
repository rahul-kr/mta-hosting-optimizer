package app

func mapUrls() {
	// filter result based on query parameters
	router.HandleFunc("/threshold/{X}", getThreshold).Methods("GET")
	// Read-all
	router.HandleFunc("/mappings", getMappings).Methods("GET")

}
