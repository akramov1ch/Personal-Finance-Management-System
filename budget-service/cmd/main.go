package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"income-expense-service/config"
	"income-expense-service/db"
	"income-expense-service/internal/handlers"
	"income-expense-service/internal/service"
	"income-expense-service/redis"

	pb "income-expense-service/grpc"

	"google.golang.org/grpc"
)

func main() {
	config.LoadConfig()

	dbConn, err := db.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	redisClient := redis.NewRedisClient()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBudgetServiceServer(grpcServer, &handlers.BudgetHandler{
		Service: service.NewBudgetService(dbConn, redisClient),
	})

	log.Printf("gRPC server is running on port %s", os.Getenv("GRPC_PORT"))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
