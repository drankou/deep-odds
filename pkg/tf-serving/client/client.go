package client

import (
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	tf "tensorflow_serving/apis"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var scaler = []float32{
	0.01111,
	0.05555,
	0.06666,
	0.003921,
	0.003921,
	0.004716,
	0.005208,
	0.028571,
	0.010526,
	0.033333,
	0.035714,
	0.041666,
	0.052631,
	0.1,
	0.06666,
	0.25,
	0.25,
	0.04166,
	0.045454,
	0.0019920,
	0.00980392,
	0.00199203,
	0.00980392,
	0.01923,
	0.004950,
}

type PredictionClient struct {
	srvClient tf.PredictionServiceClient
}

type Prediction struct {
	HomeWin float32
	Draw    float32
	AwayWin float32
}

func NewPredictionClient(address string) (*PredictionClient, error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(ctx, address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := tf.NewPredictionServiceClient(conn)
	return &PredictionClient{srvClient: c}, nil
}

func (c *PredictionClient) Predict(modelName string, inputs []float32) (*Prediction, error) {

	for i := range inputs {
		inputs[i] *= scaler[i]
	}

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

	var homeWinProb float32
	var drawProb float32
	var awayWinProb float32

	if resp.GetOutputs()["dense_2"].FloatVal[0] != 0 {
		homeWinProb = resp.GetOutputs()["dense_2"].FloatVal[0]
	}

	if resp.GetOutputs()["dense_2"].FloatVal[1] != 0 {
		awayWinProb = resp.GetOutputs()["dense_2"].FloatVal[1]
	}

	if resp.GetOutputs()["dense_2"].FloatVal[2] != 0 {
		drawProb = resp.GetOutputs()["dense_2"].FloatVal[2]
	}
	res := &Prediction{
		HomeWin: homeWinProb,
		Draw:    drawProb,
		AwayWin: awayWinProb,
	}

	return res, nil
}
