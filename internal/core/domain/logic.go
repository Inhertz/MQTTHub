package domain

import (
	"MQTTHub/internal/core/models"
	"errors"
	"strings"
	"time"
)

// Logic exports the domain logic to the application layer
type Logic struct {
}

// New creates a new Logic struct
func New() *Logic {
	return &Logic{}
}

// RegComplete completes measurement model recieved without date and sensor id
func (l Logic) RegComplete(s models.Sensor,
	m models.Measurement) (models.Measurement, error) {
	if s.ID <= 0 {
		err := errors.New("sensor id must be a valid id")
		return models.Measurement{}, err
	}
	m.MeasurementDate = time.Now()
	m.IDSensor = s.ID
	return m, nil
}

//SplitTopicTree recieves an MQTT topic string and returns a slice of the topic nodes
func (l Logic) SplitTopicTree(s string) ([]string, error) {

	ans := strings.Split(s, "/")
	if len(ans) <= 1 {
		err := errors.New("topic tree should be at least 2 levels, follow mqtt doc")
		return nil, err
	}
	return ans, nil
}
