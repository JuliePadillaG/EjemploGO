package purchase_orders

import (
	// "fmt"
	"context"
	"errors"
	"time"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Errors
var (
	ErrNotFoundPurchaseOrder = errors.New("buyer not found")
        ErrNotExistsBuyerID = errors.New("buyer_id doesn't exists")
        ErrNotExistsProductRecordsID = errors.New("product_records_id doesn't exists")
)

type Service interface {
        Save(ctx context.Context,
                orderNumber, trackingCode string,
                buyerID, productRecordID, orderStatusID int,
                orderDate *time.Time) (domain.PurchaseOrders, error)
	GetAllByBuyerID(ctx context.Context, id int) ([]domain.ReportPurchaseOrders, error)
}

type service struct{
        repository Repository
}

func NewService(r Repository) Service {
	return &service{
                repository: r,
        }
}

func (s *service) GetAllByBuyerID(ctx context.Context, buyerID int) ([]domain.ReportPurchaseOrders, error) {

        if buyerID != 0 && !s.repository.ExistsBuyersID(ctx, buyerID) {
		return nil, ErrNotExistsBuyerID
        }

        return s.repository.Get(ctx, buyerID)
}

func (s *service) Save(ctx context.Context, orderNumber, trackingCode string, buyerID, productRecordID, orderStatusID  int,  orderDate *time.Time) (domain.PurchaseOrders, error) {

        if !s.repository.ExistsBuyersID(ctx, buyerID) {
		return domain.PurchaseOrders{}, ErrNotExistsBuyerID
        }

        if !s.repository.ExistsProductRecordsID(ctx, productRecordID) {
		return domain.PurchaseOrders{}, ErrNotExistsProductRecordsID
        }

        purchaseOrders := domain.PurchaseOrders{}
        purchaseOrders.OrderNumber = orderNumber
        purchaseOrders.TrackingCode = trackingCode
        purchaseOrders.BuyerID = buyerID
        purchaseOrders.ProductRecordID = productRecordID
        purchaseOrders.OrderStatusID = orderStatusID
        purchaseOrders.OrderDate = orderDate

        id, err := s.repository.Save(ctx, purchaseOrders)
	if err != nil {
		return domain.PurchaseOrders{}, err
	}

        purchaseOrders.ID = id

        return purchaseOrders, nil
}

