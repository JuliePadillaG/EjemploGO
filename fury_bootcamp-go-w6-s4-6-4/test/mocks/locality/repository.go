package locality

import (
	"context"
	"errors"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockRepository struct {
	DataMock []domain.Locality
	Error    string
}

func (m *MockRepository) Exists(ctx context.Context, id int) bool {
	for _, elemento := range m.DataMock {
		if elemento.ID == id {
			return true
		}
	}
	return false
}

func (m *MockRepository) Create(ctx context.Context, l domain.Locality) (int, error) {
	if m.Exists(ctx, l.ID) {
		return 0, errors.New("id already exists")
	}

	return int(l.ID), nil
}

func (m *MockRepository) GetAllSellersByLocality(ctx context.Context, localityID string) ([]domain.ResponseLocality, error) {
	return []domain.ResponseLocality{}, nil
}

func (m *MockRepository) GetCarriesReport(ctx context.Context, id string) (carriesReports []domain.CarriesReport, err error) {
	return []domain.CarriesReport{}, nil
}
