package locality

import (
	"context"
	"errors"
	"strconv"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("seller not found")
	ErrRequired = errors.New("field required")
	ErrExists   = errors.New("id already exists")
	ErrRequest  = errors.New("incorrec field content")
)

type MockService struct {
	DataMock []domain.Locality
	Error    string
}

func (m *MockService) Create(ctx context.Context, l domain.Locality) (domain.Locality, error) {
	for _, element := range m.DataMock {
		if element.ID == l.ID {
			return domain.Locality{}, ErrExists
		}
	}

	return l, nil
}

func (m *MockService) GetAllSellersByLocality(ctx context.Context, localityID string) ([]domain.ResponseLocality, error) {
	id, _ := strconv.Atoi(localityID)
	for _, element := range m.DataMock {
		if element.ID == id {
			return []domain.ResponseLocality{}, ErrExists
		}
	}

	return []domain.ResponseLocality{}, nil
}

func (m *MockService) GetCarriesReport(ctx context.Context, id string) ([]domain.CarriesReport, error) {
	i, _ := strconv.Atoi(id)
	for _, element := range m.DataMock {
		if element.ID == i {
			return []domain.CarriesReport{}, ErrExists
		}
	}

	return []domain.CarriesReport{}, nil
}
