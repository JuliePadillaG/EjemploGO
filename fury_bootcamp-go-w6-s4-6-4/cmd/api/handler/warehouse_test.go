package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/warehouse"
)

var s = createService()

func createService() *gin.Engine {
	r := gin.Default()
	service := warehouse.NewServiceWarehouse([]domain.Warehouse{})
	hanler := NewWarehouse(service)
	r.GET("/warehouses", hanler.GetAll())
	r.GET("/warehouses/:id", hanler.Get())
	r.POST("/warehouses", hanler.Create())
	r.PATCH("/warehouses/:id", hanler.Update())
	r.DELETE("/warehouses/:id", hanler.Delete())
	return r
}

func createRequestTestWarehouse(method string, path string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, path, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestCreate(t *testing.T) {
	t.Run("should return 201 when create a warehouse", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodPost, "/warehouses", `{"address": "test", "telephone": "test", "warehouse_code": "test", "minimum_capacity": 1, "minimum_temperature": 1}`)

		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusCreated, w.Code)
	})
	t.Run("should return 422 when create a warehouse with invalid body", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodPost, "/warehouses", `{"telephone": "test", "warehouse_code": "test", "minimum_capacity": 1, "minimum_temperature": 1}`)

		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})
	t.Run("should return 422 when create a warehouse with invalid mininum capacity", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodPost, "/warehouses", `{"address": "test", "telephone": "test", "warehouse_code": "test", "minimum_capacity": -1, "minimum_temperature": 1}`)
		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})
	t.Run("should return 422 when create a warehouse with invalid mininum temperature", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodPost, "/warehouses", `{"address": "test", "telephone": "test", "warehouse_code": "test", "minimum_capacity": 1, "minimum_temperature": 30}`)
		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})
	t.Run("should return 409 when create a warehouse with conflict warehouse code", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodPost, "/warehouses", `{"address": "test", "telephone": "test", "warehouse_code": "test", "minimum_capacity": 1, "minimum_temperature": 1}`)

		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusConflict, w.Code)
	})
}

func TestRead(t *testing.T) {
	req, w := createRequestTestWarehouse(http.MethodPost, "/warehouses", `{"address": "test", "telephone": "test", "warehouse_code": "test", "minimum_capacity": 1, "minimum_temperature": 1}`)
	s.ServeHTTP(w, req)
	t.Run("should return 200 when get a warehouse", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodGet, "/warehouses/1", "")

		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("should return 404 when get a warehouse with invalid id", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodGet, "/warehouses/2", "")

		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("should return 400 when get a warehouse without query param", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodGet, "/warehouses/ok", "")

		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	t.Run("should return 200 when get all warehouses", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodGet, "/warehouses", "")

		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestUpdate(t *testing.T) {
	req, w := createRequestTestWarehouse(http.MethodPost, "/warehouses", `{"address": "test", "telephone": "test", "warehouse_code": "test", "minimum_capacity": 1, "minimum_temperature": 1}`)
	s.ServeHTTP(w, req)
	t.Run("should return 200 when update a warehouse", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodPatch, "/warehouses/1", `{"address": "new test", "telephone": "new test", "warehouse_code": "new test"}`)

		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("should return 422 when update a warehouse with invalid mininum capacity", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodPatch, "/warehouses/1", `{"address": "new test", "telephone": "new test", "warehouse_code": "new test", " minimum_capacity": -1, "
		minimum_temperature": 1}`)
		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	})
	t.Run("should return 404 when update a warehouse with invalid mininum temperature", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodPatch, "/warehouses/1", `{"address": "new test", "telephone": "new test", "warehouse_code": "new test", "
		minimum_capacity": 1, "minimum_temperature": 30}`)
		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})
	t.Run("should return 404 when update a warehouse with invalid id", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodPatch, "/warehouses/2", `{"address": "test", "telephone": "test", "warehouse_code": "test", "minimum_capacity": 1, "minimum_temperature": 1}`)

		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestDelete(t *testing.T) {
	req, w := createRequestTestWarehouse(http.MethodPost, "/warehouses", `{"address": "test", "telephone": "test", "warehouse_code": "test", "minimum_capacity": 1, "minimum_temperature": 1}`)
	s.ServeHTTP(w, req)
	t.Run("should return 204 when delete a warehouse", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodDelete, "/warehouses/1", "")

		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusNoContent, w.Code)
	})
	t.Run("should return 404 when delete a warehouse with invalid id", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodDelete, "/warehouses/10", "")

		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusNotFound, w.Code)
	})
	t.Run("should return 400 when delete with no query param", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestWarehouse(http.MethodDelete, "/warehouses/ok", "")

		// Act
		s.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
