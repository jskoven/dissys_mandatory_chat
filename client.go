package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	gRPC_template "github.com/jskoven/gRPC_template/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	//If connection to other computer, set up a hotspot from phone and connect to it.
	//Then find your IP and insert it before the port below. As is, it runs on localHost
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := gRPC_template.NewChatServiceClient(conn)

	for {
		Scanner := bufio.NewScanner(os.Stdin)
		Scanner.Scan()
		textToSend := Scanner.Text()
		if textToSend == "exit" {
			fmt.Println("##- Exiting chat service -##")
			break
		}

		//Message struct from proto file is created and used here, which is then sent to server.
		message := gRPC_template.Message{
			MessageToBeSent: textToSend,
		}

		//Since ServerDef.ReceiveMessage has a return value that is the answer to the message
		//we simple define response as the return value of the function:
		response, err := c.ReceiveMessage(context.Background(), &message)
		if err != nil {
			log.Fatalf("Error when calling ReceiveMessage: %s", err)
		}

		log.Printf("Response from server: %s", response.MessageToBeSent)

	}
}
