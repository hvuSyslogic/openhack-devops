package tripsgo

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

func SelectAllTrips() string {
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

func SelectAllTripsForUser(userID string) string {
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
