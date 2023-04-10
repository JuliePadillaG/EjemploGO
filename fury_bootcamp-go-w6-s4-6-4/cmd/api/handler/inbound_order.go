package handler

import (
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	inboundorder "github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/inbound_order"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/pkg/web"
)

type request_Inbound_Order struct {
	Order_date       string `json:"order_date"`
	Order_number     string `json:"order_number"`
	Employee_id      int    `json:"employee_id"`
	Product_batch_id int    `json:"product_batch_id"`
	Warehouse_id     int    `json:"warehouse_id"`
}

type Inbound_order struct {
	inbound_ordersService inboundorder.Service
}

func NewInBound_Order(bo inboundorder.Service) *Inbound_order {
	return &Inbound_order{
		inbound_ordersService: bo,
	}
}

func (bo *Inbound_order) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		inBOs, err := bo.inbound_ordersService.GetAll_inboundOrders(ctx)
		if err != nil {
			web.Error(ctx, 404, "%s", err)
			return
		}
		if len(inBOs) == 0 {
			web.Success(ctx, 200, "No existing Inbound_orders")
			return
		}
		web.Success(ctx, 200, inBOs)
	}
}

func (bo *Inbound_order) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req request_Inbound_Order

		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Error(ctx, 422, "%s", err)
			return
		}
		var emptyFiled []string
		values := reflect.ValueOf(req)
		for i := 0; i < values.NumField(); i++ {
			if values.Field(i).Interface() == reflect.Zero(values.Field(i).Type()).Interface() {
				emptyFiled = append(emptyFiled, values.Type().Field(i).Name)
			}
		}

		if len(emptyFiled) > 0 {
			for _, v := range emptyFiled {
				if v != "ID" {
					web.Error(ctx, 422, "El campo: %s es requerido", v)
				}
			}
			return
		}

		_, err := time.Parse("2006-01-02", req.Order_date)
		if err != nil {
			web.Error(ctx, 422, "invalid date example 2008-01-02")
			return
		}

		inbOrder, err := bo.inbound_ordersService.Save(ctx, req.Order_date, req.Order_number, req.Employee_id, req.Product_batch_id, req.Warehouse_id)
		if err != nil {
			web.Error(ctx, 409, "%s", err)
			return
		}
		web.Success(ctx, 201, inbOrder)
	}
}
