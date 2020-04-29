package tfclient

import (
	log "github.com/sirupsen/logrus"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	"sync"

	tf "tensorflow_serving/apis"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type PredictionClient struct {
	mu      sync.RWMutex
	rpcConn *grpc.ClientConn
	svcConn tf.PredictionServiceClient
}

type Prediction struct {
	HomeWin float32
	Draw    float32
	AwayWin float32
}

func NewClient(addr string) (*PredictionClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := tf.NewPredictionServiceClient(conn)
	return &PredictionClient{rpcConn: conn, svcConn: c}, nil
}

func (c *PredictionClient) Predict(modelName string, inputs []float32) (*Prediction, error) {
	resp, err := c.svcConn.Predict(context.Background(), &tf.PredictRequest{
		ModelSpec: &tf.ModelSpec{
			Name: modelName,
		},
		Inputs: map[string]*tensor_go_proto.TensorProto{
			"dense_input": {
				Dtype:     types_go_proto.DataType_DT_FLOAT,
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

	log.Printf("OUTPUTS: %+v", resp.GetOutputs())
	log.Printf("Model spec: %+v", resp.GetModelSpec())

	res := &Prediction{
		HomeWin: 1/resp.GetOutputs()["dense_3"].FloatVal[0],
		Draw:    1/resp.GetOutputs()["dense_3"].FloatVal[1],
		AwayWin: 1/resp.GetOutputs()["dense_3"].FloatVal[2],
	}

	return res, nil
}

func newPredictRequest(modelName string) (pr *tf.PredictRequest) {
	return &tf.PredictRequest{
		ModelSpec: &tf.ModelSpec{
			Name: modelName,
		},
		Inputs: make(map[string]*tensor_go_proto.TensorProto),
	}
}

func (c *PredictionClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.svcConn = nil
	return c.rpcConn.Close()
}
