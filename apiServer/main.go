package main

import (
	"github.com/BambooTuna/go-server-lib/config"
	subscription "github.com/CA21engineer/Sub-server/apiServer/pb"
	"github.com/CA21engineer/Sub-server/apiServer/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	serverPort := config.GetEnvString("PORT", "18080")

	lis, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	subscriptionService := &service.SubscriptionServiceImpl{}
	subscription.RegisterSubscriptionServiceServer(server, subscriptionService)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
