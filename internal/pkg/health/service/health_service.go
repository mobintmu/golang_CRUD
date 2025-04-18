package service

import (
	"context"
	"time"
)

type healthService struct {
	startTime time.Time
}

func NewHealthService() HealthService {
	return &healthService{
		startTime: time.Now(),
	}
}

func (s *healthService) CheckHealth(ctx context.Context) (*HealthStatus, error) {
	status := &HealthStatus{
		Status:    "UP",
		Timestamp: time.Now(),
		Details:   make(map[string]interface{}),
	}
	return status, nil
}
