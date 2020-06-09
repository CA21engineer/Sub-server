package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"fmt"
	"github.com/BambooTuna/go-server-lib/config"
	"google.golang.org/api/option"
	"time"
)

func main() {

	ctx := context.Background()

	jsonFile := config.GetEnvString("FIREBASE_JSON", "firebase.json")
	client, err := fcmClient(ctx, jsonFile)
	if err != nil {
		return
	}

	ticker := time.NewTicker(time.Hour * 24)
	defer ticker.Stop()
	send(ctx, client, time.Now())

	for {
		select {
		case t := <-ticker.C:
			send(ctx, client, t)
		}
	}

}

func fcmClient(ctx context.Context, jsonFile string) (*messaging.Client, error) {
	opt := option.WithCredentialsFile(jsonFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	return app.Messaging(ctx)
}

func send(ctx context.Context, client *messaging.Client, t time.Time) {
	println("Running...", t.String())
	r, _ := client.SendAll(ctx, fetchBatchMessage())
	fmt.Printf("SuccessCount: %d\n", r.SuccessCount)
	fmt.Printf("FailureCount: %d\n", r.FailureCount)

}

func fetchBatchMessage() []*messaging.Message {
	badge := 0
	message := &messaging.Message{
		APNS: &messaging.APNSConfig{
			Headers: map[string]string{
				"apns-priority": "10",
			},
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Alert: &messaging.ApsAlert{
						Title: "留年確定",
						Body:  "レポートの提出期限を超過したため、留年が確定しました",
					},
					Badge: &badge,
				},
			},
		},
		Token: "RegistrationToken",
	}
	return []*messaging.Message{message}
}
