package carry

import (
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type mockServiceCarry struct {
	dataMock         []domain.Carry
	dataMockLocality []domain.Locality
	GlobalId         int
}

func NewServiceCarry(data []domain.Carry) *mockServiceCarry {
	return &mockServiceCarry{
		dataMock: data,
		dataMockLocality: []domain.Locality{
			{
				ID: 1,
			},
		},
		GlobalId: 1,
	}
}

func (m *mockServiceCarry) Save(carry domain.Carry) (int, error) {
	for _, c := range m.dataMock {
		if c.CID == carry.CID {
			return 0, fmt.Errorf("carry with code %s already exists", carry.CID)
		}
	}

	exists := false
	for _, c := range m.dataMockLocality {
		if c.ID == carry.Locality_id {
			exists = true
			break
		}

	}
	if !exists {
		return 0, fmt.Errorf("locality with code %d not exists", carry.Locality_id)
	}
	carry.ID = m.GlobalId
	m.dataMock = append(m.dataMock, carry)
	m.GlobalId++

	return carry.ID, nil
}
