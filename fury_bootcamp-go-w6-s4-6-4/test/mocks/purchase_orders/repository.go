package purchase_orders

import (
	// "time"
	"context"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockRepository struct {
	DataMock               []domain.PurchaseOrders
	DataMockReports        []domain.ReportPurchaseOrders
	DataMockBuyers         []domain.Buyer
	DataMockProductRecords []domain.ProductRecords
	ExistsProductRecord    bool
	ExistsBuyer            bool
	Error                  string
}

func (m *MockRepository) Save(ctx context.Context, p domain.PurchaseOrders) (int, error) {

	if m.Error != "" {
		return 0, fmt.Errorf(m.Error)
	}

	var lastID int
	if len(m.DataMock) == 0 {
		lastID = 0
	} else {
		lastID = m.DataMock[len(m.DataMock)-1].ID
	}

	m.ExistsProductRecord = true
	m.ExistsBuyer = true

	lastID++
	testPurchaseOrders := domain.PurchaseOrders{
		ID:              lastID,
		OrderNumber:     p.OrderNumber,
		OrderDate:       p.OrderDate,
		TrackingCode:    p.TrackingCode,
		BuyerID:         p.BuyerID,
		ProductRecordID: p.ProductRecordID,
		OrderStatusID:   p.OrderStatusID,
	}

	m.DataMock = append(m.DataMock, testPurchaseOrders)

	return testPurchaseOrders.ID, nil
}

func (m *MockRepository) Get(ctx context.Context, id int) ([]domain.ReportPurchaseOrders, error) {

	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}

	var reports []domain.ReportPurchaseOrders
	for i := range m.DataMockReports {
		if m.DataMockReports[i].ID == id {
			reports = append(reports, m.DataMockReports[i])
		}
	}

	return reports, nil
}

func (m *MockRepository) ExistsBuyersID(ctx context.Context, buyerID int) bool {
	exists := false
	m.ExistsBuyer = true

	for i := range m.DataMockBuyers {
		if m.DataMockBuyers[i].ID == buyerID {
			exists = true
		}
	}

	return exists
}

func (m *MockRepository) ExistsProductRecordsID(ctx context.Context, productRecordID int) bool {
	exists := false
	m.ExistsProductRecord = true

	for i := range m.DataMockProductRecords {
		if m.DataMockProductRecords[i].ID == productRecordID {
			exists = true
		}
	}

	return exists
}
