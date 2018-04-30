package tripsgo

import "fmt"

func SelectTripByIdQuery(tripID string) string {
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

func DeleteTripPointsForTripQuery(tripID string) string {
	return fmt.Sprintf("UPDATE TripPoints SET Deleted = 1 WHERE TripId = '%s'", tripID)
}

func DeleteTripQuery(tripID string) string {
	return fmt.Sprintf("UPDAte Trips SET Deleted = 1 WHERE Id = '%s'", tripID)
}
