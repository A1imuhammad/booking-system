package server

import (
	authgrpc "booking-system/services/auth/internal/grpc"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func NewServer(log *slog.Logger, port int, authService authgrpc.Auth) *Server {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer, authService)

	return &Server{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (s *Server) Run() error {
	const op = "server.Run"

	log := s.log.With(slog.String("op", op), slog.Int("port", s.port))

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("gRPC server listening", slog.String("addr", listen.Addr().String()))

	if err := s.gRPCServer.Serve(listen); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}


func (s *Server) Stop() {
	const op = "grpcapp.Stop"

	s.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", s.port))

	s.gRPCServer.GracefulStop()
}