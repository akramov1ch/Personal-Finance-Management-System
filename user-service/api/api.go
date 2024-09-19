package api

import (
	"log"
	"net"
	"user-service/protos/genuser"
	"user-service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func ConnectApi() {
	// lg := logs.GetLogger("./logs/logger.log")
	lis, err := net.Listen("tcp", "localhost:7070")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	server := service.ConnServer()
	if server == nil {
		log.Fatalf("wrong to connect server")
	}

	grpcServer := grpc.NewServer()

	genuser.RegisterUserServiceServer(grpcServer, server)

	reflection.Register(grpcServer)

	log.Println("Server running on :7070")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
