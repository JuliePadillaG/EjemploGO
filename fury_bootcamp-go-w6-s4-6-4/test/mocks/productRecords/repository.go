package productRecords

import (
	"context"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockProductRecordsRepository struct {
	DataMock []domain.ProductRecords
	Error    string
}

func (ms *MockProductRecordsRepository) Save(ctx context.Context, pr domain.ProductRecords) (int, error) {
	if ms.Error != "" {
		return 0, fmt.Errorf(ms.Error)
	}
	id := 1
	pr.ID = id
	ms.DataMock = append(ms.DataMock, pr)
	return int(id), nil
}

func (ms *MockProductRecordsRepository) ExistsProductRecord(ctx context.Context, id int) bool {
	if ms.Error != "" {
		return false
	}
	for _, m := range ms.DataMock {
		if id == m.ID {
			return true
		}
	}
	return false
}

func (ms *MockProductRecordsRepository) UniqueProduct(ctx context.Context, productID int) bool {
	return true
}
