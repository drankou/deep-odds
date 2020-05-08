package client

import (
	"testing"
)

const defaultGrpcTfServer = "localhost:8500"

func TestNewPredictionClient(t *testing.T) {
	c, err := NewPredictionClient(defaultGrpcTfServer)
	if err != nil{
		t.Fatal(err)
	}

	if c == nil{
		t.Fatal("nil prediction client")
	}
}

func TestPredictionClient_Predict(t *testing.T) {
	c, err := NewPredictionClient(defaultGrpcTfServer)
	if err != nil{
		t.Fatal(err)
	}

	inputs := []float32{84, 2, 2, 80, 84, 77, 44, 10, 2, 3, 6, 8, 1, 1, 2, 0, 0, 0, 0, 6.5, 1.3, 5.5, 3.25, 4.2, 1.9}
	prediction, err := c.Predict("model_67", inputs)
	if err != nil{
		t.Fatal(err)
	}

	t.Logf("%+v", prediction)
}
