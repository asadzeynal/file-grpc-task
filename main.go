package main

import (
	"log"
	"net"

	pb "github.com/asadzeynal/file-grpc-task/gen/file/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	server := NewServer("files")

	// TODO: Implement graceful shutdown
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterFileServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting grpc server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
