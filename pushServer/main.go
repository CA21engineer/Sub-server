package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/BambooTuna/go-server-lib/config"
	"google.golang.org/api/option"
)

func main() {

	ctx := context.Background()

	jsonFile := config.GetEnvString("FIREBASE_JSON", "firebase.json")
	client, err := fcmClient(ctx, jsonFile)
	if err != nil {
		return
	}

	res, err := send(ctx, client)
	println(res)

}

func fcmClient(ctx context.Context, jsonFile string) (*messaging.Client, error) {
	opt := option.WithCredentialsFile(jsonFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	return app.Messaging(ctx)
}

func send(ctx context.Context, client *messaging.Client) (string, error) {
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
	r, _ := client.SendAll(ctx, []*messaging.Message{message})
	println(r.SuccessCount)
	println(r.FailureCount)
	println(r.Responses)
	return client.Send(ctx, message)
}
