package services

import (
	"context"
	"portsApi/internal/domain"
)

type PortRepository interface {
	CreateOrUpdatePort(ctx context.Context, port *domain.Port) error
	CountPorts(ctx context.Context) (int, error)
	GetPort(ctx context.Context, id string) (*domain.Port, error)
}
type PortService struct {
	repo PortRepository
}

func NewPortService(repo PortRepository) PortService {
	return PortService{
		repo: repo,
	}
}

func (s PortService) GetPort(ctx context.Context, id string) (*domain.Port, error) {
	return s.repo.GetPort(ctx, id)
}
func (s PortService) CountPorts(ctx context.Context) (int, error) {
	return s.repo.CountPorts(ctx)
}

func (s PortService) CreateOrUpdatePort(ctx context.Context, port *domain.Port) error {
	return s.repo.CreateOrUpdatePort(ctx, port)
}
