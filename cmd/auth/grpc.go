package main

import (
	"fmt"
	"log"
	"net"
	"os"

	grpcserver "auth-service/internal/grpc"

	authpb "github.com/BorBoris23/auth-proto/gen/auth"
	userpb "github.com/BorBoris23/auth-proto/gen/users"

	"google.golang.org/grpc"
)

func startGRPCAuthServer(authServer *grpcserver.AuthServer) {
	port := os.Getenv("AUTH_GRPC_PORT")

	grpcServer := grpc.NewServer()

	authpb.RegisterAuthServiceServer(
		grpcServer,
		authServer,
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Auth gRPC service started on :%s", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

func startGRPCUserServer(userServer *grpcserver.UserServer) {
	port := os.Getenv("USER_GRPC_PORT")

	grpcServer := grpc.NewServer()

	userpb.RegisterUserServiceServer(
		grpcServer,
		userServer,
	)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("User gRPC service started on :%s", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
