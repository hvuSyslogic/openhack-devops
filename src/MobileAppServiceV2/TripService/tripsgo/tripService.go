package tripsgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/denisenkom/go-mssqldb" //vscode deletes this import if it is not a blank import
	"github.com/gorilla/mux"
)

// Trip Service Methods

func getTripByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//Build Query
	var query = SelectTripByIdQuery(params["id"])

	//Execute Query
	row, err := FirstOrDefault(query)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "getTripsByID - Error while retrieving trip from database"))
		return
	}

	var trip Trip

	err = row.Scan(
		&trip.Id,
		&trip.Name,
		&trip.UserId,
		&trip.RecordedTimeStamp,
		&trip.EndTimeStamp,
		&trip.Rating,
		&trip.IsComplete,
		&trip.HasSimulatedOBDData,
		&trip.AverageSpeed,
		&trip.FuelUsed,
		&trip.HardStops,
		&trip.HardAccelerations,
		&trip.Distance)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Failed to scan a trip"))
		return
	}

	serializedTrip, _ := json.Marshal(trip)

	fmt.Fprintf(w, string(serializedTrip))
}

func getAllTrips(w http.ResponseWriter, r *http.Request) {

	var query = SelectAllTripsQuery()

	tripRows, err := ExecuteQuery(query)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "getAllTrips - Query Failed to Execute."))
		return
	}

	trips := []Trip{}

	for tripRows.Next() {
		var r Trip
		err := tripRows.Scan(
			&r.Id,
			&r.Name,
			&r.UserId,
			&r.RecordedTimeStamp,
			&r.EndTimeStamp,
			&r.Rating,
			&r.IsComplete,
			&r.HasSimulatedOBDData,
			&r.AverageSpeed,
			&r.FuelUsed,
			&r.HardStops,
			&r.HardAccelerations,
			&r.Distance)

		if err != nil {
			fmt.Fprintf(w, SerializeError(err, "GetAllTrips - Error scanning Trips"))
			return
		}

		trips = append(trips, r)
	}

	tripsJSON, _ := json.Marshal(trips)

	fmt.Fprintf(w, string(tripsJSON))
}

func getAllTripsForUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var query = SelectAllTripsForUserQuery(params["id"])

	tripRows, err := ExecuteQuery(query)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "getAllTripsForUser - Error while retrieving trips from database"))
		return
	}

	trips := []Trip{}

	for tripRows.Next() {
		var r Trip
		err := tripRows.Scan(&r.Id,
			&r.Name,
			&r.UserId,
			&r.RecordedTimeStamp,
			&r.EndTimeStamp,
			&r.Rating,
			&r.IsComplete,
			&r.HasSimulatedOBDData,
			&r.AverageSpeed,
			&r.FuelUsed,
			&r.HardStops,
			&r.HardAccelerations,
			&r.Distance)

		if err != nil {
			fmt.Fprintf(w, SerializeError(err, "getAllTripsForUser - Error scanning Trips"))
			return
		}

		trips = append(trips, r)
	}

	tripsJSON, _ := json.Marshal(trips)

	fmt.Fprintf(w, string(tripsJSON))
}

func deleteTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var deleteTripPointsQuery = DeleteTripPointsForTripQuery(params["id"])
	var deleteTripsQuery = DeleteTripQuery(params["id"])

	result, err := ExecuteNonQuery(deleteTripPointsQuery)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while deleting trip points from database"))
		return
	}

	log.Println(fmt.Sprintln("Deleted trip points for Trip '%s'", params["id"]))

	result, err = ExecuteNonQuery(deleteTripsQuery)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while deleting trip from database"))
		return
	}

	log.Println(fmt.Sprintln("Deleted trip '%s'", params["id"]))

	serializedResult, _ := json.Marshal(result)

	fmt.Fprintf(w, string(serializedResult))
}

func updateTrip(w http.ResponseWriter, r *http.Request) {
	tripID := r.FormValue("id")
	body, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while reading request body"))
		return
	}

	var trip Trip

	err = json.Unmarshal(body, &trip)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while decoding json"))
		return
	}

	updateQuery := fmt.Sprintf("UPDATE Trips SET Name = '%s', UserId = '%s', RecordedTimeStamp = '%s', EndTimeStamp = '%s', Rating = %d, IsComplete = '%s', HasSimulatedOBDData = '%s', AverageSpeed = %f, FuelUsed = %s, HardStops = %s, HardAccelerations = %s, MainPhotoUrl = '%s', Distance = %f, UpdatedAt = GETDATE() WHERE Id = '%s'", trip.Name, trip.UserId, trip.RecordedTimeStamp, trip.EndTimeStamp, trip.Rating, strconv.FormatBool(trip.IsComplete), strconv.FormatBool(trip.HasSimulatedOBDData), trip.AverageSpeed, trip.FuelUsed, trip.HardStops, trip.HardAccelerations, trip.Distance, tripID)

	result, err := ExecuteNonQuery(updateQuery)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while patching trip on the database"))
		return
	}

	fmt.Fprintf(w, string(result))
}

func createTrip(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("userId")

	body, err := ioutil.ReadAll(r.Body)

	var trip Trip

	err = json.Unmarshal(body, &trip)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while decoding json"))
		return
	}

	trip.UserId = userId

	insertQuery := fmt.Sprintf("DECLARE @tempReturn TABLE (TripId NVARCHAR(128)); INSERT INTO Trips (Name, UserId, RecordedTimeStamp, EndTimeStamp, Rating, IsComplete, HasSimulatedOBDData, AverageSpeed, FuelUsed, HardStops, HardAccelerations, Distance, Deleted) OUTPUT Inserted.ID INTO @tempReturn VALUES ('%s', '%s', '%s', '%s', %d, '%s', '%s', %f, '%s', '%s', '%s', '%s', %f, 'false'); SELECT TripId FROM @tempReturn", trip.Name, trip.UserId, trip.RecordedTimeStamp, trip.EndTimeStamp, trip.Rating, strconv.FormatBool(trip.IsComplete), strconv.FormatBool(trip.HasSimulatedOBDData), trip.AverageSpeed, trip.FuelUsed, trip.HardStops, trip.HardAccelerations, trip.Distance)

	var newTrip NewTrip

	result, err := ExecuteQuery(insertQuery)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while inserting trip onto database"))
		return
	}

	for result.Next() {
		err = result.Scan(&newTrip.Id)

		if err != nil {
			fmt.Fprintf(w, SerializeError(err, "Error while retrieving last id"))
		}
	}

	serializedTrip, _ := json.Marshal(newTrip)

	fmt.Fprintf(w, string(serializedTrip))
}

type NewTrip struct {
	Id string
}

// End of Trip Service Methods
