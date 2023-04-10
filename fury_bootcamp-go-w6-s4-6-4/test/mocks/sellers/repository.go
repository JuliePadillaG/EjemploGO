package sellers

import (
	"context"
	"errors"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockRepository struct {
	DataMock []domain.Seller
	Error    string
}

func (m *MockRepository) GetAll(ctx context.Context) ([]domain.Seller, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil
}

func (m *MockRepository) Get(ctx context.Context, id int) (domain.Seller, error) {

	for i, elemento := range m.DataMock {
		if elemento.ID == id {
			return m.DataMock[i], nil
		}
	}

	return domain.Seller{}, fmt.Errorf(m.Error)
}

func (m *MockRepository) Exists(ctx context.Context, cid int) bool {
	for _, elemento := range m.DataMock {
		if elemento.CID == cid {
			return true
		}
	}
	return false
}

func (m *MockRepository) LocalityExists(ctx context.Context, locality int) bool {
	return true
}

func (m *MockRepository) Save(ctx context.Context, s domain.Seller) (int, error) {
	return 1, nil
}

func (m *MockRepository) Update(ctx context.Context, s domain.Seller) error {
	return nil
}

func (m *MockRepository) Delete(ctx context.Context, id int) error {

	for i, elemento := range m.DataMock {
		if elemento.ID == id {
			m.DataMock = append(m.DataMock[:i], m.DataMock[i+1:]...)
			return nil
		}
	}

	return errors.New("seller not found")
}
