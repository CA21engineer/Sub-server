package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/BambooTuna/go-server-lib/config"
	"github.com/BambooTuna/go-server-lib/metrics"
	models2 "github.com/CA21engineer/Subs-server/apiServer/models"
	"github.com/CA21engineer/Subs-server/pushServer/models"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/api/option"
	"net/http"
	"sync"
	"time"
)

const namespace = "go_push_server"

func main() {
	mode := config.GetEnvString("MODE", "DEV")
	ctx := context.Background()
	m := metrics.CreateMetrics(namespace)
	m.Counter("Server_Start_Total", map[string]string{}).Inc()

	wg := new(sync.WaitGroup)
	wg.Add(3)

	jsonFile := config.GetEnvString("FIREBASE_JSON", "firebase.json")
	client, err := fcmClient(ctx, jsonFile)
	if err != nil {
		m.Counter("Internal_Error_Total", map[string]string{"error_message": err.Error()}).Inc()
		return
	}

	models2.ConnectDB()

	// サブスク解約通知
	cancellation := &models.PushNotification{
		Namespace: "Cancellation",
		Option: models.NotificationOpt{
			MessageGen: func(s string) *models.Message {
				return &models.Message{
					Title:   "無料期間終了が近づいています！",
					Body:    "間も無く無料期間が終了します、解約をお考えの方はお忘れのないようご注意ください！",
					Badge:   0,
					Headers: map[string]string{"apns-priority": "10"},
				}
			},
			Duration: time.Second * 10,
		},
		Client:    client,
		Schedules: map[string]*models.Schedule{},
		Metrics:   m,
	}

	// クローラー作成

	cancellationCrawler := &models.NotificationCrawler{
		PushNotification: cancellation,
		Option:           models.DefaultNotificationCrawlerOpt(),
		Execute:          fetchCrawlerExecute(mode, m),
	}

	go func() {
		processCollector := prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{Namespace: namespace})
		prometheus.MustRegister(m, processCollector)
		http.Handle("/metrics", promhttp.Handler())
		_ = http.ListenAndServe(":2112", nil)
		wg.Done()
	}()

	go func() {
		cancellation.StartTimer(ctx)
		wg.Done()
	}()

	go func() {
		cancellationCrawler.StartCrawlerTimer(ctx)
		wg.Done()
	}()

	wg.Wait()
	m.Counter("Server_End_Total", map[string]string{}).Inc()
}

func fcmClient(ctx context.Context, jsonFile string) (*messaging.Client, error) {
	opt := option.WithCredentialsFile(jsonFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	return app.Messaging(ctx)
}

func fetchCrawlerExecute(mode string, metrics metrics.Metrics) func(ctx context.Context, notification *models.PushNotification) {
	// Record Record
	type Record struct {
		UserID             string
		UserSubscriptionID string
		StartedAt          time.Time
		FreeTrial          int // 単位は月
	}

	lastTime := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	proCrawlerExecute := func(ctx context.Context, notification *models.PushNotification) {
		// 完結型の場合はここでDBを読み込んで通知すべきユーザーをリストアップしてAddScheduleする
		var records []*Record
		sql := fmt.Sprintf("select user_subscriptions.user_id,user_subscriptions.user_subscription_id,user_subscriptions.started_at, subscriptions.free_trial from user_subscriptions join subscriptions on user_subscriptions.subscription_id = subscriptions.subscription_id where user_subscriptions.updated_at > '%v'", lastTime)
		if err := models2.DB.Raw(sql).Scan(&records).Error; err != nil {
			metrics.Counter("DB_Error_Total", map[string]string{"error_message": err.Error()}).Inc()
			return
		}
		lastTime = time.Now()
		for _, record := range records {
			// 1日前に通知
			timeToNotify := record.StartedAt.Add(time.Hour * 24 * time.Duration(record.FreeTrial-1))
			schedule := models.ApplyPlan(timeToNotify, record.UserID)
			notification.AddSchedule(record.UserSubscriptionID, schedule)
		}
	}

	devCrawlerExecute := func(ctx context.Context, notification *models.PushNotification) {
		// 完結型の場合はここでDBを読み込んで通知すべきユーザーをリストアップしてAddScheduleする
		token := config.GetEnvString("PUSH_TOKEN", "token")
		lastTime = time.Now()
		timeToNotify := time.Now().Add(time.Second * 10)
		schedule := models.ApplyPlan(timeToNotify, token)
		notification.AddSchedule(token, schedule)
	}

	if mode == "PRO" {
		return proCrawlerExecute
	}
	return devCrawlerExecute
}
