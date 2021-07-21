package service

import (
	"context"

	"github.com/MonsterYNH/api/v1/health"
)

type HealthService struct {
	health.UnimplementedHealthServiceServer
}

func (service *HealthService) HealthCheck(ctx context.Context, request *health.HealthRequest) (*health.HealthResponse, error) {
	return &health.HealthResponse{
		Message: "OK",
	}, nil
}
