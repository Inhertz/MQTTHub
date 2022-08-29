package domain

import (
	"MQTTHub/internal/core/models"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRegComplete(t *testing.T) {
	domain := New()

	answer, err := domain.RegComplete(models.Sensor{ID: 1, DeviceAdd: "test"}, models.Measurement{Value: 21, Unit: "Celcius"})
	if err != nil {
		t.Logf("expected: %v, got: %v", nil, err)
	}

	require.Equal(t, models.Measurement{Value: 21, Unit: "Celcius", Date: time.Now(), IDSensor: 1}, answer)
}

func TestSplitTopicTree(t *testing.T) {
	domain := New()

	answer, err := domain.SplitTopicTree("test/long/tree/sensor")

	if err != nil {
		t.Fatalf("expected: %v, got %v", nil, err)
	}

	require.Equal(t, []string{"test", "long", "tree", "sensor"}, answer)
}
