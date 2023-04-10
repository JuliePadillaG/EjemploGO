package inboundorder

import (
	"context"
	"errors"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockRepositoryIBO struct {
	DataMock     []domain.Inbound_order
	DataMockEmp  []domain.Employee
	Err          string
	MethodCalled bool
}

func (m *MockRepositoryIBO) GetAll(ctx context.Context) ([]domain.Inbound_order, error) {
	m.MethodCalled = true
	if m.Err != "" {
		return []domain.Inbound_order{}, errors.New(m.Err)
	}
	return m.DataMock, nil
}
func (m *MockRepositoryIBO) ExistsEmployee(ctx context.Context, id_employee int) bool {
	m.MethodCalled = true
	for _, value := range m.DataMockEmp {
		if value.ID == id_employee {
			return true
		}
	}
	return false
}
func (m *MockRepositoryIBO) ExistsInboundOrder(ctx context.Context, order_number string) bool {
	m.MethodCalled = true
	for _, value := range m.DataMock {
		if value.Order_number == order_number {
			return true
		}
	}
	return false
}

func (m *MockRepositoryIBO) Save(ctx context.Context, b_order domain.Inbound_order) (int, error) {
	m.MethodCalled = true
	var id int = 1
	if m.Err != "" {
		return 0, errors.New(m.Err)
	}
	if len(m.DataMock) > 0 {
		id = m.DataMock[len(m.DataMock)-1].ID + 1
	}
	b_order.ID = id
	m.DataMock = append(m.DataMock, b_order)
	return id, nil
}
