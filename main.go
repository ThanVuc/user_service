package main

import (
	"log"
	"user_service/cmd"
)

func main() {
	log.Println("gRPC servers are running...\n")
	cmd.RunGRPCServer()
	// run the grpc server
}
