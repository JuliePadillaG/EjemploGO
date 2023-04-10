package employee

import (
	"context"
	"errors"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockRepositoryEmployee struct {
	DataMock     []domain.Employee
	DatamockIBO  []domain.ReportInBO
	Err          string
	MethodCalled bool
}

func (m *MockRepositoryEmployee) Exists(ctx context.Context, cardNumberID string) bool {
	for _, value := range m.DataMock {
		if value.CardNumberID == cardNumberID {
			return true
		}
	}
	return false
}
func (m *MockRepositoryEmployee) Save(ctx context.Context, e domain.Employee) (int, error) {
	m.MethodCalled = true
	var id int = 1
	if m.Err != "" {
		return 0, errors.New(m.Err)
	}
	if len(m.DataMock) > 0 {
		id = m.DataMock[len(m.DataMock)-1].ID + 1
	}
	e.ID = id
	m.DataMock = append(m.DataMock, e)
	return id, nil
}
func (m *MockRepositoryEmployee) GetAll(ctx context.Context) ([]domain.Employee, error) {
	m.MethodCalled = true
	if m.Err != "" {
		return []domain.Employee{}, errors.New(m.Err)
	}
	return m.DataMock, nil
}
func (m *MockRepositoryEmployee) Get(ctx context.Context, id int) (domain.Employee, error) {
	m.MethodCalled = true
	for _, value := range m.DataMock {
		if value.ID == id {
			return value, nil
		}
	}
	return domain.Employee{}, errors.New("ID not found")
}
func (m *MockRepositoryEmployee) Update(ctx context.Context, e domain.Employee) error {
	m.MethodCalled = true
	for i := range m.DataMock {
		if m.DataMock[i].ID == e.ID {
			m.DataMock[i].FirstName = e.FirstName
			m.DataMock[i].LastName = e.LastName
			m.DataMock[i].WarehouseID = e.WarehouseID
			return nil
		}
	}
	return errors.New("Not updated")
}
func (m *MockRepositoryEmployee) Delete(ctx context.Context, id int) error {
	m.MethodCalled = true
	for i := range m.DataMock {
		if m.DataMock[i].ID == id {
			m.DataMock = append(m.DataMock[:i], m.DataMock[i+1:]...)
			return nil
		}
	}
	return errors.New("Not eliminated")
}

func (m *MockRepositoryEmployee) ReportInboundOrders(ctx context.Context) ([]domain.ReportInBO, error) {
	m.MethodCalled = true
	if m.Err != "" {
		return []domain.ReportInBO{}, errors.New(m.Err)
	}
	return m.DatamockIBO, nil
}

func (m *MockRepositoryEmployee) ReportInboundOrdersByID(ctx context.Context, id int) ([]domain.ReportInBO, error) {
	m.MethodCalled = true
	if m.Err != "" {
		return []domain.ReportInBO{}, errors.New(m.Err)
	}
	for _, bo := range m.DatamockIBO {
		if bo.ID == id {
			return []domain.ReportInBO{bo}, nil
		}
	}
	return []domain.ReportInBO{}, errors.New("ID not found")
}
