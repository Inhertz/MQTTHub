package serializer

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Adapter struct{}

// NewAdapter creates a new Adapter
func NewAdapter() *Adapter {
	return &Adapter{}
}

// Decode takes a JSON binary input and a pointer to the model to unmarshal
func (jsona Adapter) Decode(input []byte, model interface{}) error {
	if err := json.Unmarshal(input, model); err != nil {
		return errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return nil
}

// Encode takes a struct, marshals it and returns a JSON byte slice
func (jsona Adapter) Encode(input interface{}) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
