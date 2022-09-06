package config

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Validatable interface {
	Validate() error
}

func GetServerInterceptor(logger *zap.Logger) func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		msg := req.(proto.Message)
		logger.Debug(fmt.Sprintf("Request: %s %s", info.FullMethod, protojson.Format(msg)))
		var err error
		var resp interface{}
		if validatable, ok := req.(Validatable); ok {
			err = validatable.Validate()
		}
		if err == nil {
			resp, err = handler(ctx, req)
		}
		if err == nil {
			if validatable, ok := resp.(Validatable); ok {
				err = validatable.Validate()
			}
		}
		if err == nil && resp != nil {
			msg = resp.(proto.Message)
			logger.Debug(fmt.Sprintf("Response: %s %s", info.FullMethod, protojson.Format(msg)))
		} else if errStatus := status.Convert(err); errStatus.Code() == codes.PermissionDenied {
			logger.Warn(errStatus.Message())
		} else if errStatus.Code() == codes.InvalidArgument {
			logger.Warn(errStatus.Message())
		} else if errStatus.Code() == codes.NotFound {
			logger.Warn(errStatus.Message())
		} else {
			logger.Error(errStatus.Message())
		}
		return resp, err
	}
}
