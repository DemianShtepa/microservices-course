package main

import (
	"context"
	"fmt"
	"github.com/DemianShtepa/microservices-course/auth/pkg/grpc/v1/user"
	"github.com/brianvoe/gofakeit/v7"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

const grpcPort = 8001

type server struct {
	user.UnimplementedUserV1Server
}

func (s *server) Create(ctx context.Context, request *user.CreateRequest) (*user.CreateResponse, error) {
	return &user.CreateResponse{
		Id: 1,
	}, nil
}

func (s *server) Get(ctx context.Context, request *user.GetRequest) (*user.GetResponse, error) {
	return &user.GetResponse{
		User: &user.User{
			Id:        gofakeit.Uint64(),
			Name:      gofakeit.Name(),
			Email:     gofakeit.Email(),
			Role:      user.Role(gofakeit.RandomMapKey(user.Role_name).(int32)),
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
		},
	}, nil
}

func (s *server) Update(ctx context.Context, request *user.UpdateRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, request *user.DeleteRequest) (*emptypb.Empty, error) {
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
	user.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %s", listener.Addr())

	if err = s.Serve(listener); err != nil {
		log.Fatalf("failed to to serve: %s", err)
	}
}
