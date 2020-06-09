package main

import (
	"fmt"
	"log"
	"net"

	"Subs-server/apiServer/models"
	subscription "Subs-server/apiServer/pb"
	"Subs-server/apiServer/service"

	"github.com/BambooTuna/go-server-lib/config"

	"google.golang.org/grpc"
)

func main() {

	serverPort := config.GetEnvString("PORT", "18080")

	lis, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Printf("==========%#v\n", models.DB)

	server := grpc.NewServer()
	subscriptionService := &service.SubscriptionServiceImpl{}
	subscription.RegisterSubscriptionServiceServer(server, subscriptionService)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
