package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/product_records"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/pkg/web"
)

type requestProductRecords struct {
	LastUpdateDate    string   		`json:"last_update_date"`
	PurchasePrice 	  *float64     	`json:"purchase_price"`
	SalePrice   	  *float64    	`json:"sale_price"`
	ProductID         *int 			`json:"products_id"`
}

type ProductRecords struct {
	productRecordsService product_records.Service
}

func NewProductRecord(service product_records.Service) *ProductRecords {
	return &ProductRecords{
		productRecordsService: service,
	}
}

// CreateProductRecord godoc
// @Summary Create a product record
// @Tags Product Records
// @Description create a product record
// @Accept  json
// @Produce  json
// @Param carry body domain.ProductRecords true "ProductRecords"
// @Success 201 {object} web.response
// @Router /api/v1/productRecords [post]
func (pr *ProductRecords) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var req_product_records requestProductRecords

		if err := ctx.ShouldBindJSON(&req_product_records); err != nil {
			web.Error(ctx, http.StatusBadRequest, err.Error())
			return
		}
		if req_product_records.LastUpdateDate == "" {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", "last_update_date is required")
			return
		}
		if req_product_records.PurchasePrice == nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", "purchase_price is required")
			return
		}
		if req_product_records.SalePrice == nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", "sale_price is required")
			return
		}
		if req_product_records.ProductID == nil {
			web.Error(ctx, http.StatusUnprocessableEntity, "%s", "products_id is required")
			return
		}

		product_records, err := pr.productRecordsService.Save(ctx, req_product_records.LastUpdateDate, *req_product_records.PurchasePrice, *req_product_records.SalePrice, *req_product_records.ProductID)
		if err != nil {
			if err.Error() == "error: product id doesn't exists" {
				web.Error(ctx, http.StatusConflict, "%s", err)
			} else {
				web.Error(ctx, http.StatusInternalServerError, "%s", err)
			}
		}
		
		web.Success(ctx, http.StatusOK, product_records)
	}
}