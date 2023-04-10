package locality

import (
	"context"
	"errors"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("seller not found")
	ErrRequired = errors.New("field required")
	ErrExists   = errors.New("id already exists")
	ErrRequest  = errors.New("incorrec field content")
)

type Service interface {
	Create(ctx context.Context, l domain.Locality) (domain.Locality, error)
	GetAllSellersByLocality(ctx context.Context, localityID string) ([]domain.ResponseLocality, error)
	GetCarriesReport(ctx context.Context, id string) ([]domain.CarriesReport, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Create(ctx context.Context, l domain.Locality) (domain.Locality, error) {

	_, err := s.repository.Create(ctx, l)
	if err != nil {
		return domain.Locality{}, err
	}

	return l, nil
}

func (s *service) GetAllSellersByLocality(ctx context.Context, localityID string) ([]domain.ResponseLocality, error) {
	return s.repository.GetAllSellersByLocality(ctx, localityID)
}

func (s *service) GetCarriesReport(ctx context.Context, id string) (carriesReports []domain.CarriesReport, err error) {

	carriesReports, err = s.repository.GetCarriesReport(ctx, id)
	if err != nil {
		return nil, err
	}

	return carriesReports, nil
}
