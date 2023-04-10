package warehouse

import (
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// type mockServiceWarehouse struct {
// 	mockRepository mockRepositoryWarehouse
// }

// func NewServiceWarehouse(mockRepo mockRepositoryWarehouse) *mockServiceWarehouse {
// 	return &mockServiceWarehouse{mockRepo}
// }

// func (m *mockServiceWarehouse) Get(id int) (domain.Warehouse, error) {
// 	return m.mockRepository.Get(context.Background(), id)
// }

// func (m *mockServiceWarehouse) GetAll() ([]domain.Warehouse, error) {
// 	return m.mockRepository.GetAll(context.Background())
// }

// func (m *mockServiceWarehouse) Save(w domain.Warehouse) (int, error) {
// 	return m.mockRepository.Save(context.Background(), w)
// }

// func (m *mockServiceWarehouse) Update(w domain.Warehouse, id int) (domain.Warehouse, error) {
// 	originalWarehouse, err := m.Get(id)
// 	if err != nil {
// 		return domain.Warehouse{}, err
// 	}
// 	if w.Address == "" {
// 		w.Address = originalWarehouse.Address
// 	}

// 	if w.Telephone == "" {
// 		w.Telephone = originalWarehouse.Telephone
// 	}

// 	if w.WarehouseCode == "" {
// 		w.WarehouseCode = originalWarehouse.WarehouseCode
// 	}

// 	if w.MinimumCapacity == nil {
// 		w.MinimumCapacity = originalWarehouse.MinimumCapacity
// 	}

// 	if w.MinimumTemperature == nil {
// 		w.MinimumTemperature = originalWarehouse.MinimumTemperature
// 	}
// 	w.ID = id

// 	return w, m.mockRepository.Update(context.Background(), w)

// }

// func (m *mockServiceWarehouse) Delete(id int) error {
// 	return m.mockRepository.Delete(context.Background(), id)
// }

type mockServiceWarehouse struct {
	dataMock []domain.Warehouse
	GlobalId int
}

func NewServiceWarehouse(data []domain.Warehouse) *mockServiceWarehouse {
	return &mockServiceWarehouse{
		dataMock: data,
		GlobalId: 1,
	}
}

func (m *mockServiceWarehouse) Get(id int) (domain.Warehouse, error) {
	for _, w := range m.dataMock {
		if w.ID == id {
			return w, nil
		}
	}
	return domain.Warehouse{}, fmt.Errorf("warehouse not found")
}

func (m *mockServiceWarehouse) GetAll() ([]domain.Warehouse, error) {
	return m.dataMock, nil
}

func (m *mockServiceWarehouse) Save(w domain.Warehouse) (int, error) {
	for _, warehouse := range m.dataMock {
		if warehouse.WarehouseCode == w.WarehouseCode {
			return 0, fmt.Errorf("warehouse code already exists")
		}
	}
	w.ID = m.GlobalId
	m.dataMock = append(m.dataMock, w)
	m.GlobalId++
	return w.ID, nil
}

func (m *mockServiceWarehouse) Update(w domain.Warehouse, id int) (domain.Warehouse, error) {
	originalWarehouse, err := m.Get(id)
	if err != nil {
		return domain.Warehouse{}, err
	}
	if w.Address == "" {
		w.Address = originalWarehouse.Address
	}

	if w.Telephone == "" {
		w.Telephone = originalWarehouse.Telephone
	}

	if w.WarehouseCode == "" {
		w.WarehouseCode = originalWarehouse.WarehouseCode
	}

	if w.MinimumCapacity == nil {
		w.MinimumCapacity = originalWarehouse.MinimumCapacity
	}

	if w.MinimumTemperature == nil {
		w.MinimumTemperature = originalWarehouse.MinimumTemperature
	}
	w.ID = id

	for i, wh := range m.dataMock {
		if wh.ID == id {
			m.dataMock[i] = w
			return w, nil
		}
	}
	return domain.Warehouse{}, fmt.Errorf("warehouse not found")
}

func (m *mockServiceWarehouse) Delete(id int) error {
	for i, w := range m.dataMock {
		if w.ID == id {
			m.dataMock = append(m.dataMock[:i], m.dataMock[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("warehouse not found")
}
