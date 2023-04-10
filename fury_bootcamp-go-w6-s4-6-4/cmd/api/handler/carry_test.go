package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/carry"
	"github.com/stretchr/testify/assert"
)

var carryService = createCarryService()

func createCarryService() *gin.Engine {
	r := gin.Default()
	service := carry.NewServiceCarry([]domain.Carry{})
	handler := NewCarry(service)

	r.POST("/carries", handler.Create())
	return r
}

func createRequestTestCarry(method string, path string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, path, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestCreateCarry(t *testing.T) {
	t.Run("should create a new carry", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestCarry(http.MethodPost, "/carries", `{"cid":"1", "company_name":"name", "address":"address", "telephone":"telephone", "locality_id":1}`)

		// Act
		carryService.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, `{"data":{"id":1,"cid":"1","company_name":"name","address":"address","telephone":"telephone","locality_id":1}}`, w.Body.String())
	})
	t.Run("should return 422 when create a carry with invalid body", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestCarry(http.MethodPost, "/carries", `{"cid":"1", "company_name":"name", "address":"address"}`)

		// Act
		carryService.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})
	t.Run("should not create a new carry with same carrycode", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestCarry(http.MethodPost, "/carries", `{"cid":"1", "company_name":"name", "address":"address", "telephone":"telephone", "locality_id":1}`)

		// Act
		carryService.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusConflict, w.Code)
		assert.Equal(t, `{"code":"conflict","message":"carry with code 1 already exists"}`, w.Body.String())
	})
	t.Run("should not create a new carry with non existent locality", func(t *testing.T) {
		// Arrange
		req, w := createRequestTestCarry(http.MethodPost, "/carries", `{"cid":"2", "company_name":"name", "address":"address", "telephone":"telephone", "locality_id":2}`)

		// Act
		carryService.ServeHTTP(w, req)

		// Assert
		assert.Equal(t, http.StatusConflict, w.Code)
		assert.Equal(t, `{"code":"conflict","message":"locality with code 2 not exists"}`, w.Body.String())
	})
}
