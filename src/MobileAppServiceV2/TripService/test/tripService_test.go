package tripsgo

import (
	"log"
	"net/http"
	"testing"

	tripSvc "github.com/Azure-Samples/openhack-devops/src/MobileAppServiceV2/TripService/tripsgo"
)

var healthcheckStatus = `{"message": "Trip Service Healthcheck","status": "Healthy"}`

var testTripUpdateOK = `{
	"Id: "",
	"Name":"Trip CREATE TEST",
	"UserId":"GO_TEST",
    "RecordedTimeStamp": "2018-04-19T19:08:16.03Z",
    "EndTimeStamp": "2018-04-19T19:42:49.573Z",
	"Rating":95,
	"IsComplete":false,
	"HasSimulatedOBDData":true,
	"AverageSpeed":100,
	"FuelUsed":10.27193484,
	"HardStops":2,
	"HardAccelerations":4,
	"Distance":30.0275486,
	"Created":"2018-01-01T12:00:00Z",
	"UpdatedAt":"2001-01-01T12:00:00Z"
}`

var testTripUpdateINVALID = `{
	"Id: "",
	"Name":"",
	"UserId":"GO_TEST",
    "RecordedTimeStamp": "",
    "EndTimeStamp": "",
	"Rating": ,
	"IsComplete":false,
	"HasSimulatedOBDData":true,
	"AverageSpeed":0,
	"FuelUsed":0.0,
	"HardStops":0,
	"HardAccelerations":0,
	"Distance":0.0,
	"Created":"",
	"UpdatedAt":""
}`

var testTripCreate1 = `{
	"Name":"Trip CREATE TEST",
	"UserId":"GO_TEST",
    "RecordedTimeStamp": "2018-04-19T19:08:16.03Z",
    "EndTimeStamp": "2018-04-19T19:42:49.573Z",
	"Rating":95,
	"IsComplete":false,
	"HasSimulatedOBDData":true,
	"AverageSpeed":100,
	"FuelUsed":10.27193484,
	"HardStops":2,
	"HardAccelerations":4,
	"Distance":30.0275486,
	"Created":"2018-01-01T12:00:00Z",
	"UpdatedAt":"2001-01-01T12:00:00Z"
}`

func TestTrip(t *testing.T) {
	router := tripSvc.NewRouter()

	// notFoundError := `{"error_code":"NOT_FOUND", "message":"NOT_FOUND"}`
	// nameRequiredError := `{"error_code":"INVALID_DATA","message":"INVALID_DATA","details":[{"field":"name","error":"cannot be blank"}]}`

	var apiTest1 = [5]apiTestCase{
		{"t0 - healthcheck", "GET", "/api/healthcheck/trips", "", http.StatusOK, healthcheckStatus, ""},
		{"t1 - get all trips", "GET", "/api/trips", "", http.StatusOK, "", ""},
		{"t2 - get a nonexistent trip", "GET", "/api/trips/99999", "", http.StatusNotFound, "", ""},
		{"t3 - create a trip", "POST", "/api/trips", testTripCreate1, http.StatusOK, "", ""},
		{"t4 - update a trip", "PATCH", "/api/trips/", testTripUpdateOK, http.StatusOK, testTripUpdateOK, ""},
	}

	runAPITests(t, router, apiTest1[0:4])

	log.Print(apiTest1[3].actualResponse)

}
