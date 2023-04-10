package warehouse

import (
	"context"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type mockRepositoryWarehouse struct {
	DataMock []domain.Warehouse
	GlobalId int
}

func NewRepositoryWarehouse(data []domain.Warehouse) *mockRepositoryWarehouse {
	return &mockRepositoryWarehouse{
		DataMock: data,
		GlobalId: 1,
	}
}
func (m *mockRepositoryWarehouse) GetAll(ctx context.Context) ([]domain.Warehouse, error) {
	return m.DataMock, nil
}

func (m *mockRepositoryWarehouse) Get(ctx context.Context, id int) (domain.Warehouse, error) {
	for _, w := range m.DataMock {
		if w.ID == id {
			return w, nil
		}
	}
	return domain.Warehouse{}, fmt.Errorf("warehouse not found")
}

func (m *mockRepositoryWarehouse) Exists(ctx context.Context, warehouseCode string) bool {
	for _, w := range m.DataMock {
		if w.WarehouseCode == warehouseCode {
			return true
		}
	}
	return false
}

func (m *mockRepositoryWarehouse) Save(ctx context.Context, w domain.Warehouse) (int, error) {
	if m.Exists(ctx, w.WarehouseCode) {
		return 0, fmt.Errorf("warehouse already exists")
	}
	w.ID = m.GlobalId
	m.DataMock = append(m.DataMock, w)
	m.GlobalId++
	return w.ID, nil
}

func (m *mockRepositoryWarehouse) Update(ctx context.Context, w domain.Warehouse) error {
	for i, wh := range m.DataMock {
		if wh.ID == w.ID {
			m.DataMock[i] = w
			return nil
		}
	}
	return fmt.Errorf("warehouse not found")
}

func (m *mockRepositoryWarehouse) Delete(ctx context.Context, id int) error {
	for i, w := range m.DataMock {
		if w.ID == id {
			m.DataMock = append(m.DataMock[:i], m.DataMock[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("warehouse not found")
}
