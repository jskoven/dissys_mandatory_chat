package gRPC_template

import (
	"bufio"
	"log"
	"os"
	"time"

	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Received message body from client: %s", message.Body)
	log.Printf("Received time from client: %s", message.TimeNow)
	Scanner := bufio.NewScanner(os.Stdin)
	log.Printf("What do you wish to respond with?")
	Scanner.Scan()
	textToSend := Scanner.Text()
	return &Message{Body: textToSend, TimeNow: time.Now().String()}, nil
}
