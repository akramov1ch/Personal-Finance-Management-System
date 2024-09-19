package methods

import (
	grp "income-expense-service/grpc"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectUser() grp.UserServiceClient {
	conn, err := grpc.NewClient("localhost:7070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("error connect user micro...", err)
	}

	client := grp.NewUserServiceClient(conn)
	return client
}
