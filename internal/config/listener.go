package config

import (
	"fmt"

	"github.com/nostressdev/runner"
	"github.com/nostressdev/runner/listener"
	"go.uber.org/zap"
)

const grpcListenerGroup = "GRPC"
const probesListenerGroup = "PROBE"

func NewGrpcListener(logger *zap.Logger, variableProvider runner.VariableProvider) *listener.Resource {
	if config, err := listener.NewConfigFromProvider(
		variableProvider,
		grpcListenerGroup,
		false,
	); err != nil {
		logger.Fatal(fmt.Sprintf("failed to create grpc network listener, check %s_ADDR and %s_PORT variables", grpcListenerGroup, grpcListenerGroup), zap.Error(err))
		return nil
	} else {
		return listener.New(config)
	}
}

func NewProbeListener(logger *zap.Logger, variableProvider runner.VariableProvider) *listener.Resource {
	if config, err := listener.NewConfigFromProvider(
		variableProvider,
		probesListenerGroup,
		false,
	); err != nil {
		logger.Fatal(fmt.Sprintf("failed to create grpc network listener, check %s_ADDR and %s_PORT variables", probesListenerGroup, probesListenerGroup), zap.Error(err))
		return nil
	} else {
		return listener.New(config)
	}
}
