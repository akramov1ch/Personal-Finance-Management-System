package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"notifaction-service/internal/config"
	"notifaction-service/internal/kafka"
	notifactionpb "notifaction-service/protos/notifaction"
	"notifaction-service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	consumer := kafka.ConnKafka(cfg)

	notificationServer := service.NewNotificationServer(consumer)
	go notificationServer.ListenForKafkaMessages()

	go func() {
		websocketPort := cfg.WebSocketserver.Http.Port
		websocketHost := cfg.WebSocketserver.Http.Host
		log.Println("WebSocket server started at port:", websocketPort)
		http.HandleFunc("/ws", notificationServer.WebSocketHandler)
		if err := http.ListenAndServe(websocketHost+":"+websocketPort, nil); err != nil {
			log.Fatalf("Failed to start WebSocket server: %v", err)
		}
	}()

	grpcPort := cfg.NotifactionServer.Http.Port
	grpcHost := cfg.NotifactionServer.Http.Host
	lis, err := net.Listen("tcp", grpcHost+":"+grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcInterceptor),
	)
	reflection.Register(grpcServer)

	notifactionpb.RegisterNotificationServiceServer(grpcServer, notificationServer)

	log.Println("gRPC server running on port", grpcPort)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("failed while listening:", err)
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
