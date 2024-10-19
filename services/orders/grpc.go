package main

import (
	"demo-grpc/services/orders/handler"
	"demo-grpc/services/orders/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	gRPCServer := grpc.NewServer()

	// register our grpc services
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(gRPCServer, orderService)

	log.Println("Starting gRPC server on", s.addr)

	return gRPCServer.Serve(lis)
}
