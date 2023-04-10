package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/productRecords"
	"github.com/stretchr/testify/assert"
)

var dataPR = []domain.ProductRecords{
	{
		ID:             1,
		LastUpdateDate: "2022-12-04",
		PurchasePrice:  20.9,
		SalePrice:      90.8,
		ProductID:      1,
	},
	{
		ID:             2,
		LastUpdateDate: "2022-12-04",
		PurchasePrice:  20.9,
		SalePrice:      90.8,
		ProductID:      1,
	},
}

func createServerProductRecords(mockService *productRecords.MockServiceProductRecords) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	handler := NewProductRecord(mockService)
	r := gin.Default()
	pr := r.Group("/products")

	pr.POST("/productRecords", handler.Create())

	return r
}

func createRequestTestProductRecord(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application-json")
	return req, httptest.NewRecorder()
}

func TestCreateOkProductRecord(t *testing.T) {

	//Arrange
	pr := domain.ProductRecords{
		ID:             3,
		LastUpdateDate: "2022-12-04",
		PurchasePrice:  109.0,
		SalePrice:      123.9,
		ProductID:      1,
	}
	var productRecordNew map[string]domain.ProductRecords

	mockService := productRecords.MockServiceProductRecords{
		DataMock: dataPR,
	}

	//Act
	server := createServerProductRecords(&mockService)
	request, received := createRequestTestProductRecord(http.MethodPost, "/products/productRecords",
		`{
            "last_update_date": "2022-12-04",
            "purchase_price": 109.0,
            "sale_price": 123.9,
            "products_id": 1
        }`)

	server.ServeHTTP(received, request)
	t.Log(received.Body.String())
	err := json.Unmarshal(received.Body.Bytes(), &productRecordNew)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 200, received.Code)
	assert.Equal(t, pr, productRecordNew["data"])
}

// func TestCreateFailProductRecord(t *testing.T) {

// 	samples := []struct {
// 		Name   string
// 		Method string
// 		Path   string
// 		Body   string
// 		Code   int
// 	}{
// 		{
// 			Name:   "last_update_date is required",
// 			Method: http.MethodPost, Path: "/products/productRecords",
// 			Body: "{purchase_price: 30.8, sale_price: 78.9, products_id: 1}",
// 			Code: 400,
// 		},
// 	}

// 	s := new(productRecords.MockServiceProductRecords)
// 	h := NewProductRecord(s)
// 	server := utils.ServerInit(h, "/products/productRecords")

// 	for _, ts := range samples {
// 		t.Run(ts.Name, func(t *testing.T) {
// 			req, rr := utils.ServerRequest(ts.Method, ts.Path, ts.Body)

// 			server.ServeHTTP(rr, req)

// 			assert.Equal(t, ts.Code, rr.Code)
// 		})
// 	}
// }
