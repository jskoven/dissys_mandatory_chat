package main

import (
	"log"
	"net"

	dissys_mandatory_chat "github.com/jskoven/dissys_mandatory_chat/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	//hello
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %t", err)
	}

	//Server struct
	s := dissys_mandatory_chat.Server{}

	//grpc server command
	grpcSever := grpc.NewServer()

	//Starting server
	dissys_mandatory_chat.RegisterChatServiceServer(grpcSever, &s)
	if err := grpcSever.Serve(lis); err != nil {
		log.Fatalf("Failed")
	}

}
