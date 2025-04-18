package service

import (
	"context"
	"time"
)

type HealthService interface {
	CheckHealth(ctx context.Context) (*HealthStatus, error)
}

type HealthStatus struct {
	Status    string                 `json:"status"`
	Timestamp time.Time              `json:"timestamp"`
	Details   map[string]interface{} `json:"details"`
}
