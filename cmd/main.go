package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/nostressdev/feeds-framework/internal/feeds"
	"github.com/nostressdev/feeds-framework/internal/storage"
	"github.com/nostressdev/feeds-framework/internal/utils"
	pb "github.com/nostressdev/feeds-framework/proto"
	"github.com/nostressdev/signer"
	"go.uber.org/zap"

	"github.com/apple/foundationdb/bindings/go/src/fdb"
	"github.com/apple/foundationdb/bindings/go/src/fdb/subspace"
	"google.golang.org/grpc"
)

var (
	host      string
	port      string
	secretKey string
	logging   string
	nodeNumber string
)

func loadEnvironmentVariables() {
	if host = os.Getenv("HOST"); host == "" {
		log.Fatalln("$HOST environment variable should be specified")
	}
	if port = os.Getenv("PORT"); port == "" {
		log.Fatalln("$PORT environment variable should be specified")
	}
	if secretKey = os.Getenv("SECRET_KEY"); secretKey == "" {
		log.Fatalln("SECRET_KEY environment variable should be specified")
	}
	if logging = os.Getenv("LOGGING"); logging == "" {
		log.Fatalln("LOGGING environment variable should be specified")
	}
	if nodeNumber = os.Getenv("NODE_NUMBER"); nodeNumber == "" {
		log.Fatalln("NODE_NUMBER environment variable should be specified")
	}
}

func createNetworkListener() net.Listener {
	addr := fmt.Sprintf("%s:%s", host, port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start serving with addr: %v\n", addr)
	return listener
}

var serviceSubspace = subspace.Sub("feeds_service")

func createGrpcNetworkListener() net.Listener {
	addr := fmt.Sprintf("%s:%s", host, port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start serving grpc with addr: %v\n", addr)
	return listener
}

func main() {
	fdb.MustAPIVersion(600)
	db := fdb.MustOpenDefault()

	loadEnvironmentVariables()

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
		DB: db,
		 Subspace: serviceSubspace,
		 Generator: node,
		})

	signer := signer.NewSignerJWT(signer.TokenProviderConfig{
		Expiration: time.Hour,
		SecretKey:  []byte(secretKey),
	})

	var logger *zap.Logger
	if logging == "DEVELOPMENT" {
		logger, err = zap.NewDevelopment()
	} else if logging == "PRODUCTION" {
		logger, err = zap.NewProduction()
	} else {
		log.Fatalln(fmt.Sprintf("unknown logging mode: %s", logging))
	}
	if err != nil {
		log.Fatalln("failed to create zap logger")
	}



	server := grpc.NewServer(utils.GetServerInterceptor(logger, signer))
	pb.RegisterFeedsServer(server, feeds.New(&feeds.Config{
		Signer:  signer,
		Storage: feedsStorage,
	}))
	listener := createGrpcNetworkListener()
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
