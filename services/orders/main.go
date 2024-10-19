package main

import "log"

func main() {
	httpServer := NewHttpServer(":8000")
	go httpServer.Run()

	gRPCServer := NewGRPCServer(":50051")
	if err := gRPCServer.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
