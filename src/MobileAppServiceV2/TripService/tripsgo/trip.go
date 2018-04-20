package tripsgo

import (
	"time"
)

type Trip struct {

	// Trip Id
	Id string `json:"Id,omitempty"`

	Name string `json:"Name,omitempty"`

	// User's unique identity
	UserId string `json:"UserId,omitempty"`

	RecordedTimeStamp string `json:"RecordedTimeStamp,omitempty"`

	EndTimeStamp string `json:"EndTimeStamp,omitempty"`

	Rating int32 `json:"Rating,omitempty"`

	IsComplete bool `json:"IsComplete,omitempty"`

	HasSimulatedOBDData bool `json:"HasSimulatedOBDData,omitempty"`

	AverageSpeed float32 `json:"AverageSpeed,omitempty"`

	FuelUsed float32 `json:"FuelUsed,omitempty"`

	HardStops int64 `json:"HardStops,omitempty"`

	HardAccelerations int64 `json:"HardAccelerations,omitempty"`

	Distance float32 `json:"Distance,omitempty"`

	Created time.Time `json:"Created,omitempty"`

	UpdatedAt time.Time `json:"UpdatedAt,omitempty"`

	Deleted bool `json:"Deleted,omitempty"`
}
