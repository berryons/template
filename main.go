package main

import (
	"fmt"
	"github.com/berryons/log"
	"github.com/berryons/server"
	"github.com/berryons/template/config"
	"google.golang.org/grpc"
)

func main() {
	// Set Environment Variables
	conf := config.Load()

	// Set Logger
	if len(conf.LogDir) != 0 {
		log.SetLogger(
			conf.LogLevel,
			conf.LogDir,
			conf.ServerName,
			conf.LogFileLevel,
			2,
			false,
		)
	}

	runServer(conf.ServerNetwork, fmt.Sprintf("%s:%s", conf.ServerAddress, conf.ServerPort))
}

func runServer(network, address string) {
	// TODO: Add Unary Interceptors
	var unaryInterceptors []grpc.UnaryServerInterceptor

	// TODO: Add Stream Interceptors
	var streamInterceptors []grpc.StreamServerInterceptor

	// Generate gRPC Server
	grpcServer := server.New(network, address, unaryInterceptors, streamInterceptors)

	// TODO: Register gRPC Handler

	// Run gRPC Server
	grpcServer.Run()
}
