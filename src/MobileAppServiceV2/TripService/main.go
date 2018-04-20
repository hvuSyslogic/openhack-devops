package main

import (
	"log"
	"net/http"

	sw "github.com/Azure-Samples/openhack-devops/src/MobileAppServiceV2/TripService/tripsgo"
)

func main() {
	log.Printf("Server started on port 8080")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
