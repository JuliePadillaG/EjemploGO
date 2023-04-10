package carry

import (
	"context"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type mockRepositoryCarry struct {
	DataMock         []domain.Carry
	DataMockLocality []domain.Locality
	GlobalId         int
}

func NewRepositoryCarry(data []domain.Carry) *mockRepositoryCarry {
	return &mockRepositoryCarry{
		DataMock: data,
		DataMockLocality: []domain.Locality{
			{
				ID: 1,
			},
		},
		GlobalId: 1,
	}
}

func (m *mockRepositoryCarry) Exists(ctx context.Context, carryCode string) bool {
	for _, c := range m.DataMock {
		if c.CID == carryCode {
			return true
		}
	}
	return false
}

func (m *mockRepositoryCarry) ExistsLocality(ctx context.Context, localityCode int) bool {
	for _, c := range m.DataMockLocality {
		if c.ID == localityCode {
			return true
		}
	}
	return false
}
func (m *mockRepositoryCarry) Save(ctx context.Context, carry domain.Carry) (int, error) {
	if m.Exists(ctx, carry.CID) {
		return 0, fmt.Errorf("carry with code %s already exists", carry.CID)
	}

	if !m.ExistsLocality(ctx, carry.Locality_id) {
		return 0, fmt.Errorf("locality with code %d not exists", carry.Locality_id)
	}

	carry.ID = m.GlobalId
	m.DataMock = append(m.DataMock, carry)
	m.GlobalId++

	return carry.ID, nil
}
