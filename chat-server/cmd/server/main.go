package main

import (
	"context"
	"fmt"
	"github.com/DemianShtepa/microservices-course/chat-server/pkg/grpc/v1/chat"
	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

const grpcPort = 8002

type server struct {
	chat.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, request *chat.CreateRequest) (*chat.CreateResponse, error) {
	return &chat.CreateResponse{Id: gofakeit.Uint64()}, nil
}

func (s *server) Delete(ctx context.Context, request *chat.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, request *chat.SendMessageRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}
	defer listener.Close()

	s := grpc.NewServer()
	reflection.Register(s)
	chat.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %s", listener.Addr())

	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
