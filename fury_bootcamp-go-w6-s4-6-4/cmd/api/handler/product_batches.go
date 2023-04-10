package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	productbatches "github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/product_batches"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/pkg/web"
)

type Product_batches struct {
	ID                 int    `json:"id"`
	BatchNumber        int    `json:"section_number"`
	CurrentQuantity    int    `json:"current_quantity"`
	CurrentTemperature int    `json:"current_temperature"`
	DueDate            string `json:"due_date"`
	InitialQuantity    int    `json:"initial_quantity"`
	ManufacturingDate  string `json:"manufacturing_date"`
	ManufacturingHour  int    `json:"manufacturing_hour"`
	MinimumTemperature int    `json:"minimum_temperature"`
	ProductId          int    `json:"product_id"`
	SectionId          int    `json:"section_id"`
}

type ProductBatches struct {
	productBatchesService productbatches.Service
}

func NewProductBatches(pb productbatches.Service) *ProductBatches {
	return &ProductBatches{
		productBatchesService: pb,
	}
}

// CreateProductBatches godoc
// @Summary Create product batches
// @Tags Product_batches
// @Description create product_batches
// @Accept  json
// @Produce  json
// @Param product_batches body domain.Product_batches true "product_batches"
// @Success 201 {object} web.response
// @Router /api/v1/productbatches [post]
func (pb *ProductBatches) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req domain.Product_batches

		//2001-12-25
		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", err)
			return
		}

		err := ValidateFormatDate(req.DueDate)
		if err != nil {
			web.Error(ctx, http.StatusConflict, "%s", err)
			return
		}

		err = ValidateFormatDate(req.ManufacturingDate)
		if err != nil {
			web.Error(ctx, http.StatusConflict, "%s", err)
			return
		}

		id, err := pb.productBatchesService.CreatePB(ctx, req)
		if err != nil {
			web.Error(ctx, http.StatusConflict, "%s", err)
			return
		}
		req.ID = id
		log.Println("Soy la request")
		log.Println(req)
		web.Success(ctx, http.StatusCreated, req)
	}
}

// GetReportProduct godoc
// @Summary Get ReportProduct
// @Tags ReportProduct
// @Description get report product
// @Accept  json
// @Produce  json
// @Param id query int false "BatchNumber ID to get report (optional)"
// @Success 200 {object} web.response
// @Router /api/v1/reportProducts/ [get]
func (pb *ProductBatches) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, "%s", err)
			return
		}
		data, err := pb.productBatchesService.ReadPB(ctx, idInt)
		if err != nil {
			web.Error(ctx, http.StatusNotFound, "%s", err)
			return
		}
		web.Success(ctx, http.StatusOK, data)
	}
}

func ValidateFormatDate(date string) error {

	_, err := time.Parse("2006-01-02", date)
	if err != nil {

		return errors.New("invalid date example 2008-01-02")
	}

	return nil
}
