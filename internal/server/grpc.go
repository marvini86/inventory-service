package server

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

// GrpcServer represents a gRPC server
type GrpcServer struct {
	grpcServer *grpc.Server
	port       string
}

// NewGrpcServer creates a new GrpcServer instance
func NewGrpcServer() *GrpcServer {
	serverPort := os.Getenv("GRPC_PORT")
	if serverPort == "" {
		serverPort = "50051"
	}
	return &GrpcServer{grpcServer: grpc.NewServer(), port: serverPort}
}

// Run runs the gRPC server
func (s *GrpcServer) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	defer l.Close()

	log.Printf("Starting gRPC server on port %s", s.port)

	serviceRegister := NewServiceRegister(s.grpcServer)
	serviceRegister.Register()

	if err = s.grpcServer.Serve(l); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}
