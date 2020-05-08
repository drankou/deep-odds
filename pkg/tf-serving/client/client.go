package client

import (
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	tf "tensorflow_serving/apis"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type PredictionClient struct {
	srvClient tf.PredictionServiceClient
}

type Prediction struct {
	HomeWin float32
	Draw    float32
	AwayWin float32
}

func NewPredictionClient(address string) (*PredictionClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := tf.NewPredictionServiceClient(conn)
	return &PredictionClient{srvClient: c}, nil
}

func (c *PredictionClient) Predict(modelName string, inputs []float32) (*Prediction, error) {
	resp, err := c.srvClient.Predict(context.Background(), &tf.PredictRequest{
		ModelSpec: &tf.ModelSpec{
			Name: modelName,
		},
		Inputs: map[string]*tensor_go_proto.TensorProto{
			"dense_input": {
				Dtype:    types_go_proto.DataType_DT_FLOAT,
				FloatVal: inputs,
				TensorShape: &tensor_shape_go_proto.TensorShapeProto{
					Dim: []*tensor_shape_go_proto.TensorShapeProto_Dim{{Size: 1}, {Size: 25}},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var homeOdds float32
	var drawOdds float32
	var awayOdds float32

	if resp.GetOutputs()["dense_3"].FloatVal[0] != 0 {
		homeOdds = 1 / resp.GetOutputs()["dense_3"].FloatVal[0]
	}

	if resp.GetOutputs()["dense_3"].FloatVal[1] != 0 {
		drawOdds = 1 / resp.GetOutputs()["dense_3"].FloatVal[1]
	}

	if resp.GetOutputs()["dense_3"].FloatVal[2] != 0 {
		awayOdds = 1 / resp.GetOutputs()["dense_3"].FloatVal[2]
	}
	res := &Prediction{
		HomeWin: homeOdds,
		Draw:    drawOdds,
		AwayWin: awayOdds,
	}

	return res, nil
}
