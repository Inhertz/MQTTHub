package mqtt

import (
	"MQTTHub/internal/application"
	"encoding/json"
	"log"

	paho "github.com/eclipse/paho.mqtt.golang"
)

type hub struct {
	client   paho.Client
	settings hubSettings
	handler  paho.MessageHandler
}

type hubSettings struct {
	ClientIdentifier string `json:"clientIdentifier"`
	Topic            string `json:"topic"`
	QualityOfService byte   `json:"qualityOfService"`
}

// const (
// 	server   = "127.0.0.1"
// 	port     = 1883
// 	protocol = "tcp"
// )

// func GetBroker() string {
// 	return fmt.Sprintf("%s://%s:%d",
// 		protocol, server, port)
// }

func newHub(settings []byte, connStr string) *hub {
	var hubSettings hubSettings
	if err := json.Unmarshal(settings, &hubSettings); err != nil {
		log.Fatalf("fatal error on marshalling hub settings: %v", err)
	}
	opts := paho.NewClientOptions()
	opts.AddBroker(connStr)
	opts.SetClientID(hubSettings.ClientIdentifier)
	client := paho.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return &hub{
		client:   client,
		settings: hubSettings,
	}
}

func (h hub) Collect(api application.APIPort) {
	h.handler = func(client paho.Client, msg paho.Message) {
		err := api.CreateMeasurement(msg.Topic(), msg.Payload())
		if err != nil {
			log.Print(err)
		}
	}
	if token := h.client.Subscribe(h.settings.Topic,
		h.settings.QualityOfService,
		h.handler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
}
