package main

import (
	"log"
	"net"

	gRPC_testing "github.com/jskoven/gRPC_template/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	//hello
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %t", err)
	}

	s := gRPC_testing.Server{}

	grpcSever := grpc.NewServer()

	gRPC_testing.RegisterChatServiceServer(grpcSever, &s)

	if err := grpcSever.Serve(lis); err != nil {
		log.Fatalf("Failed")
	}

}
