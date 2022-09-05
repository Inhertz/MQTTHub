package domain

import (
	"MQTTHub/internal/core/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegComplete(t *testing.T) {
	domain := New()

	answer, err := domain.RegComplete(models.Sensor{ID: 1, DeviceAdd: "test"}, models.Measurement{Value: 21, Unit: "Celcius"})
	if err != nil {
		t.Logf("expected: %v, got: %v", nil, err)
	}

	expected := models.Measurement{Value: 21, Unit: "Celcius", IDSensor: 1}

	//evaluates everithing except date
	require.Equal(t, expected.Value, answer.Value)
	require.Equal(t, expected.Unit, answer.Unit)
	require.Equal(t, expected.IDSensor, answer.IDSensor)
	require.NotEqual(t, expected.Date, answer.Date)

}

func TestSplitTopicTree(t *testing.T) {

	topicTreePairsOk := []struct {
		name   string
		input  string
		output []string
	}{
		{"OK 4 levels", "test/long/tree/sensor", []string{"test", "long", "tree", "sensor"}},
		{"OK 2 levels", "test/short", []string{"test", "short"}},
	}

	topicTreePairsFail := []struct {
		name   string
		input  string
		output string
	}{
		{"wrong 1 level", "test", "topic tree should be at least 2 levels, follow mqtt doc"},
		{"wrong separator", "test-wrong-separator", "topic tree should be at least 2 levels, follow mqtt doc"},
	}
	domain := New()

	for _, topicTree := range topicTreePairsOk {
		t.Run(topicTree.name, func(t *testing.T) {
			answer, err := domain.SplitTopicTree(topicTree.input)

			if err != nil {
				t.Errorf("expected: %v, got %v", nil, err)
			}

			require.Equal(t, topicTree.output, answer)
		})
	}

	for _, topicTree := range topicTreePairsFail {
		t.Run(topicTree.name, func(t *testing.T) {
			answer, err := domain.SplitTopicTree(topicTree.input)

			if answer != nil {
				t.Errorf("expected: %v, got %v", nil, answer)
			}

			require.Equal(t, topicTree.output, err.Error())
		})
	}

}
