package application

import (
	"MQTTHub/internal/core/domain"
	"MQTTHub/internal/core/models"
)

type Application struct {
	db         DbPort
	serializer SerializerPort
	logic      *domain.Logic //struct, not an interface!
}

// NewApplication creates a new Application
func NewApplication(db DbPort, se SerializerPort, logic *domain.Logic) *Application {
	return &Application{db: db, serializer: se, logic: logic}
}

// CreateMeasurement is a use case, it takes topic and payload and inserts into db
func (api Application) CreateMeasurement(topic string, payload []byte) error {
	plainTree, err := api.logic.SplitTopicTree(topic)
	macAdd := plainTree[len(plainTree)-1]
	if err != nil {
		return err
	}
	var s models.Sensor
	err = api.db.Find("device_add", macAdd, &s)
	if err != nil {
		return err
	}
	var r models.Measurement
	err = api.serializer.Decode(payload, &r)
	if err != nil {
		return err
	}
	answer, err := api.logic.RegComplete(s, r)
	if err != nil {
		return err
	}
	err = api.db.Create(&answer)
	if err != nil {
		return err
	}

	return nil
}
