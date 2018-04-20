package tripsgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TripPoint Service Methods

func getTripPoints(w http.ResponseWriter, r *http.Request) {
	query := "SELECT [Id], [TripId], [Latitude], [Longitude], [Speed], [RecordedTimeStamp], [Sequence], [RPM], [ShortTermFuelBank], [LongTermFuelBank], [ThrottlePosition], [RelativeThrottlePosition], [Runtime], [DistanceWithMalfunctionLight], [EngineLoad], [EngineFuelRate], [VIN] FROM [dbo].[TripPoints] WHERE Deleted = 0"

	statement, err := ExecuteQuery(query)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while retrieving trip points from database"))
		return
	}

	got := []TripPoint{}

	for statement.Next() {
		var r TripPoint
		err := statement.Scan(&r.Id, &r.TripId, &r.Latitude, &r.Longitude, &r.Speed, &r.RecordedTimeStamp, &r.Sequence, &r.RPM, &r.ShortTermFuelBank, &r.LongTermFuelBank, &r.ThrottlePosition, &r.RelativeThrottlePosition, &r.Runtime, &r.DistanceWithMalfunctionLight, &r.EngineLoad, &r.EngineFuelRate, &r.VIN)

		if err != nil {
			fmt.Fprintf(w, SerializeError(err, "Error scanning Trip Points"))
			return
		}

		got = append(got, r)
	}

	serializedReturn, _ := json.Marshal(got)

	fmt.Fprintf(w, string(serializedReturn))
}


func GetTripPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripPointId := params["id"]


	query := "SELECT [Id], [TripId], [Latitude], [Longitude], [Speed], [RecordedTimeStamp], [Sequence], [RPM], [ShortTermFuelBank], [LongTermFuelBank], [ThrottlePosition], [RelativeThrottlePosition], [Runtime], [DistanceWithMalfunctionLight], [EngineLoad], [EngineFuelRate], [VIN] FROM TripPoints WHERE Id = '" + tripPointId + "' AND Deleted = 0"

	row, err := FirstOrDefault(query)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while retrieving trip point from database"))
		return
	}

	var tripPoint TripPoint

	err = row.Scan(&tripPoint.Id, &tripPoint.TripId, &tripPoint.Latitude, &tripPoint.Longitude, &tripPoint.Speed, &tripPoint.RecordedTimeStamp, &tripPoint.Sequence, &tripPoint.RPM, &tripPoint.ShortTermFuelBank, &tripPoint.LongTermFuelBank, &tripPoint.ThrottlePosition, &tripPoint.RelativeThrottlePosition, &tripPoint.Runtime, &tripPoint.DistanceWithMalfunctionLight, &tripPoint.EngineLoad, &tripPoint.EngineFuelRate, &tripPoint.VIN)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Failed to scan a trip point"))
		return
	}

	serializedTripPoint, _ := json.Marshal(tripPoint)

	fmt.Fprintf(w, string(serializedTripPoint))
}

func createTripPoint(w http.ResponseWriter, r *http.Request) {
	tripId := r.FormValue("tripId")

	body, err := ioutil.ReadAll(r.Body)

	var tripPoint TripPoint

	err = json.Unmarshal(body, &tripPoint)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while decoding json"))
		return
	}

	tripPoint.TripId = tripId

	insertQuery := fmt.Sprintf("DECLARE @tempReturn TABLE (TripPointId NVARCHAR(128)); INSERT INTO TripPoints ([TripId], [Latitude], [Longitude], [Speed], [RecordedTimeStamp], [Sequence], [RPM], [ShortTermFuelBank], [LongTermFuelBank], [ThrottlePosition], [RelativeThrottlePosition], [Runtime], [DistanceWithMalfunctionLight], [EngineLoad], [EngineFuelRate], [MassFlowRate], [HasOBDData], [HasSimulatedOBDData], [VIN], [Deleted]) OUTPUT Inserted.ID INTO @tempReturn VALUES ('%s', '%s', '%s', '%s', '%s', %d, '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s', 'false'); SELECT TripPointId FROM @tempReturn",
		tripPoint.TripId,
		tripPoint.Latitude,
		tripPoint.Longitude,
		tripPoint.Speed,
		tripPoint.RecordedTimeStamp,
		tripPoint.Sequence,
		tripPoint.RPM,
		tripPoint.ShortTermFuelBank,
		tripPoint.LongTermFuelBank,
		tripPoint.ThrottlePosition,
		tripPoint.RelativeThrottlePosition,
		tripPoint.Runtime,
		tripPoint.DistanceWithMalfunctionLight,
		tripPoint.EngineLoad,
		tripPoint.MassFlowRate,
		tripPoint.EngineFuelRate,
		strconv.FormatBool(tripPoint.HasOBDData),
		strconv.FormatBool(tripPoint.HasSimulatedOBDData),
		tripPoint.VIN)

	fmt.Fprintf(w, insertQuery)

	// var newTripPoint NewTripPoint

	// result, err := ExecuteQuery(insertQuery)

	// if err != nil {
	// 	fmt.Fprintf(w, SerializeError(err, "Error while inserting Trip Point onto database"))
	// 	return
	// }

	// for result.Next() {
	// 	err = result.Scan(&newTripPoint.Id)

	// 	if err != nil {
	// 		fmt.Fprintf(w, SerializeError(err, "Error while retrieving last id"))
	// 	}
	// }

	// serializedTripPoint, _ := json.Marshal(newTripPoint)

	// fmt.Fprintf(w, string(serializedTripPoint))
}

func updateTripPoint(w http.ResponseWriter, r *http.Request) {
	tripPointId := r.FormValue("id")

	body, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while reading request body"))
		return
	}

	var tripPoint TripPoint

	err = json.Unmarshal(body, &tripPoint)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while decoding json"))
		return
	}

	updateQuery := fmt.Sprintf("UPDATE [TripPoints] SET [TripId] = '%s',[Latitude] = '%s',[Longitude] = '%s',[Speed] = '%s',[RecordedTimeStamp] = '%s',[Sequence] = %d,[RPM] = '%s',[ShortTermFuelBank] = '%s',[LongTermFuelBank] = '%s',[ThrottlePosition] = '%s',[RelativeThrottlePosition] = '%s',[Runtime] = '%s',[DistanceWithMalfunctionLight] = '%s',[EngineLoad] = '%s',[MassFlowRate] = '%s',[EngineFuelRate] = '%s',[HasOBDData] = '%s',[HasSimulatedOBDData] = '%s',[VIN] = '%s' WHERE Id = '%s'",
		tripPoint.TripId,
		tripPoint.TripId,
		tripPoint.Latitude,
		tripPoint.Longitude,
		tripPoint.Speed,
		tripPoint.RecordedTimeStamp,
		tripPoint.Sequence,
		tripPoint.RPM,
		tripPoint.ShortTermFuelBank,
		tripPoint.LongTermFuelBank,
		tripPoint.ThrottlePosition,
		tripPoint.RelativeThrottlePosition,
		tripPoint.Runtime,
		tripPoint.DistanceWithMalfunctionLight,
		tripPoint.EngineLoad,
		tripPoint.MassFlowRate,
		tripPoint.EngineFuelRate,
		strconv.FormatBool(tripPoint.HasOBDData),
		strconv.FormatBool(tripPoint.HasSimulatedOBDData),
		tripPoint.VIN,
		tripPointId)

	result, err := ExecuteNonQuery(updateQuery)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while patching Trip Point on the database"))
		return
	}

	fmt.Fprintf(w, string(result))
}

func deleteTripPoint(w http.ResponseWriter, r *http.Request) {
	tripPointId := r.FormValue("id")

	deleteTripPointQuery := fmt.Sprintf("UPDATE TripPoints SET Deleted = 1 WHERE Id = '%s'", tripPointId)

	result, err := ExecuteNonQuery(deleteTripPointQuery)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while deleting trip point from database"))
		return
	}

	serializedResult, _ := json.Marshal(result)

	fmt.Fprintf(w, string(serializedResult))
}

func getMaxSequence(w http.ResponseWriter, r *http.Request) {
	tripId := r.FormValue("id")

	query := fmt.Sprintf("SELECT MAX(Sequence) as MaxSequence FROM TripPoints where tripid = '%s'", tripId)

	row, err := FirstOrDefault(query)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while querying Max Sequence"))
		return
	}

	var MaxSequence string

	err = row.Scan(&MaxSequence)

	if err != nil {
		fmt.Fprintf(w, SerializeError(err, "Error while obtaining max sequence"))
		return
	}

	fmt.Fprintf(w, MaxSequence)
}

// End of Trip Point Service Methods

type newTripPoint struct {
	Id string
}
