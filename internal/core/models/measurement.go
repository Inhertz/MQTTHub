package models

import "time"

type Measurement struct {
	ID       int
	Value    float32
	Unit     string
	Date     time.Time
	IDSensor int
}
