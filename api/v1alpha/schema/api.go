package schema

//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --gogo_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. schema.proto
