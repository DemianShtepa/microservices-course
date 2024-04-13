package main

import (
	"context"
	"fmt"
	"github.com/DemianShtepa/microservices-course/auth/pkg/grpc/v1/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const address = "localhost:8001"

func main() {
	connection, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %s", err)
	}
	defer connection.Close()

	client := user.NewUserV1Client(connection)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, err := client.Get(ctx, &user.GetRequest{Id: 12})
	if err != nil {
		log.Fatalf("failed to request: %s", err)
	}

	fmt.Println(response)
}
