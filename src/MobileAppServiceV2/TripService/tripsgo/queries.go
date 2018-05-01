package tripsgo

import (
	"fmt"
	"strconv"
)

// SelectTripByIDQuery - REQUIRED tripID value
func SelectTripByIDQuery(tripID string) string {
	return `SELECT
		Id,
		Name,
		UserId,
		RecordedTimeStamp,
		EndTimeStamp,
		Rating,
		IsComplete,
		HasSimulatedOBDData,
		AverageSpeed,
		FuelUsed,
		HardStops,
		HardAccelerations,
		Distance
		FROM Trips
		WHERE Id = '` + tripID + `'
		AND Deleted = 0`
}

// SelectAllTripsQuery - select all trips
func SelectAllTripsQuery() string {
	return `SELECT
	Id,
	Name,
	UserId,
	RecordedTimeStamp,
	EndTimeStamp,
	Rating,
	IsComplete,
	HasSimulatedOBDData,
	AverageSpeed,
	FuelUsed,
	HardStops,
	HardAccelerations,
	Distance
	FROM Trips
	WHERE Deleted = 0`
}

// SelectAllTripsForUserQuery REQUIRED userID
func SelectAllTripsForUserQuery(userID string) string {
	return `SELECT
	Id,
	Name,
	UserId,
	RecordedTimeStamp,
	EndTimeStamp,
	Rating,
	IsComplete,
	HasSimulatedOBDData,
	AverageSpeed,
	FuelUsed,
	HardStops,
	HardAccelerations,
	Distance
	FROM Trips
	WHERE UserId LIKE '%` + userID + `'
	AND Deleted = 0`
}

// DeleteTripPointsForTripQuery - REQUIRED tripID
func DeleteTripPointsForTripQuery(tripID string) string {
	return fmt.Sprintf("UPDATE TripPoints SET Deleted = 1 WHERE TripId = '%s'", tripID)
}

// DeleteTripQuery - REQUIRED tripID
func DeleteTripQuery(tripID string) string {
	return fmt.Sprintf("UPDAte Trips SET Deleted = 1 WHERE Id = '%s'", tripID)
}

// UpdateTripQuery - REQUIRED trip object and tripID
func UpdateTripQuery(trip Trip, tripID string) string {
	var query = `UPDATE Trips SET
	Name = '%s',
	UserId = '%s',
	RecordedTimeStamp = '%s',
	EndTimeStamp = '%s',
	Rating = %d,
	IsComplete = '%s',
	HasSimulatedOBDData = '%s',
	AverageSpeed = %f,
	FuelUsed = %s,
	HardStops = %s,
	HardAccelerations = %s,
	MainPhotoUrl = '%s',
	Distance = %f,
	UpdatedAt = GETDATE()
	WHERE Id = '%s'`

	return fmt.Sprintf(
		query,
		trip.Name,
		trip.UserID,
		trip.RecordedTimeStamp,
		trip.EndTimeStamp,
		trip.Rating,
		strconv.FormatBool(trip.IsComplete),
		strconv.FormatBool(trip.HasSimulatedOBDData),
		trip.AverageSpeed,
		trip.FuelUsed,
		trip.HardStops,
		trip.HardAccelerations,
		trip.Distance, tripID)

}

func createTripQuery(trip Trip) string {
	var query = `DECLARE @tempReturn
		TABLE (TripId NVARCHAR(128));
		INSERT INTO Trips (
			Name,
			UserId,
			RecordedTimeStamp,
			EndTimeStamp,
			Rating,
			IsComplete,
			HasSimulatedOBDData,
			AverageSpeed,
			FuelUsed,
			HardStops,
			HardAccelerations,
			Distance,
			Deleted)
			OUTPUT Inserted.ID
			INTO @tempReturn
			VALUES ('%s',
				'%s',
				'%s',
				'%s',
				%d,
				'%s',
				'%s',
				%f,
				'%s',
				'%s',
				'%s',
				'%s',
				%f,
				'false');
			SELECT TripId FROM @tempReturn`
	return fmt.Sprintf(
		query,
		trip.Name,
		trip.UserID,
		trip.RecordedTimeStamp,
		trip.EndTimeStamp,
		trip.Rating,
		strconv.FormatBool(trip.IsComplete),
		strconv.FormatBool(trip.HasSimulatedOBDData),
		trip.AverageSpeed,
		trip.FuelUsed,
		trip.HardStops,
		trip.HardAccelerations,
		trip.Distance)
}
