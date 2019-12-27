package infra

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/inaohiro/jwt-sample/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type HttpServer struct {
	Server *http.Server
}

func NewHttpServer(cfg *config.ServerEnv, handler http.Handler) *HttpServer {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port),
		Handler: handler,
	}

	return &HttpServer{Server: server}
}

func (s *HttpServer) ListenAndServe() error {
	return s.Server.ListenAndServe()
}

// GrpcServer struct needs to have server address because the gRPC server requires Listener instance
type GrpcServer struct {
	addr   string
	server *grpc.Server
}

type GrpcServices interface {
	Register(*grpc.Server)
}

func NewGrpcServer(cfg *config.GrpcEnv, services []GrpcServices) *GrpcServer {
	server := grpc.NewServer()

	for _, svc := range services {
		svc.Register(server)
	}
	reflection.Register(server)

	return &GrpcServer{
		addr:   fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port),
		server: server,
	}
}

func (s *GrpcServer) ListenAndServe() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return s.server.Serve(lis)
}
