package productRecords

import (
	"context"
	"errors"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockServiceProductRecords struct {
	DataMock []domain.ProductRecords
	Error    string
}

func (m *MockServiceProductRecords) Save(ctx context.Context, last_update_date string, purchase_price float64, sale_price float64, products_id int) (domain.ProductRecords, error) {
	if m.Error != "" {
		return domain.ProductRecords{}, errors.New(m.Error)
	}

	id := 1

	if len(m.DataMock) > 0 {
		id = m.DataMock[len(m.DataMock)-1].ID + 1
	}

	var newProductRecord domain.ProductRecords

	newProductRecord.ID = id
	newProductRecord.LastUpdateDate = last_update_date
	newProductRecord.PurchasePrice = purchase_price
	newProductRecord.SalePrice = sale_price
	newProductRecord.ProductID = products_id

	m.DataMock = append(m.DataMock, newProductRecord)

	return newProductRecord, nil
}
