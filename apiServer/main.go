package main

import (
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/CA21engineer/Subs-server/apiServer/models"

	"github.com/BambooTuna/go-server-lib/config"
	"github.com/BambooTuna/go-server-lib/metrics"
	subscription "github.com/CA21engineer/Subs-server/apiServer/pb"
	"github.com/CA21engineer/Subs-server/apiServer/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const namespace = "go_server"

func main() {
	m := metrics.CreateMetrics(namespace)
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

	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)
	subscriptionService := &service.SubscriptionServiceImpl{}
	subscription.RegisterSubscriptionServiceServer(server, subscriptionService)
	reflection.Register(server)

	// monitoring metrics, process and grpc
	go func() {
		processCollector := prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{Namespace: namespace})
		prometheus.MustRegister(m, processCollector)
		grpc_prometheus.Register(server)
		http.Handle("/metrics", promhttp.Handler())
		_ = http.ListenAndServe(":2112", nil)
	}()

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
