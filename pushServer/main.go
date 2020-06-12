package main

import (
	"context"
	"github.com/BambooTuna/go-server-lib/config"
	"github.com/BambooTuna/go-server-lib/metrics"
	"github.com/CA21engineer/Subs-server/pushServer/models"
	"sync"
	"time"
)

const namespace = "go_push_server"

func main() {
	ctx := context.Background()
	m := metrics.CreateMetrics(namespace)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	jsonFile := config.GetEnvString("FIREBASE_JSON", "firebase.json")
	client, err := FcmClient(ctx, jsonFile)
	if err != nil {
		m.Counter("Internal_Error_Total", map[string]string{"error_message": err.Error()}).Inc()
		return
	}

	// サブスク解約通知
	cancellation := models.DefaultPushNotification("Cancellation", client, m)

	// クローラー作成
	cancellationCrawler := models.NotificationCrawler{
		PushNotification: cancellation,
		Option:           models.DefaultNotificationCrawlerOpt(),
		Execute: func(ctx context.Context, notification *models.PushNotification) {
			schedule := models.ApplyPlan(time.Now().Add(time.Second*30), "push_token")
			notification.AddSchedule(schedule)
		},
	}

	go func() {
		cancellation.StartTimer(ctx)
		wg.Done()
	}()

	go func() {
		cancellationCrawler.StartCrawlerTimer(ctx)
		wg.Done()
	}()

}
