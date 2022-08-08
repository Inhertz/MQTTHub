package mqtt

import (
	"MQTTHub/internal/application"
	"fmt"
	"log"
	"time"
)

type Adapter struct {
	api application.APIPort
}

// NewAdapter creates a new Adapter
func NewAdapter(api application.APIPort) *Adapter {
	return &Adapter{api: api}
}

// Run starts the MQTT client collector
func (mqtta Adapter) Run(settings []byte, connStr string) {
	h := newHub(settings, connStr)
	h.Collect(mqtta.api)
	var in string
	for {
		fmt.Scanln(&in)
		if in == "exit" {
			log.Println("finishing client")
			break
		}
		time.Sleep(time.Second)
	}
}
