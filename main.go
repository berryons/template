package main

import (
	"fmt"
	"github.com/berryons/log"
	"github.com/berryons/server"
	"github.com/berryons/template/variable"
	"google.golang.org/grpc"
)

func main() {
	// Set Environment Variables
	envVar := variable.New()

	// Set Logger
	if len(envVar.LogDir) != 0 {
		log.SetLogger(
			envVar.LogLevel,
			envVar.LogDir,
			envVar.ServerName,
			envVar.LogLevel,
			2,
			false,
		)
	}

	runServer(envVar.ServerNetwork, fmt.Sprintf("%s:%s", envVar.ServerAddress, envVar.ServerPort))
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
