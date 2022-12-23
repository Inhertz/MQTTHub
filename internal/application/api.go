package application

import (
	"MQTTHub/internal/core/domain"
	"MQTTHub/internal/core/models"
	"fmt"
	"log"
)

type Application struct {
	db         DbPort
	serializer SerializerPort
	bot        BotPort
	logic      *domain.Logic //struct, not an interface!
}

// NewApplication creates a new Application
func NewApplication(db DbPort, se SerializerPort, bot BotPort, logic *domain.Logic) *Application {
	return &Application{db: db, serializer: se, bot: bot, logic: logic}
}

// CreateMeasurement is a use case, it takes topic and payload and inserts into db
func (api Application) CreateMeasurement(topic string, payload []byte) error {
	plainTree, err := api.logic.SplitTopicTree(topic)
	macAdd := plainTree[len(plainTree)-1]
	if err != nil {
		api.pushError(topic)
		return err
	}
	var s models.Sensor
	err = api.db.Find("device_add", macAdd, &s)
	if err != nil {
		api.pushError(topic)
		return err
	}
	var r models.Measurement
	err = api.serializer.Decode(payload, &r)
	if err != nil {
		api.pushError(topic)
		return err
	}
	answer, err := api.logic.RegComplete(s, r)
	if err != nil {
		api.pushError(topic)
		return err
	}
	err = api.db.Create(&answer)
	if err != nil {
		api.pushError(topic)
		return err
	}

	return nil
}

func (api Application) pushError(topic string) {

	err := api.bot.PushMessage(fmt.Sprintf("Se acaba de producir un error en el Topic: %s. \n Para más información revise los registros del Recolector", topic))
	if err != nil {
		log.Print(err)
	}
}
