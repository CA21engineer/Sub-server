package main

import (
	"Subs-server/apiServer/models"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	subscription "Subs-server/apiServer/pb"
	"Subs-server/apiServer/service"

	"github.com/BambooTuna/go-server-lib/config"
	"github.com/BambooTuna/go-server-lib/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {
	m := metrics.CreateMetrics("go_server")
	go func() {
		prometheus.MustRegister(m)
		http.Handle("/metrics", promhttp.Handler())
		_ = http.ListenAndServe(":2112", nil)
	}()

	go func() {
		health := m.Gauge("health", map[string]string{})
		health.Set(200)
		ticker := time.NewTicker(time.Minute * 1)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				health.Set(200)
			}
		}
	}()

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
