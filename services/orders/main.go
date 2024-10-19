package main

import "log"

func main() {
	gRPCServer := NewGRPCServer(":50051")
	if err := gRPCServer.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
