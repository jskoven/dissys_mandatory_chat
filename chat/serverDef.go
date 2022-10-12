package dissys_mandatory_chat

import (
	"bufio"
	"log"
	"os"

	"golang.org/x/net/context"
)

// Server struct
type Server struct {
	//Must contain this, whatever it is
	UnimplementedChatServiceServer
	messages         []string
	amountOfMessages int
}

func (s *Server) ReceiveMessageOld(ctx context.Context, message *Message) (*Message, error) {
	//log.Printf("## Received message body from client: %s ##", message.MessageToBeSent)
	log.Printf("%s: %s", message.ClientUsername, message.MessageToBeSent)
	//Scanner only used so as to be able to send back messages to client from server.
	Scanner := bufio.NewScanner(os.Stdin)
	log.Printf("## What do you wish to respond with? ##")
	Scanner.Scan()
	textToSend := Scanner.Text()
	//The return value is the same message struct as is received, but with overwritten and changed values.
	return &Message{MessageToBeSent: textToSend}, nil
}

func (s *Server) ReceiveMessage(ctx context.Context) (*[]string, error) {
	return &s.messages, nil
}

func (s *Server) SendMessage(ctx context.Context, message *Message) {
	s.messages[s.amountOfMessages+1] = message.MessageToBeSent
}
