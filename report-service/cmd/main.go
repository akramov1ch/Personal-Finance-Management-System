package main

import (
	"context"
	"log"
	"net"
	"report-service/internal/config"
	"report-service/protos/reporting"
	"report-service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	notificationServer := service.NewNotificationServer()

	lis, err := net.Listen("tcp", cfg.ReportingServer.Http.Host+":"+cfg.ReportingServer.Http.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcInterceptor),
	)
	reflection.Register(grpcServer)

	reporting.RegisterReportingServiceServer(grpcServer, notificationServer)

	log.Println("main: server running  port", cfg.ReportingServer.Http.Port)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatal("failed while listening:", (err))
	}

}

func grpcInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	m, err := handler(ctx, req)
	if err != nil {
		log.Printf("RPC failed with error: %v", err)
		return nil, err
	}
	return m, nil
}
