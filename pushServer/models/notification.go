package models

import (
	"context"
	"firebase.google.com/go/messaging"
	"github.com/BambooTuna/go-server-lib/metrics"
	"time"
)

// PushNotification PushNotification
type PushNotification struct {
	Namespace string
	Option    NotificationOpt
	Client    *messaging.Client
	Schedules map[string]*Schedule
	Metrics   metrics.Metrics
}

// DefaultPushNotification DefaultPushNotification
func DefaultPushNotification(namespace string, client *messaging.Client, metrics metrics.Metrics) *PushNotification {
	return &PushNotification{Namespace: namespace, Option: DefaultNotificationOpt(), Client: client, Schedules: map[string]*Schedule{}, Metrics: metrics}
}

// StartTimer StartTimerは別スレッドで実行してください
func (p PushNotification) StartTimer(ctx context.Context) {
	ticker := time.NewTicker(p.Option.Duration)
	defer ticker.Stop()
	p.executeSchedule(ctx)
	for {
		select {
		case <-ticker.C:
			p.executeSchedule(ctx)
		}
	}
}

// AddSchedule AddSchedule
func (p *PushNotification) AddSchedule(id string, schedule *Schedule) {
	if !schedule.Executed() {
		p.Schedules[id] = schedule
	}
}

func (p PushNotification) executeSchedule(ctx context.Context) {
	p.Metrics.Gauge(p.Namespace+"_Queue_Size", map[string]string{}).Set(float64(len(p.Schedules)))
	for k, v := range p.Schedules {
		if v.CanExecute() {
			if err := v.Execute(ctx, p.sendMessage); err != nil {
				p.Metrics.Counter(p.Namespace+"_Error_Total", map[string]string{"error_message": err.Error()}).Inc()
			}
			delete(p.Schedules, k)
		}
	}
}

func (p PushNotification) sendMessage(ctx context.Context, token string) error {
	m := p.Option.MessageGen(token)
	message := &messaging.Message{
		APNS: &messaging.APNSConfig{
			Headers: m.Headers,
			Payload: &messaging.APNSPayload{
				Aps: &messaging.Aps{
					Alert: &messaging.ApsAlert{
						Title: m.Title,
						Body:  m.Body,
					},
					Badge: &m.Badge,
				},
			},
		},
		Token: token,
	}
	r, err := p.Client.SendAll(ctx, []*messaging.Message{message})
	p.Metrics.Counter(p.Namespace+"_Result_Total", map[string]string{"type": "Success"}).Add(float64(r.SuccessCount))
	p.Metrics.Counter(p.Namespace+"_Result_Total", map[string]string{"type": "Failure"}).Add(float64(r.FailureCount))
	return err
}
