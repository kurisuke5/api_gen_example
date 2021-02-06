package health

import (
	"context"

	"golang.org/x/xerrors"
)

type Service struct {
	repository Repository
}

func NewHealthService(repo Repository) *Service {
	return &Service{repository: repo}
}

func (s *Service) Check(ctx context.Context) error {
	if err := s.repository.Check(ctx); err != nil {
		return xerrors.Errorf("unhealthy database: %w", err)
	}
	return nil
}
