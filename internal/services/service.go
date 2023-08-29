package services

import (
	"context"
	"portsApi/internal/domain"
	"github.com/google/uuid"
)

type PortService struct {
}

func NewPortService() PortService {
	return PortService{}
}

func (s PortService) GetPort(ctx context.Context, id string) (*domain.Port, error) {
	randomID := uuid.New().String()
	return domain.NewPort(randomID, randomID, randomID, randomID, randomID,
		 []string{randomID}, []string{randomID}, []float64{1.0, 2.0}, randomID, randomID,  nil)
}