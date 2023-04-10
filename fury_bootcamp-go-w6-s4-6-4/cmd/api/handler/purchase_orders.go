package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

	// "strconv"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/purchase_orders"
	custom "github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/pkg/custom_datatypes"

	// "github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/pkg/web"
)

const (
        BEGIN_DATE = "1000-01-01T00:00:00+00:00"
)

var (
        ErrNotFoundPurchaseOrders = errors.New("there is no purchase orders")
        ErrDateTime = errors.New("the date is not a date of the dates range")
)

type RequestPurchaseOrders struct {
        ID              int             `json:"id"`
	OrderNumber     string          `json:"order_number"            binding:"required"`
	OrderDate       custom.MyTime   `json:"order_date"              binding:"required"`
	TrackingCode    string          `json:"tracking_code"           binding:"required"`
	BuyerID         int             `json:"buyer_id"                binding:"required"`
	ProductRecordID int             `json:"product_record_id"       binding:"required"`
	OrderStatusID   int             `json:"order_status_id"`
}


type RequestQueryPurchaseOrders struct {
        ID      int     `json:"id"      form:"id"`
}

type PurchaseOrders struct {
	purchaseOrderService purchase_orders.Service
}

func NewPurchaseOrders(p purchase_orders.Service) *PurchaseOrders {
	return &PurchaseOrders{
		purchaseOrderService: p,
	}
}

func (p *PurchaseOrders) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
                var req RequestPurchaseOrders

                if err := ctx.ShouldBindJSON(&req); err != nil {
                        log.Print(err)
                        web.Error(ctx, http.StatusUnprocessableEntity, err.Error())
                        return
		}

                previosTime := time.Time(req.OrderDate)

                begin, err := time.Parse(time.RFC3339, BEGIN_DATE)
                if err != nil {
                    panic(err)
                }

                if previosTime.After(begin) == false {
                        web.Error(ctx, http.StatusBadRequest, ErrDateTime.Error())
                        return
                }

		purchaseOrder, err := p.purchaseOrderService.Save(ctx, req.OrderNumber, req.TrackingCode,  req.BuyerID, req.ProductRecordID, req.OrderStatusID, &previosTime)
		if err != nil {
                        web.Error(ctx, http.StatusConflict, err.Error())
			return
		}

                web.Success(ctx, http.StatusCreated, purchaseOrder)
        }
}

func (p *PurchaseOrders) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
	        var req RequestQueryPurchaseOrders

                if err := ctx.ShouldBindQuery(&req); err != nil {
                        web.Error(ctx, http.StatusUnprocessableEntity, err.Error())
                        return
                }

                reportPurchaseOrders, err := p.purchaseOrderService.GetAllByBuyerID(ctx, req.ID)
                if err != nil {
                        web.Error(ctx, http.StatusNotFound, ErrNotFoundPurchaseOrders.Error())
                        return
                }

	        if len(reportPurchaseOrders) == 0 {
                        web.Error(ctx, http.StatusNotFound, ErrNotFoundPurchaseOrders.Error())
	        	return
	        }

                web.Success(ctx, http.StatusOK, reportPurchaseOrders)
        }
}

