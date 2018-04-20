package main

import (
	"log"
	"net/http"

	//sw "gopkg.in/Azure-Samples/openhack-devops/src/MobileAppServiceV2/TripService/tripsgo.v1.0.0"
	sw "github.com/Azure-Samples/openhack-devops/src/MobileAppServiceV2/TripService/tripsgo"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
