package carry

import (
	"context"
	"errors"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type Service interface {
	Save(c domain.Carry) (int, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Save(c domain.Carry) (int, error) {
	if s.repository.Exists(context.Background(), c.CID) {
		return 0, errors.New("carry code already exists")
	}

	if !s.repository.ExistsLocality(context.Background(), c.Locality_id) {
		return 0, errors.New("locality code doesn't exists")
	}

	return s.repository.Save(context.Background(), c)
}
