package models

import "time"

type Measurement struct {
	ID               int
	MeasurementValue float32
	Unit             string
	MeasurementDate  time.Time
	IDSensor         int
}
