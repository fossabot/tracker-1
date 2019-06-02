package api

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --gogo_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. dts.proto

var (
	// ErrModuleNotFound occurs when a module cannot be found in the graph
	ErrModuleNotFound = status.Error(codes.NotFound, "failed to locate module")
	// ErrPartialDeletion occurs when a partial deletion occurs during Put
	ErrPartialDeletion = status.Error(codes.Internal, "failed to delete removed edges")
	// ErrPartialInsertion occurs when a partial insertion occurs during Put
	ErrPartialInsertion = status.Error(codes.Internal,"failed to insert new edges")
	// ErrUnimplemented occurs when a method has not yet been implemented
	ErrUnimplemented = status.Error(codes.Unimplemented, "unimplemented")
)
