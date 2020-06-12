package models

import (
	"context"
	"errors"
	"time"
)

// Schedule Schedule
type Schedule struct {
	scheduledTime time.Time // Scheduled execution time
	destination   string    // Push Token
	executed      bool      // is executed
}

// ApplyPlan ApplyPlan
func ApplyPlan(t time.Time, pushToken string) *Schedule {
	return &Schedule{scheduledTime: t, destination: pushToken, executed: false}
}

// Execute Execute
func (s Schedule) Execute(ctx context.Context, f func(context.Context, string) error) error {
	if s.executed {
		return errors.New("executed")
	}
	return f(ctx, s.destination)
}

// CanExecute CanExecute
func (s Schedule) CanExecute() bool {
	return s.scheduledTime.Before(time.Now())
}

// Executed Executed
func (s Schedule) Executed() bool {
	return s.executed
}
