package main

import (
	"log"
	"net"

	grpcauth "auth-service/internal/grpc"

	authpb "github.com/BorBoris23/auth-proto/gen/auth"

	"google.golang.org/grpc"
)

func startGRPCServer(authServer *grpcauth.AuthServer) {
	grpcServer := grpc.NewServer()

	authpb.RegisterAuthServiceServer(
		grpcServer,
		authServer,
	)

	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Auth gRPC service started on :9090")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
