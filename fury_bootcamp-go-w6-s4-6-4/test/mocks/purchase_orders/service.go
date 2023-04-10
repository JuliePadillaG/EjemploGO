package purchase_orders

import (
	"context"
	"fmt"
	"time"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockService struct {
	DataMock                []domain.PurchaseOrders
	DataMockBuyers          []domain.Buyer
	DataMockReports         []domain.ReportPurchaseOrders
	GetAllBySellerWasCalled bool
	Error                   string
}

func (m *MockService) Save(ctx context.Context, orderNumber, trackingCode string, buyerID, productRecordID, orderStatusID int, orderDate *time.Time) (domain.PurchaseOrders, error) {

	if m.Error != "" {
		return domain.PurchaseOrders{}, fmt.Errorf(m.Error)
	}

	// m.

	var lastID int
	if len(m.DataMock) == 0 {
		lastID = 0
	} else {
		lastID = len(m.DataMock)
		// fmt.Printf("sdjfsldk %+v\n", lastID)
	}

	lastID++
	testPurchaseOrders := domain.PurchaseOrders{
		ID:              lastID,
		OrderNumber:     orderNumber,
		OrderDate:       orderDate,
		TrackingCode:    trackingCode,
		BuyerID:         buyerID,
		ProductRecordID: productRecordID,
		OrderStatusID:   orderStatusID,
	}

	m.DataMock = append(m.DataMock, testPurchaseOrders)

	return testPurchaseOrders, nil
}

func (m *MockService) GetAllByBuyerID(ctx context.Context, buyerID int) ([]domain.ReportPurchaseOrders, error) {

	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}

	if buyerID == 0 {
		return m.DataMockReports, nil
	}

	var reports []domain.ReportPurchaseOrders
	for i := range m.DataMockReports {
		if m.DataMockReports[i].ID == buyerID {
			reports = append(reports, m.DataMockReports[i])
		}
	}

	return reports, nil
}
