package main

import (
	"log"
	"net"

	gRPC_template "github.com/jskoven/gRPC_template/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	//hello
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %t", err)
	}

	//Server struct
	s := gRPC_template.Server{}

	//grpc server command
	grpcSever := grpc.NewServer()

	//Starting server
	gRPC_template.RegisterChatServiceServer(grpcSever, &s)
	if err := grpcSever.Serve(lis); err != nil {
		log.Fatalf("Failed")
	}

}
