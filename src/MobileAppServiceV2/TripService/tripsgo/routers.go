package tripsgo

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/api/",
		Index,
	},

	Route{
		"CreateTrip",
		strings.ToUpper("Post"),
		"/api/trips",
		createTrip,
	},

	Route{
		"CreateTripPoint",
		strings.ToUpper("Post"),
		"/api/trips/{tripID}/trippoints",
		createTripPoint,
	},

	Route{
		"DeleteTrip",
		strings.ToUpper("Delete"),
		"/api/trips/{tripID}",
		deleteTrip,
	},

	Route{
		"DeleteTripPoint",
		strings.ToUpper("Delete"),
		"/api/trips/{tripID}/trippoints/{tripPointID}",
		deleteTripPoint,
	},

	Route{
		"GetAllTrips",
		strings.ToUpper("Get"),
		"/api/trips",
		getAllTrips,
	},

	Route{
		"GetTripById",
		strings.ToUpper("Get"),
		"/api/trips/{tripID}",
		getTripByID,
	},

	Route{
		"GetTripPointByID",
		strings.ToUpper("Get"),
		"/api/trips/{tripID}/trippoints/{tripPointID}",
		getTripPointByID,
	},

	Route{
		"GetTripPoints",
		strings.ToUpper("Get"),
		"/api/trips/{tripID}/trippoints",
		getTripPoints,
	},

	Route{
		"HealthcheckGet",
		strings.ToUpper("Get"),
		"/api/healthcheck",
		healthcheckGet,
	},

	Route{
		"UpdateTrip",
		strings.ToUpper("Patch"),
		"/api/trips/{tripID}",
		updateTrip,
	},

	Route{
		"UpdateTripPoint",
		strings.ToUpper("Patch"),
		"/api/trips/{tripID}/trippoints/{tripPointID}",
		updateTripPoint,
	},
}
