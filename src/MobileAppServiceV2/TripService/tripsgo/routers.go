package tripsgo

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// Route - object representing a route handler
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - Route handler collection
type Routes []Route

// NewRouter - Constructor
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

// Index - Default route handler for service base uri
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Trips Service")
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
		"GetAllTripsForUser",
		strings.ToUpper("Get"),
		"/api/trips/user/{userID}",
		getAllTripsForUser,
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
		"/api/healthcheck/trips",
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
