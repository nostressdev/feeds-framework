package main

import (
	"fmt"
	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"github.com/bwmarrin/snowflake"
	"github.com/nostressdev/feeds-framework/internal/config"
	"github.com/nostressdev/feeds-framework/internal/feeds"
	"github.com/nostressdev/feeds-framework/internal/storage"
	"github.com/nostressdev/runner"
	"github.com/nostressdev/runner/state"
	"github.com/nostressdev/signer"
	"go.uber.org/zap"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	secretKey  string
	nodeNumber string
)

func loadEnvironmentVariables() {
	if secretKey = os.Getenv("SECRET_KEY"); secretKey == "" {
		log.Fatalln("SECRET_KEY environment variable should be specified")
	}
	if nodeNumber = os.Getenv("NODE_NUMBER"); nodeNumber == "" {
		log.Fatalln("NODE_NUMBER environment variable should be specified")
	}
}

var serviceSubspace = subspace.Sub("feeds_service")

func main() {
	logger := config.NewLogger()

	fdb.MustAPIVersion(600)
	db := fdb.MustOpenDefault()

	loadEnvironmentVariables()
	environmentVariableProvider := runner.NewEnvironmentVariableProvider()
	grpcListener := config.NewGrpcListener(logger, environmentVariableProvider)
	probeListener := config.NewProbeListener(logger, environmentVariableProvider)

	number, err := strconv.Atoi(nodeNumber)
	if err != nil {
		log.Fatalln("failed to parse node number")
	}
	node, err := snowflake.NewNode(int64(number))
	if err != nil {
		fmt.Println(err)
		return
	}

	feedsStorage := storage.NewFeedsFDB(&storage.ConfigFeedsFDB{
		DB:        db,
		Subspace:  serviceSubspace,
		Generator: node,
	})

	sgn := signer.NewSignerJWT(signer.TokenProviderConfig{
		Expiration: time.Hour,
		SecretKey:  []byte(secretKey),
	})

	feedsServer := feeds.New(&feeds.Config{
		Signer:  sgn,
		Storage: feedsStorage,
	})

	app := runner.New(runner.DefaultConfig(),
		[]runner.Resource{
			grpcListener,
			probeListener,
		},
		[]runner.Job{
			config.NewFeedsServerJob(
				logger,
				feedsServer,
				grpcListener,
				config.GetServerInterceptor(logger),
			),
		},
	)
	app.AddJob(state.NewReadyLiveHttpJob(probeListener, app.State))

	logger.Info("feeds API server starting...")

	if err := app.Run(); err != nil {
		logger.Fatal("feeds API error during server execution", zap.Error(err))
	}
	logger.Info("feeds API server execution successfully finished!!!")
}
