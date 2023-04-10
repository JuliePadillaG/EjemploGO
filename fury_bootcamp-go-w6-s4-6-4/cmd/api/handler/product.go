package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/product"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/pkg/web"
)

type requestProduct struct {
	Description    string   `json:"description"`
	ExpirationRate *int     `json:"expiration_rate"`
	FreezingRate   *int     `json:"freezing_rate"`
	Height         *float32 `json:"height"`
	Length         *float32 `json:"length"`
	Netweight      *float32 `json:"netweight"`
	ProductCode    string   `json:"product_code"`
	RecomFreezTemp *float32 `json:"recommended_freezing_temperature"`
	Width          *float32 `json:"width"`
	ProductTypeID  *int     `json:"product_type_id"`
	SellerID       *int     `json:"seller_id"`
}

// Se debe generar la estructura del controlador que tenga como campo el servicio
type Product struct {
	productService product.Service
}

// Se debe generar la función que retorne el controlador
func NewProduct(w product.Service) *Product {
	return &Product{
		productService: w,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Success 200 {object} web.response
// @Router /api/v1/products [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := p.productService.GetAll(c)
		if err != nil {
			web.Error(c, 404, "%s", err)
			return
		}
		if len(products) == 0 {
			web.Success(c, 200, "no existing products")
			return
		}
		web.Success(c, 200, products)
	}
}

// ListOneProduct godoc
// @Summary List one product
// @Tags Products
// @Description get product by ID
// @Accept  json
// @Produce  json
// @Success 200 {object} web.response
// @Router /api/v1/products/:id [get]
// @Param id path int true "Product ID"
func (p *Product) Get() gin.HandlerFunc {
	return func(c *gin.Context) {

		var pr domain.Product

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "%s", err)
			return
		}

		pr, err = p.productService.Get(c, int(id))
		if err != nil {
			web.Error(c, http.StatusNotFound, "%s", err)
			return
		}

		web.Success(c, http.StatusOK, pr)
	}
}

// CreateProduct godoc
// @Summary Create a product
// @Tags Products
// @Description create a product
// @Accept  json
// @Produce  json
// @Param carry body domain.Product true "Product"
// @Success 201 {object} web.response
// @Router /api/v1/products [post]
func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req requestProduct

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusBadRequest, "%s", err)
			return
		}

		if req.Description == "" {
			web.Error(c, http.StatusUnprocessableEntity, "%s", "description is required")
			return
		}
		if req.ExpirationRate == nil {
			web.Error(c, http.StatusUnprocessableEntity, "%s", "expiration_rate is required")
			return
		}
		if req.FreezingRate == nil {
			web.Error(c, http.StatusUnprocessableEntity, "%s", "freezing_rate is required")
			return
		}
		if req.Height == nil {
			web.Error(c, http.StatusUnprocessableEntity, "%s", "height is required")
			return
		}
		if req.Length == nil {
			web.Error(c, http.StatusUnprocessableEntity, "%s", "length is required")
			return
		}
		if req.Netweight == nil {
			web.Error(c, http.StatusUnprocessableEntity, "%s", "netweight is required")
			return
		}
		if req.ProductCode == "" {
			web.Error(c, http.StatusUnprocessableEntity, "%s", "product code is required")
			return
		}
		if req.RecomFreezTemp == nil {
			web.Error(c, http.StatusUnprocessableEntity, "%s", "recommended_freezing_temperature is required")
			return
		}
		if req.Width == nil {
			web.Error(c, http.StatusUnprocessableEntity, "%s", "width is required")
			return
		}
		if req.ProductTypeID == nil {
			web.Error(c, http.StatusUnprocessableEntity, "%s", "product_type_id is required")
			return
		}
		if req.SellerID == nil {
			web.Error(c, http.StatusUnprocessableEntity, "%s", "seller_id is required")
			return
		}
		
		pr, err := p.productService.Save(c, req.Description, *req.ExpirationRate, *req.FreezingRate, *req.Height, *req.Length, *req.Netweight, req.ProductCode, *req.RecomFreezTemp, *req.Width, *req.ProductTypeID, *req.SellerID)
		if err != nil {
			if err.Error() == "product_code already exists" {
				web.Error(c, 409, "%s", err)
			} else {
				web.Error(c, 500, "%s", err)
			}
			return
		}
		web.Success(c, 201, pr)
	}
}

// UpdateProduct godoc
// @Summary Update product
// @Tags Products
// @Description update product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 200 {object} web.response
// @Router /api/v1/products/:id [patch]
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "%s", err)
			return
		}

		var req requestProduct

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusBadRequest, "%s", err)
			return
		}

		pr, err := p.productService.Update(c, int(id), req.Description, req.ExpirationRate, req.FreezingRate, req.Height, req.Length, req.Netweight, req.ProductCode, req.RecomFreezTemp, req.Width, req.ProductTypeID, req.SellerID)
		if err != nil {
			web.Error(c, http.StatusNotFound, "%s", err)
			return
		}

		web.Success(c, http.StatusOK, pr)
	}
}

// DeleteProduct godoc
// @Summary Delete product
// @Tags Products
// @Description delete product
// @Accept  json
// @Produce  json
// @Param id path int true "Product ID"
// @Success 204 {object} web.response
// @Router /api/v1/products/:id [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, http.StatusBadRequest, "%s", err)
			return
		}

		err = p.productService.Delete(c, int(id))
		if err != nil {
			if err.Error() == "product not found" {
				web.Error(c, http.StatusNotFound, "%s", err)
			} else {
				web.Error(c, 500, "%s", err)
			}
			return
		}

		web.Success(c, http.StatusNoContent, "El producto ha sido eliminado exitosamente")
	}
}

// GetReportRecords godoc
// @Summary Get report records
// @Tags Products
// @Description get report records
// @Accept  json
// @Produce  json
// @Param id query string true "Product id"
// @Success 200 {object} web.response
// @Router /api/v1/products/reportRecords [get]
func (p *Product) GetReportRecords() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := ctx.Query("id")

		reports, err := p.productService.GetProductRecords(ctx, id)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, err.Error())
			return
		}

		web.Success(ctx, http.StatusOK, reports)
	}
}