package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	inboundorder "github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/inbound_order"
	"github.com/stretchr/testify/assert"
)

var data_ibo = []domain.Inbound_order{
	{
		ID:               1,
		Order_date:       "2021-04-04",
		Order_number:     "order#1",
		Employee_id:      4,
		Product_batch_id: 1,
		Warehouse_id:     1,
	},
	{
		ID:               2,
		Order_date:       "2022-01-09",
		Order_number:     "order#1",
		Employee_id:      4,
		Product_batch_id: 1,
		Warehouse_id:     1,
	},
}

func createServerIBO(mockService *inboundorder.MockServiceIBO) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	handler := NewInBound_Order(mockService)
	r := gin.Default()
	ibor := r.Group("/inboundOrders")
	ibor.GET("/", handler.GetAll())
	ibor.POST("/", handler.Create())
	return r
}

func createRequestTestIBO(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application-json")
	return req, httptest.NewRecorder()
}

func Test_GetAll_Handler_IBO(t *testing.T) {
	t.Run("Get All Ok", func(t *testing.T) {
		//Arrange
		ibosExpected := data_ibo
		var ibosResult map[string][]domain.Inbound_order
		var dat []domain.Inbound_order
		dat = append(dat, data_ibo...)
		myMockS := inboundorder.MockServiceIBO{DataMock: dat}
		server := createServerIBO(&myMockS)

		//Act
		req, rec := createRequestTestEmployee(http.MethodGet, "/inboundOrders/", "")
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &ibosResult)
		//Assert
		assert.True(t, myMockS.MethodCalled)
		assert.Nil(t, err)
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, ibosExpected, ibosResult["data"])
	})

	t.Run("no in bound reports exist", func(t *testing.T) {
		//Arrange
		messageExpected := `{"data":"No existing Inbound_orders"}`
		myMockS := inboundorder.MockServiceIBO{}
		server := createServerIBO(&myMockS)

		//Act
		req, rec := createRequestTestEmployee(http.MethodGet, "/inboundOrders/", "")
		server.ServeHTTP(rec, req)

		//Assert
		assert.True(t, myMockS.MethodCalled)
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, messageExpected, rec.Body.String())
	})
}

func Test_Create_Handler_IBO(t *testing.T) {
	t.Run("Create Ok", func(t *testing.T) {
		iboexpected := domain.Inbound_order{
			ID:               3,
			Order_date:       "2021-04-04",
			Order_number:     "order#1",
			Employee_id:      4,
			Product_batch_id: 1,
			Warehouse_id:     1,
		}
		var ibosResult map[string]domain.Inbound_order
		var dat []domain.Inbound_order
		dat = append(dat, data_ibo...)
		myMockS := inboundorder.MockServiceIBO{DataMock: dat}
		server := createServerIBO(&myMockS)

		//Act
		req, rec := createRequestTestEmployee(http.MethodPost, "/inboundOrders/",
			`{
			"order_date": "2021-04-04",
			"order_number": "order#1",
			"employee_id": 4,
			"product_batch_id": 1,
			"warehouse_id": 1
		}`)
		server.ServeHTTP(rec, req)

		err := json.Unmarshal(rec.Body.Bytes(), &ibosResult)

		//Assert
		assert.True(t, myMockS.MethodCalled)
		assert.Nil(t, err)
		assert.Equal(t, 201, rec.Code)
		assert.Equal(t, iboexpected, ibosResult["data"])
	})

	t.Run("Create fail date", func(t *testing.T) {
		//Arrange
		errorExpected := "invalid date example 2008-01-02"
		var errorResult map[string]string
		var dat []domain.Inbound_order
		dat = append(dat, data_ibo...)
		myMockS := inboundorder.MockServiceIBO{DataMock: dat}
		server := createServerIBO(&myMockS)

		//Act
		req, rec := createRequestTestEmployee(http.MethodPost, "/inboundOrders/",
			`{
			"order_date": "04-04-2021",
			"order_number": "order#1",
			"employee_id": 4,
			"product_batch_id": 1,
			"warehouse_id": 1
		}`)
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &errorResult)

		//Assert
		assert.False(t, myMockS.MethodCalled)
		assert.Nil(t, err)
		assert.Equal(t, 422, rec.Code)
		assert.Equal(t, errorExpected, errorResult["message"])
	})

	t.Run("Required", func(t *testing.T) {
		//Arrange
		errorExpected := "El campo: Order_number es requerido"
		var errorResult map[string]string
		var dat []domain.Inbound_order
		dat = append(dat, data_ibo...)
		myMockS := inboundorder.MockServiceIBO{DataMock: dat}
		server := createServerIBO(&myMockS)

		//Act
		req, rec := createRequestTestEmployee(http.MethodPost, "/inboundOrders/",
			`{
			"order_date": "04-04-2021",
			"employee_id": 4,
			"product_batch_id": 1,
			"warehouse_id": 1
		}`)
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &errorResult)

		//Assert
		assert.False(t, myMockS.MethodCalled)
		assert.Nil(t, err)
		assert.Equal(t, 422, rec.Code)
		assert.Equal(t, errorExpected, errorResult["message"])
	})
}
