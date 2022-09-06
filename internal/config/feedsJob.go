package config

import (
	"context"
	"errors"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/runner/listener"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net/http"
)

const PrometheusPortEnv = "PROMETHEUS_PORT"

type FeedsServerJob struct {
	feedsServer      proto.FeedsServer
	server           *grpc.Server
	listenerResource *listener.Resource
	httpServer       *http.Server
	logger           *zap.Logger
	interceptor      func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)
}

func NewFeedsServerJob(logger *zap.Logger, feedsServer proto.FeedsServer, listenerResource *listener.Resource, interceptor func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error)) *FeedsServerJob {
	return &FeedsServerJob{
		logger:           logger,
		feedsServer:      feedsServer,
		listenerResource: listenerResource,
		interceptor:      interceptor,
	}
}

func (job *FeedsServerJob) Run() error {
	grpcMetrics := grpc_prometheus.NewServerMetrics()
	job.server = grpc.NewServer(
		grpc.StreamInterceptor(grpcMetrics.StreamServerInterceptor()),
		grpc.ChainUnaryInterceptor(grpcMetrics.UnaryServerInterceptor(), job.interceptor),
	)
	proto.RegisterFeedsServer(job.server, job.feedsServer)
	reflection.Register(job.server)
	//go func() {
	//	prometheusPort := os.Getenv(PrometheusPortEnv)
	//	if prometheusPort == "" {
	//		job.logger.Fatal(fmt.Sprintf("%s environment variable is not set", PrometheusPortEnv))
	//	}
	//	grpcMetrics.InitializeMetrics(job.server)
	//	reg := prometheus.NewRegistry()
	//	reg.MustRegister(grpcMetrics)
	//	job.httpServer = &http.Server{Handler: promhttp.HandlerFor(reg, promhttp.HandlerOpts{}), Addr: fmt.Sprintf("0.0.0.0:%s", prometheusPort)}
	//	if err := job.httpServer.ListenAndServe(); err != nil {
	//		job.logger.Fatal("httpServer listen error", zap.Error(err))
	//	}
	//}()
	return job.server.Serve(job.listenerResource.Listener)
}

func (job *FeedsServerJob) Shutdown(ctx context.Context) error {
	go func() {
		if err := job.httpServer.Shutdown(ctx); err != nil {
			job.logger.Fatal("httpServer shutdown error", zap.Error(err))
		}
	}()
	done := make(chan struct{})
	go func() {
		job.server.GracefulStop()
		done <- struct{}{}
	}()
	select {
	case <-ctx.Done():
		return errors.New("reached timeout stopping profiles server job")
	case <-done:
		return nil
	}
}
