package main

import (
	"MQTTHub/internal/adapters/bot"
	"MQTTHub/internal/adapters/db"
	"MQTTHub/internal/adapters/mqtt"
	"MQTTHub/internal/adapters/serializer"
	"MQTTHub/internal/application"
	"MQTTHub/internal/core/domain"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// github.com/denisenkom/go-mssqldb
var dbConnStr = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s",
	os.Getenv("DB_SERVER"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"))

// Paho format
var mqttConnStr = fmt.Sprintf("%s://%s:%s",
	os.Getenv("MQTT_PROTOCOL"),
	os.Getenv("MQTT_SERVER"),
	os.Getenv("MQTT_PORT"))

var hubSettingsPath = os.Getenv("HUB_SETTINGS_PATH")

var tgBotToken = os.Getenv("TELEGRAM_BOT_TOKEN")
var tgTarget, _ = strconv.Atoi(os.Getenv("TELEGRAM_BOT_TARGET"))

func main() {

	dbAdapter, err := db.NewAdapter(dbConnStr)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	jsonAdapter := serializer.NewAdapter()
	TgBotAdapter, err := bot.NewAdapter(tgBotToken, int64(tgTarget))
	if err != nil {
		log.Fatalf("failed to initiate telegram bot: %v", err)
	}
	domainLogic := domain.New()

	appApi := application.NewApplication(dbAdapter, jsonAdapter, TgBotAdapter, domainLogic)

	mqttAdapter := mqtt.NewAdapter(appApi)
	hubSettings, err := ioutil.ReadFile(hubSettingsPath)
	if err != nil {
		log.Fatalf("failed to open mqtt hub settings: %v", err)
	}
	mqttAdapter.Run(hubSettings, mqttConnStr)
}
