#!/bin/bash

mkdir -p vendor/tensorflow/core/framework
cd vendor/tensorflow/core/framework
wget -N https://raw.githubusercontent.com/tensorflow/tensorflow/master/tensorflow/core/framework/{attr_value,graph,op_def,variable,versions,node_def,function,resource_handle,tensor,tensor_shape,types}.proto
cd ..

mkdir -p protobuf
cd protobuf
wget -N https://raw.githubusercontent.com/tensorflow/tensorflow/master/tensorflow/core/protobuf/{meta_graph,saver,saved_object_graph,struct,trackable_object_graph}.proto
cd ..

mkdir -p example
cd example
wget -N https://raw.githubusercontent.com/tensorflow/tensorflow/master/tensorflow/core/example/{example,feature}.proto
cd ../../..

mkdir -p tensorflow_serving/apis
cd tensorflow_serving/apis
wget -N https://raw.githubusercontent.com/tensorflow/serving/master/tensorflow_serving/apis/{classification,model,input,prediction_service,predict,get_model_metadata,inference,regression}.proto
cd ../..

protoc --go_out=. tensorflow/core/framework/*.proto
protoc --go_out=. tensorflow/core/protobuf/*.proto
protoc --go_out=. tensorflow/core/example/*.proto
protoc --go_out=plugins=grpc:. tensorflow_serving/apis/*.proto