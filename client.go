package main

import (
	"context"
	"log"

	"github.com/jeferagudeloc/go-grpc-server-example/greeting"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to server: %v", err)
	}

	defer conn.Close()

	c := greeting.NewGreetingServiceClient(conn)

	greeting := greeting.Greeting{
		From:    "jeferagudeloc",
		To:      "perrito",
		Content: "hi perrito",
	}

	response, err := c.SendGreetings(context.Background(), &greeting)

	if err != nil {
		log.Fatalf("Error when trying to say hello: %v", err)
	}

	log.Printf("Response from server: %v", response)

}
