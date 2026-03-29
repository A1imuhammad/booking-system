package main

import (
	pb "booking-system/protoc/gen/go/auth"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAuthServer
}

func (s *server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	return &pb.RegisterResponse{}, nil
}

func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{Token: "test-token"}, nil
}

func (s *server) IsAdmin(ctx context.Context, in *pb.IsAdminRequest) (*pb.IsAdminResponse, error) {
	return &pb.IsAdminResponse{IsAdmin: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServer(s, &server{})

	log.Println("Auth service running on :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
