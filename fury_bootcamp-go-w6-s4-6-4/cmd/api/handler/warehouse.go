package handler

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/warehouse"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/pkg/web"
)

type Warehouse struct {
	warehouseService warehouse.Service
}

func NewWarehouse(w warehouse.Service) *Warehouse {
	return &Warehouse{
		warehouseService: w,
	}
}

// ListOneWarehouse godoc
// @Summary List one warehouse
// @Tags Warehouses
// @Description get product by ID
// @Accept  json
// @Produce  json
// @Success 200 {object} web.response
// @Router /api/v1/warehouses/:id [get]
// @Param id path int true "Warehouse ID"
func (w *Warehouse) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		warehouse, err := w.warehouseService.Get(id)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}
		web.Success(c, http.StatusOK, warehouse)
	}
}

// ListWarehouses godoc
// @Summary List warehouses
// @Tags Warehouses
// @Description get warehouses
// @Accept  json
// @Produce  json
// @Success 200 {object} web.response
// @Router /api/v1/warehouses [get]
func (w *Warehouse) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouses, err := w.warehouseService.GetAll()
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}
		web.Success(c, http.StatusOK, warehouses)
	}
}

// CreateWarehouse godoc
// @Summary Create warehouse
// @Tags Warehouses
// @Description create warehouse
// @Accept  json
// @Produce  json
// @Param warehouse body domain.Warehouse true "Warehouse"
// @Success 201 {object} web.response
// @Router /api/v1/warehouses [post]
func (w *Warehouse) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		warehouse := domain.Warehouse{}
		if err := c.ShouldBindJSON(&warehouse); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		emptyFields := []string{}
		values := reflect.ValueOf(warehouse)
		for i := 0; i < values.NumField(); i++ {
			if values.Type().Field(i).Name == "ID" {
				continue
			}
			if values.Field(i).IsZero() {
				emptyFields = append(emptyFields, values.Type().Field(i).Name)
			}
		}
		if len(emptyFields) > 0 {
			web.Error(c, http.StatusUnprocessableEntity, "empty fields: "+strings.Join(emptyFields, ", "))
			return
		}

		if *warehouse.MinimumCapacity < 0 {
			web.Error(c, http.StatusUnprocessableEntity, "minimum capacity must be greater than 0")
			return
		}

		if *warehouse.MinimumTemperature > 20 || *warehouse.MinimumTemperature < -10 {
			web.Error(c, http.StatusUnprocessableEntity, "minimum temperature must be between -10 and 20")
			return
		}

		id, err := w.warehouseService.Save(warehouse)
		if err != nil {
			web.Error(c, http.StatusConflict, err.Error())
			return
		}

		warehouse.ID = id

		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		web.Success(c, http.StatusCreated, warehouse)
	}
}

// UpdateWarehouse godoc
// @Summary Update warehouse
// @Tags Warehouses
// @Description update warehouse
// @Accept  json
// @Produce  json
// @Param id path int true "Warehouse ID"
// @Success 200 {object} web.response
// @Router /api/v1/warehouses/:id [patch]
func (w *Warehouse) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		warehouse := domain.Warehouse{}

		if err := c.ShouldBindJSON(&warehouse); err != nil {
			web.Error(c, http.StatusUnprocessableEntity, err.Error())
			return
		}

		if warehouse.MinimumCapacity != nil {
			if *warehouse.MinimumCapacity < 0 {
				web.Error(c, http.StatusUnprocessableEntity, "minimum capacity must be greater than 0")
				return
			}
		}
		if warehouse.MinimumTemperature != nil {
			if *warehouse.MinimumTemperature > 20 || *warehouse.MinimumTemperature < -10 {
				web.Error(c, http.StatusUnprocessableEntity, "minimum temperature must be between -10 and 20")
				return
			}
		}

		updateWarehouse, err := w.warehouseService.Update(warehouse, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web.Success(c, http.StatusOK, updateWarehouse)
	}
}

// DeleteWarehouse godoc
// @Summary Delete warehouse
// @Tags Warehouses
// @Description delete warehouse
// @Accept  json
// @Produce  json
// @Param id path int true "Warehouse ID"
// @Success 204 {object} web.response
// @Router /api/v1/warehouses/:id [delete]
func (w *Warehouse) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := w.warehouseService.Delete(id); err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web.Success(c, http.StatusNoContent, gin.H{"message": "warehouse deleted"})
	}
}
