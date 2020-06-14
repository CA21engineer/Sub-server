package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"
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
	wg := new(sync.WaitGroup)
	wg.Add(4)

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
	go func() {
		subscriptionService := &service.SubscriptionServiceImpl{}
		subscription.RegisterSubscriptionServiceServer(server, subscriptionService)
		reflection.Register(server)
		_ = server.Serve(lis)
		wg.Done()
	}()

	// monitoring metrics, process and grpc
	go func() {
		processCollector := prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{Namespace: namespace})
		prometheus.MustRegister(m, processCollector)
		grpc_prometheus.Register(server)
		http.Handle("/metrics", promhttp.Handler())
		_ = http.ListenAndServe(":2112", nil)
		wg.Done()
	}()

	go func() {
		serverPort := "9090"
		r := gin.Default()
		r.Use(
			reverseProxy(
				"/",
				&url.URL{Scheme: "http", Host: config.GetEnvString("PROMETHEUS_NAMESPACE", "prometheus-service-server") + ":80"},
			),
		)
		_ = r.Run(fmt.Sprintf(":%s", serverPort))
		wg.Done()
	}()

	go func() {
		serverPort := "3000"
		r := gin.Default()
		r.Use(
			reverseProxy(
				"/",
				&url.URL{Scheme: "http", Host: config.GetEnvString("GRAFANA_NAMESPACE", "grafana-service") + ":80"},
			),
		)
		_ = r.Run(fmt.Sprintf(":%s", serverPort))
		wg.Done()
	}()

	wg.Wait()
}

func reverseProxy(urlPrefix string, target *url.URL) gin.HandlerFunc {
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.FlushInterval = -1

	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, urlPrefix) {
			c.Request.URL.Path = strings.Replace(c.Request.URL.Path, urlPrefix, "", 1)
			proxy.ServeHTTP(c.Writer, c.Request)
		}
	}
}
