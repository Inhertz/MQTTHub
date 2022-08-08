package application

// APIPort is the technology neutral
// port for driving adapters
type APIPort interface {
	CreateMeasurement(topic string, payload []byte) error
}

// DbPort is the port for a db adapter
type DbPort interface {
	Create(modelP interface{}) error
	Find(parameter string, value string, modelP interface{}) error
}

// Serializer is the por for a marshaller adapter (json, bson, msgpack...)
type SerializerPort interface {
	Decode(input []byte, model interface{}) error
	Encode(input interface{}) ([]byte, error)
}
