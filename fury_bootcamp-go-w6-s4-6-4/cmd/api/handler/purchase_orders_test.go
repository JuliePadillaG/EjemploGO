package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/purchase_orders"
)

func createServerPurchaseOrders(mockService purchase_orders.MockService) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	handler := NewPurchaseOrders(&mockService)
	router := gin.Default()

	pr := router.Group("/api/v1/")

	pu := pr.Group("purchaseOrders")
	pu.POST("", handler.Create())

	pe := pr.Group("buyers")
	pe.GET("reportPurchaseOrders", handler.Get())

	return router
}

func createRequestPO(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestHandlerPurchaseOrdersCreate(t *testing.T) {

	t.Run("TestHandlerPurchaseOrdersCreateOk", func(t *testing.T) {
		// Arrange
		orderDate := time.Now()
		database := []domain.PurchaseOrders{
			{
				OrderNumber:     "232345",
				OrderDate:       &orderDate,
				TrackingCode:    "bar",
				BuyerID:         1,
				ProductRecordID: 1,
				OrderStatusID:   1,
			},
		}

		idExpected := 2
		requestBodyJSON := struct {
			OrderNumber     string `json:"order_number"`
			OrderDate       string `json:"order_date"`
			TrackingCode    string `json:"tracking_code"`
			BuyerID         int    `json:"buyer_id"`
			ProductRecordID int    `json:"product_record_id"`
			OrderStatusID   int    `json:"order_status_id"`
		}{
			OrderNumber:     "232346",
			OrderDate:       "2021-11-12",
			TrackingCode:    "foo",
			BuyerID:         1,
			ProductRecordID: 1,
			OrderStatusID:   1,
		}

		productToSave, err := json.Marshal(&requestBodyJSON)
		if err != nil {
			panic(err)
		}

		// t.Logf("%+v\n", string(productToSave))

		mockService := purchase_orders.MockService{
			DataMock: database,
		}

		var resp map[string]domain.PurchaseOrders

		// Act
		router := createServerPurchaseOrders(mockService)
		req, rr := createRequestPO(http.MethodPost, `/api/v1/purchaseOrders`, string(productToSave))
		router.ServeHTTP(rr, req)

		err = json.Unmarshal(rr.Body.Bytes(), &resp)
		if err != nil {
			panic(err)
		}

		// t.Logf(" %+v\n\n%+v \n", rr, req)
		// t.Logf(" %+v\n", resp)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, http.StatusCreated, rr.Code)
		assert.Equal(t, idExpected, resp["data"].ID)

	})

	t.Run("TestHandlerPurchaseOrdersCreateBuyerIDNotFoundFail", func(t *testing.T) {
		// Arrange
		errMessage := "buyer_id doesn't exists"
		orderDate := time.Now()
		database := []domain.PurchaseOrders{
			{
				OrderNumber:     "232345",
				OrderDate:       &orderDate,
				TrackingCode:    "bar",
				BuyerID:         1,
				ProductRecordID: 1,
				OrderStatusID:   1,
			},
		}

		// idExpected := 2
		requestBodyJSON := struct {
			OrderNumber     string `json:"order_number"`
			OrderDate       string `json:"order_date"`
			TrackingCode    string `json:"tracking_code"`
			BuyerID         int    `json:"buyer_id"`
			ProductRecordID int    `json:"product_record_id"`
			OrderStatusID   int    `json:"order_status_id"`
		}{
			OrderNumber:     "232346",
			OrderDate:       "1999-11-10",
			TrackingCode:    "foo",
			BuyerID:         1,
			ProductRecordID: 1,
			OrderStatusID:   1,
		}

		productToSave, err := json.Marshal(&requestBodyJSON)
		if err != nil {
			panic(err)
		}

		mockService := purchase_orders.MockService{
			DataMock: database,
			Error:    errMessage,
		}

		var resp map[string]string
		expected := map[string]string{
			"code":    "conflict",
			"message": "buyer_id doesn't exists",
		}

		// Act
		router := createServerPurchaseOrders(mockService)
		req, rr := createRequestPO(http.MethodPost, `/api/v1/purchaseOrders`, string(productToSave))
		router.ServeHTTP(rr, req)

		err = json.Unmarshal(rr.Body.Bytes(), &resp)
		if err != nil {
			panic(err)
		}

		// t.Logf(" %+v\n\n%+v \n", rr, req)
		// t.Logf(" %+v\n", resp)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected, resp)
	})

	t.Run("TestHandlerPurchaseOrdersCreateFail", func(t *testing.T) {
		// Arrange
		errMessage := "Key: 'RequestPurchaseOrders.TrackingCode' Error:Field validation for 'TrackingCode' failed on the 'required' tag"
		database := []domain.PurchaseOrders{}

		requestBodyJSON := struct {
			OrderNumber     string `json:"order_number"`
			OrderDate       string `json:"order_date"`
			TrackingCode    string `json:"tracking_code"`
			BuyerID         int    `json:"buyer_id"`
			ProductRecordID int    `json:"product_record_id"`
			OrderStatusID   int    `json:"order_status_id"`
		}{
			OrderNumber:     "232346",
			OrderDate:       "2021-11-12",
			BuyerID:         1,
			ProductRecordID: 1,
			OrderStatusID:   1,
		}

		productToSave, err := json.Marshal(&requestBodyJSON)
		if err != nil {
			panic(err)
		}

		// Missing "FirstName" field

		mockService := purchase_orders.MockService{
			DataMock: database,
			Error:    errMessage,
		}

		var resp map[string]string

		// Act
		router := createServerPurchaseOrders(mockService)
		req, rr := createRequestPO(http.MethodPost, `/api/v1/purchaseOrders`, string(productToSave))
		router.ServeHTTP(rr, req)
		err = json.Unmarshal(rr.Body.Bytes(), &resp)

		// fmt.Printf("--- %+v\n", rr)
		// fmt.Printf("--- %+v\n", err)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		assert.Equal(t, errMessage, resp["message"])
	})

	t.Run("TestHandlerPurchaseOrdersCreateDateFail", func(t *testing.T) {
		// Arrange
		errMessage := "Key: 'RequestPurchaseOrders.TrackingCode' Error:Field validation for 'TrackingCode' failed on the 'required' tag"
		database := []domain.PurchaseOrders{}

		requestBodyJSON := struct {
			OrderNumber     string `json:"order_number"`
			OrderDate       string `json:"order_date"`
			TrackingCode    string `json:"tracking_code"`
			BuyerID         int    `json:"buyer_id"`
			ProductRecordID int    `json:"product_record_id"`
			OrderStatusID   int    `json:"order_status_id"`
		}{
			OrderNumber:     "232346",
			OrderDate:       "0999-11-12",
			TrackingCode:    "foo",
			BuyerID:         1,
			ProductRecordID: 1,
			OrderStatusID:   1,
		}

		productToSave, err := json.Marshal(&requestBodyJSON)
		if err != nil {
			panic(err)
		}

		// Missing "FirstName" field

		mockService := purchase_orders.MockService{
			DataMock: database,
			Error:    errMessage,
		}

		var resp map[string]string

		// Act
		router := createServerPurchaseOrders(mockService)
		req, rr := createRequestPO(http.MethodPost, `/api/v1/purchaseOrders`, string(productToSave))
		router.ServeHTTP(rr, req)
		err = json.Unmarshal(rr.Body.Bytes(), &resp)

		// fmt.Printf("--- %+v\n", rr)
		// fmt.Printf("--- %+v\n", err)

		// Assert
		assert.Nil(t, err)
		// assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
		// assert.Equal(t, errMessage, resp["message"])
	})

}

func TestHandlerReport(t *testing.T) {

	t.Run("TestHandlerReportFindByIdExistentAll", func(t *testing.T) {
		// arrange
		database := []domain.ReportPurchaseOrders{
			{
				ID:                  1,
				CardNumberID:        "402323",
				FirstName:           "Peter",
				LastName:            "Peter",
				PurchaseOrdersCount: 3,
			},
		}

		mockService := purchase_orders.MockService{
			DataMockReports: database,
		}

		var resp map[string][]domain.ReportPurchaseOrders

		router := createServerPurchaseOrders(mockService)
		req, rr := createRequestPO(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders", "")

		// Act
		router.ServeHTTP(rr, req)
		err := json.Unmarshal(rr.Body.Bytes(), &resp)

		// fmt.Printf("--- %+v\n", rr)
		// fmt.Printf("--- %+v\n", err)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, mockService.DataMockReports[0], resp["data"][0])
	})

	t.Run("TestHandlerReportFindll", func(t *testing.T) {
		// arrange
		database := []domain.ReportPurchaseOrders{
			{
				ID:                  1,
				CardNumberID:        "402323",
				FirstName:           "Peter",
				LastName:            "Peter",
				PurchaseOrdersCount: 3,
			},
		}

		mockService := purchase_orders.MockService{
			DataMockReports: database,
		}

		var resp map[string][]domain.ReportPurchaseOrders

		router := createServerPurchaseOrders(mockService)
		req, rr := createRequestPO(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders?id=1", "")

		// Act
		router.ServeHTTP(rr, req)
		err := json.Unmarshal(rr.Body.Bytes(), &resp)

		// fmt.Printf("--- %+v\n", rr)
		// fmt.Printf("--- %+v\n", err)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, rr.Code)
		assert.Equal(t, mockService.DataMockReports[0], resp["data"][0])
	})

	t.Run("TestHandlerReportFindFail", func(t *testing.T) {
		// arrange
		database := []domain.ReportPurchaseOrders{}

		mockService := purchase_orders.MockService{
			DataMockReports: database,
			// Error: "not Found",
		}

		var resp map[string][]domain.ReportPurchaseOrders

		router := createServerPurchaseOrders(mockService)
		req, rr := createRequestPO(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders?id=1", "")

		// Act
		router.ServeHTTP(rr, req)
		err := json.Unmarshal(rr.Body.Bytes(), &resp)

		// t.Logf("--- %+v\n", resp)
		// t.Logf("--- %+v\n", err)

		// Assert
		assert.NotNil(t, err)
	})

	t.Run("TestHandlerReportFindFailbuyerID", func(t *testing.T) {
		// arrange
		database := []domain.ReportPurchaseOrders{}

		mockService := purchase_orders.MockService{
			DataMockReports: database,
			// Error: "not Found",
		}

		var resp map[string]string
		expected := map[string]string{
			"code":    "not_found",
			"message": "there is no purchase orders",
		}

		router := createServerPurchaseOrders(mockService)
		req, rr := createRequestPO(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders?buyerID=foo", "")

		// Act
		router.ServeHTTP(rr, req)
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		if err != nil {
			panic(err)
		}

		t.Logf("--- %+v\n\n%+v\n", req, rr)
		t.Logf("--- %+v\n", err)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected, resp)
	})

	t.Run("TestHandlerReportBuyerIDNotFound", func(t *testing.T) {
		// arrange
		database := []domain.ReportPurchaseOrders{
			{
				ID:                  1,
				CardNumberID:        "402323",
				FirstName:           "Peter",
				LastName:            "Peter",
				PurchaseOrdersCount: 3,
			},
		}

		mockService := purchase_orders.MockService{
			DataMockReports: database,
			Error:           "there is no purchase orders",
		}

		var resp map[string]string
		expected := map[string]string{
			"code":    "not_found",
			"message": "there is no purchase orders",
		}

		router := createServerPurchaseOrders(mockService)
		req, rr := createRequestPO(http.MethodGet, "/api/v1/buyers/reportPurchaseOrders?id=2", "")

		// Act
		router.ServeHTTP(rr, req)
		err := json.Unmarshal(rr.Body.Bytes(), &resp)

		// t.Logf("b --- %+v\n", rr)
		// t.Logf("b --- %+v\n", err)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, expected, resp)
		// assert.Equal(t, http.StatusOK, rr.Code)
		// assert.Equal(t, mockService.DataMockReports[0], resp["data"][0])
	})
}
