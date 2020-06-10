package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/CA21engineer/Subs-server/apiServer/models"

	"github.com/BambooTuna/go-server-lib/config"
	"github.com/BambooTuna/go-server-lib/metrics"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
	"github.com/CA21engineer/Subs-server/apiServer/service"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const namespace = "go_server"

func main() {
	m := metrics.CreateMetrics(namespace)
	go func() {
		processCollector := prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{Namespace: namespace})
		prometheus.MustRegister(m, processCollector)
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

	models.ConnectDB()

	server := grpc.NewServer()
	subscriptionService := &service.SubscriptionServiceImpl{}
	subscription.RegisterSubscriptionServiceServer(server, subscriptionService)
	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
