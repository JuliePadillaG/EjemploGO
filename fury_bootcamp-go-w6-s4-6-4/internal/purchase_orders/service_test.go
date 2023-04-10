package purchase_orders

import (
	// "errors"
	"context"
	"testing"
	"time"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/purchase_orders"
	"github.com/stretchr/testify/assert"
)


func TestServiceCreateOk(t *testing.T) {

        t.Run("Create a new purchase orders OK", func (t *testing.T) {
                // Arrange
                orderDate := time.Now()
                productRecords := []domain.ProductRecords {
                        {
        	        	ID: 1,
                                LastUpdateDate: "2323/10/02",
                                PurchasePrice: 4.5,
                                SalePrice: 3.4,
                                ProductID: 1,
                        },
                }

                buyers := []domain.Buyer{
                        {
        	        	ID:           1,
                                CardNumberID: "232345",
                                FirstName:    "hello",
                                LastName:     "World",
                        },
                }

                newPurchaseOrders := domain.PurchaseOrders{
                        OrderNumber: "232345",
                        OrderDate: &orderDate,
                        TrackingCode: "bar",
                        BuyerID: 1,
                        ProductRecordID: productRecords[0].ID,
                        OrderStatusID: buyers[0].ID,
                }

                mockRepository := purchase_orders.MockRepository{
                        DataMock: []domain.PurchaseOrders{},
                        DataMockBuyers: buyers,
                        DataMockProductRecords: productRecords,
                        ExistsProductRecord: false,
                        ExistsBuyer: false,
                }

                productIDExpected := 1

                // Act
                service := NewService(&mockRepository)
                ctx := context.Background()
                result, err := service.Save(ctx, 
                        newPurchaseOrders.OrderNumber, 
                        newPurchaseOrders.TrackingCode, 
                        newPurchaseOrders.BuyerID, 
                        newPurchaseOrders.ProductRecordID, 
                        newPurchaseOrders.OrderStatusID, 
                        newPurchaseOrders.OrderDate) 

                // Assert
                assert.Nil(t, err)
                assert.True(t, mockRepository.ExistsProductRecord)
                assert.True(t, mockRepository.ExistsBuyer)
                assert.Equal(t, mockRepository.DataMock[0], result)
                assert.Equal(t, productIDExpected, result.ID)
        })
        

        t.Run("Cannot create a new Purchase because buyerID not found Fail", func (t *testing.T) {

	        // Arrange
                errMessage := "buyer_id doesn't exists"
                database := []domain.PurchaseOrders {}

                orderDate := time.Now()
                newPurchaseOrders := domain.PurchaseOrders {
                        OrderNumber: "232345",
                        OrderDate: &orderDate,
                        TrackingCode: "bar",
                        BuyerID: 1,
                        ProductRecordID: 1,
                        OrderStatusID: 1,
                }

                mockRepository := purchase_orders.MockRepository{
                        DataMock: database,
                        DataMockBuyers: []domain.Buyer{},
                        Error: errMessage,
                }

                productIDExpected := 0

                // Act
                service := NewService(&mockRepository)
                ctx := context.Background()
                result, err := service.Save(ctx, 
                        newPurchaseOrders.OrderNumber, 
                        newPurchaseOrders.TrackingCode, 
                        newPurchaseOrders.BuyerID, 
                        newPurchaseOrders.ProductRecordID, 
                        newPurchaseOrders.OrderStatusID, 
                        newPurchaseOrders.OrderDate) 

                // t.Log(err)

                // Assert
	        assert.NotNil(t, err)
                assert.Equal(t, errMessage, err.Error())
	        assert.Empty(t, result)
	        assert.Equal(t, productIDExpected, result.ID)
        })

}


func TestGetAllByBuyerIDOk(t *testing.T) {

        t.Run("GetAllByBuyerIDSuccess", func (t *testing.T) {
                // Arrange
                buyers := []domain.Buyer{
                        {
        	        	ID:           1,
                                CardNumberID: "232345",
                                FirstName:    "hello",
                                LastName:     "World",
                        },
                }

                reports := []domain.ReportPurchaseOrders {
                        {
                                ID: 1,
                                CardNumberID: "232345",
                                FirstName: "foo",
                                LastName: "bar",
                                PurchaseOrdersCount: 1,
                        },
                        {
                                ID: 2,
                                CardNumberID: "232346",
                                FirstName: "foo1",
                                LastName: "bar1",
                                PurchaseOrdersCount: 2,
                        },
                }

                mockRepository := purchase_orders.MockRepository{
                        DataMockReports: reports,
                        DataMockBuyers: buyers,
                }

                // productIDExpected := 1
                buyerIDToGet := buyers[0].ID
                lengthReports := 1

                // Act
                service := NewService(&mockRepository)
                ctx := context.Background()
                result, err := service.GetAllByBuyerID(ctx, buyerIDToGet) 

                // Assert
                assert.Nil(t, err)
                assert.Equal(t, lengthReports, len(result))
        })

        t.Run("GetAllByBuyerIDFail", func (t *testing.T) {
                // Arrange
                errMessage := "buyer_id doesn't exists"
                reports := []domain.ReportPurchaseOrders {
                        {
                                ID: 1,
                                CardNumberID: "232345",
                                FirstName: "foo",
                                LastName: "bar",
                                PurchaseOrdersCount: 1,
                        },
                        {
                                ID: 2,
                                CardNumberID: "232346",
                                FirstName: "foo",
                                LastName: "bar",
                                PurchaseOrdersCount: 2,
                        },
                }

                mockRepository := purchase_orders.MockRepository{
                        DataMockReports: reports,
                        Error: errMessage,
                }

                // productIDExpected := 1
                buyerIDToGet := 4

                // Act
                service := NewService(&mockRepository)
                ctx := context.Background()
                result, err := service.GetAllByBuyerID(ctx, buyerIDToGet) 

                // Assert
                assert.NotNil(t, err)
                //assert.EqualError(err.Error(), errors.New(errMessage).Error())
                assert.NotEqual(t, len(result), len(reports))
        })
}
