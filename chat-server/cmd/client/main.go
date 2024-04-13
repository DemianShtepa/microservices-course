package main

import (
	"context"
	"fmt"
	"github.com/DemianShtepa/microservices-course/chat-server/pkg/grpc/v1/chat"
	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const address = "localhost:8002"

func main() {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	defer connection.Close()

	client := chat.NewChatV1Client(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := client.Create(ctx, &chat.CreateRequest{Names: []string{gofakeit.Name()}})
	if err != nil {
		log.Fatalf("failed to request: %s", err)
	}

	fmt.Println(response)
}
