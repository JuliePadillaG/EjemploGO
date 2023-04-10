package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/employee"
	"github.com/stretchr/testify/assert"
)

var data = []domain.Employee{
	{
		ID:           1,
		CardNumberID: "402323",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  1,
	},
	{
		ID:           2,
		CardNumberID: "402322",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  3,
	},
}

func createServerEmployee(mockService *employee.MockServiceEmployee) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	handler := NewEmployee(mockService)
	r := gin.Default()
	er := r.Group("/employees")
	er.POST("/", handler.Create())
	er.GET("/", handler.GetAll())
	er.GET("/:id", handler.Get())
	er.PATCH("/:id", handler.Update())
	er.DELETE("/:id", handler.Delete())
	er.GET("/reportInboundOrders", handler.Report_InboundOrders())

	return r
}

func createRequestTestEmployee(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application-json")
	return req, httptest.NewRecorder()
}

func Test_create_ok_Handler(t *testing.T) {
	t.Run("Create Ok", func(t *testing.T) {
		//Arrange
		employeeExpected := domain.Employee{
			ID:           3,
			CardNumberID: "402324",
			FirstName:    "Christian",
			LastName:     "Jhonson",
			WarehouseID:  7,
		}
		var employeeNew map[string]domain.Employee
		var dat []domain.Employee
		dat = append(dat, data...)
		myMockS := employee.MockServiceEmployee{DataMock: dat}
		server := createServerEmployee(&myMockS)
		//Act
		req, rec := createRequestTestEmployee(http.MethodPost, "/employees/",
			`{
			"card_number_id": "402324",
			"first_name": "Christian",
			"last_name": "Jhonson",
			"warehouse_id": 7
		 }`)

		server.ServeHTTP(rec, req)

		err := json.Unmarshal(rec.Body.Bytes(), &employeeNew)

		//Assert
		assert.True(t, myMockS.MethodCalled)
		assert.Nil(t, err)
		assert.Equal(t, 201, rec.Code)
		assert.Equal(t, employeeExpected, employeeNew["data"])
	})

	//Create sin WareHouse
	t.Run("Create Ok sin WareHouse", func(t *testing.T) {
		//Arrange
		employeeExpected := domain.Employee{
			ID:           3,
			CardNumberID: "402325",
			FirstName:    "Christian",
			LastName:     "Jhonson",
			WarehouseID:  0,
		}
		var employeeNew map[string]domain.Employee
		var dat []domain.Employee
		dat = append(dat, data...)
		myMockS := employee.MockServiceEmployee{DataMock: dat}
		server := createServerEmployee(&myMockS)

		//Act
		req, rec := createRequestTestEmployee(http.MethodPost, "/employees/",
			`{
			"card_number_id": "402325",
			"first_name": "Christian",
			"last_name": "Jhonson"
		 }`)
		server.ServeHTTP(rec, req)

		err := json.Unmarshal(rec.Body.Bytes(), &employeeNew)

		//ASSERT
		assert.True(t, myMockS.MethodCalled)
		assert.Nil(t, err)
		assert.Equal(t, 201, rec.Code)
		assert.Equal(t, employeeExpected, employeeNew["data"])
	})

}

func Test_create_fail_Handler(t *testing.T) {
	//Card number
	t.Run("Create fail card number", func(t *testing.T) {
		//Arrange
		errorExpected := "card_number_id is required"
		var errorResult map[string]string

		var dat []domain.Employee
		dat = append(dat, data...)
		myMockS := employee.MockServiceEmployee{DataMock: dat}
		server := createServerEmployee(&myMockS)
		req, rec := createRequestTestEmployee(http.MethodPost, "/employees/",
			`{
			"first_name": "Christian",
			"last_name": "Jhonson",
			"warehouse_id": 7
		}`)
		//Act
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &errorResult)

		//Assert
		assert.False(t, myMockS.MethodCalled)
		assert.Nil(t, err)
		assert.Equal(t, 422, rec.Code)
		assert.Equal(t, errorExpected, errorResult["message"])
	})

	//first name
	t.Run("Create fail first name", func(t *testing.T) {
		//Arrange
		errorExpected := "first_name is required"
		var errorResult map[string]string

		var dat []domain.Employee
		dat = append(dat, data...)
		myMockS := employee.MockServiceEmployee{DataMock: dat}
		server := createServerEmployee(&myMockS)
		req, rec := createRequestTestEmployee(http.MethodPost, "/employees/",
			`{
		"card_number_id": "402325",
        "last_name": "Jhonson",
        "warehouse_id": 7
		}`)

		//Act
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &errorResult)

		//Assert
		assert.False(t, myMockS.MethodCalled)
		assert.Nil(t, err)
		assert.Equal(t, 422, rec.Code)
		assert.Equal(t, errorExpected, errorResult["message"])
	})

	//Last name
	t.Run("Create fail last name", func(t *testing.T) {
		//Arrange
		errorExpected := "last_name is required"
		var errorResult map[string]string

		var dat []domain.Employee
		dat = append(dat, data...)
		myMockS := employee.MockServiceEmployee{DataMock: dat}
		server := createServerEmployee(&myMockS)
		req, rec := createRequestTestEmployee(http.MethodPost, "/employees/",
			`{
			"card_number_id": "402325",
			"first_name": "Christian",
			"warehouse_id": 7
		}`)

		//Act
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &errorResult)

		//Assert
		assert.False(t, myMockS.MethodCalled)
		assert.Nil(t, err)
		assert.Equal(t, 422, rec.Code)
		assert.Equal(t, errorExpected, errorResult["message"])
	})

	//Warehouse ID
	t.Run("Create fail wherehouse id", func(t *testing.T) {
		//Arrange
		errorExpected := "WareHouseID cannot be negative"
		var errorResult map[string]string

		var dat []domain.Employee
		dat = append(dat, data...)
		myMockS := employee.MockServiceEmployee{DataMock: dat}
		server := createServerEmployee(&myMockS)
		req, rec := createRequestTestEmployee(http.MethodPost, "/employees/",
			`{
			"card_number_id": "402324",
			"first_name": "Christian",
			"last_name": "Jhonson",
			"warehouse_id": -7
		}`)

		//Act
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &errorResult)

		//Assert
		assert.False(t, myMockS.MethodCalled)
		assert.Nil(t, err)
		assert.Equal(t, 422, rec.Code)
		assert.Equal(t, errorExpected, errorResult["message"])
	})

	//Error
	t.Run("Create fail Error", func(t *testing.T) {
		//Arrange
		errorExpected := "Error en el server"
		var errorResult map[string]string

		var dat []domain.Employee
		dat = append(dat, data...)
		myMockS := employee.MockServiceEmployee{DataMock: dat, Err: errorExpected}
		server := createServerEmployee(&myMockS)
		req, rec := createRequestTestEmployee(http.MethodPost, "/employees/",
			`{
			"card_number_id": "402324",
			"first_name": "Christian",
			"last_name": "Jhonson",
			"warehouse_id": 9
		}`)

		//Act
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &errorResult)

		//Assert
		assert.True(t, myMockS.MethodCalled)
		assert.Nil(t, err)
		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, errorExpected, errorResult["message"])
	})
}

func Test_create_conflict_Handler(t *testing.T) {
	//Arrage
	errorExpected := "The card_number_id already exists"
	var errorResult map[string]string
	var dat []domain.Employee
	dat = append(dat, data...)
	myMockS := employee.MockServiceEmployee{DataMock: dat}
	server := createServerEmployee(&myMockS)
	req, rec := createRequestTestEmployee(http.MethodPost, "/employees/",
		`{
        "card_number_id": "402323",
        "first_name": "Christian",
        "last_name": "Jhonson",
        "warehouse_id": 7
     }`)
	//Act
	server.ServeHTTP(rec, req)
	err := json.Unmarshal(rec.Body.Bytes(), &errorResult)
	//Assert
	assert.True(t, myMockS.MethodCalled)
	assert.Nil(t, err)
	assert.Equal(t, 409, rec.Code)
	assert.Equal(t, errorExpected, errorResult["message"])
}

func Test_find_all_Handler(t *testing.T) {
	t.Run("Find all Ok", func(t *testing.T) {
		//Arrange
		employeesExpected := data
		var employeesResult map[string][]domain.Employee
		var dat []domain.Employee
		dat = append(dat, data...)
		myMockS := employee.MockServiceEmployee{DataMock: dat}
		server := createServerEmployee(&myMockS)
		req, rec := createRequestTestEmployee(http.MethodGet, "/employees/", "")
		//Act
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &employeesResult)
		//Assert
		assert.True(t, myMockS.MethodCalled)
		assert.Nil(t, err)
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, employeesExpected, employeesResult["data"])
	})

	//Mensaje Cuando no exixten empleados aún
	t.Run("No existing employees", func(t *testing.T) {
		//Arrange
		messageExpected := `{"data":"No existing employees"}`

		myMockS := employee.MockServiceEmployee{}
		server := createServerEmployee(&myMockS)
		req, rec := createRequestTestEmployee(http.MethodGet, "/employees/", "")
		//Act
		server.ServeHTTP(rec, req)
		//Assert
		assert.True(t, myMockS.MethodCalled)
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, messageExpected, rec.Body.String())
	})

}

func Test_find_all_fail_Handler(t *testing.T) {
	//Arrange
	errorExpected := "Error en Get All de service"
	var errorResult map[string]string
	var dat []domain.Employee
	dat = append(dat, data...)
	myMockS := employee.MockServiceEmployee{DataMock: dat, Err: errorExpected}
	server := createServerEmployee(&myMockS)
	req, rec := createRequestTestEmployee(http.MethodGet, "/employees/", "")
	//Act
	server.ServeHTTP(rec, req)

	err := json.Unmarshal(rec.Body.Bytes(), &errorResult)
	//Assert
	assert.True(t, myMockS.MethodCalled)
	assert.Nil(t, err)
	assert.Equal(t, 404, rec.Code)
	assert.Equal(t, errorExpected, errorResult["message"])
}

func Test_find_by_id_non_existent_Handler(t *testing.T) {
	//Arrange
	errorExpected := "employee not found"
	var errorResult map[string]string
	var dat []domain.Employee
	dat = append(dat, data...)
	myMockS := employee.MockServiceEmployee{DataMock: dat}
	server := createServerEmployee(&myMockS)
	req, rec := createRequestTestEmployee(http.MethodGet, "/employees/4", "")
	//Act
	server.ServeHTTP(rec, req)

	err := json.Unmarshal(rec.Body.Bytes(), &errorResult)
	//Assert
	assert.True(t, myMockS.MethodCalled)
	assert.Nil(t, err)
	assert.Equal(t, 404, rec.Code)
	assert.Equal(t, errorExpected, errorResult["message"])
}

func Test_find_by_id_existent_Handler(t *testing.T) {
	//Arrange
	employeeExpected := domain.Employee{
		ID:           1,
		CardNumberID: "402323",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  1,
	}
	var employeeResult map[string]domain.Employee
	var dat []domain.Employee
	dat = append(dat, data...)
	myMockS := employee.MockServiceEmployee{DataMock: dat}
	server := createServerEmployee(&myMockS)
	req, rec := createRequestTestEmployee(http.MethodGet, "/employees/1", "")
	//Act
	server.ServeHTTP(rec, req)

	err := json.Unmarshal(rec.Body.Bytes(), &employeeResult)
	//Assert
	assert.True(t, myMockS.MethodCalled)
	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, employeeExpected, employeeResult["data"])
}

func Test_update_ok_Handler(t *testing.T) {
	//Arrange
	employeeExpected := domain.Employee{
		ID:           1,
		CardNumberID: "402323",
		FirstName:    "Jhon Updated",
		LastName:     "Doe Updated",
		WarehouseID:  1,
	}
	var employeeResult map[string]domain.Employee
	var dat []domain.Employee
	dat = append(dat, data...)
	myMockS := employee.MockServiceEmployee{DataMock: dat}
	server := createServerEmployee(&myMockS)
	req, rec := createRequestTestEmployee(http.MethodPatch, "/employees/1",
		`{
        "first_name": "Jhon Updated",
        "last_name": "Doe Updated",
        "warehouse_id": 1
     }`)
	//Act
	server.ServeHTTP(rec, req)

	err := json.Unmarshal(rec.Body.Bytes(), &employeeResult)
	//Assert
	assert.True(t, myMockS.MethodCalled)
	assert.Nil(t, err)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, employeeExpected, employeeResult["data"])
}

func Test_update_non_existent_Handler(t *testing.T) {
	//Arrange
	errorExpected := "employee not found"
	var errorResult map[string]string
	var dat []domain.Employee
	dat = append(dat, data...)
	myMockS := employee.MockServiceEmployee{DataMock: dat}
	server := createServerEmployee(&myMockS)
	req, rec := createRequestTestEmployee(http.MethodPatch, "/employees/9",
		`{
        "first_name": "Jhon Updated",
        "last_name": "Doe Updated",
        "warehouse_id": 1
     }`)
	//Act
	server.ServeHTTP(rec, req)

	err := json.Unmarshal(rec.Body.Bytes(), &errorResult)
	//Assert
	assert.True(t, myMockS.MethodCalled)
	assert.Nil(t, err)
	assert.Equal(t, 404, rec.Code)
	assert.Equal(t, errorExpected, errorResult["message"])
}

func Test_update_fail_Handler(t *testing.T) {
	//Arrange
	errorExpected := "card_number_id field cannot be updated"
	var errorResult map[string]string
	var dat []domain.Employee
	dat = append(dat, data...)
	myMockS := employee.MockServiceEmployee{DataMock: dat}
	server := createServerEmployee(&myMockS)
	req, rec := createRequestTestEmployee(http.MethodPatch, "/employees/1",
		`{
		"card_number_id": "402325",
        "first_name": "Jhon Updated",
        "last_name": "Doe Updated",
        "warehouse_id": 1
     }`)
	//Act
	server.ServeHTTP(rec, req)

	err := json.Unmarshal(rec.Body.Bytes(), &errorResult)
	//Assert
	assert.False(t, myMockS.MethodCalled)
	assert.Nil(t, err)
	assert.Equal(t, 400, rec.Code)
	assert.Equal(t, errorExpected, errorResult["message"])
}
func Test_delete_non_existent_Handler(t *testing.T) {
	//Arrange
	errorExpected := "employee not found"
	var errorResult map[string]string
	var dat []domain.Employee
	dat = append(dat, data...)
	myMockS := employee.MockServiceEmployee{DataMock: dat}
	server := createServerEmployee(&myMockS)
	req, rec := createRequestTestEmployee(http.MethodDelete, "/employees/9", "")
	//Act
	server.ServeHTTP(rec, req)

	err := json.Unmarshal(rec.Body.Bytes(), &errorResult)
	//Assert
	assert.True(t, myMockS.MethodCalled)
	assert.Nil(t, err)
	assert.Equal(t, 404, rec.Code)
	assert.Equal(t, errorExpected, errorResult["message"])
}

func Test_delete_ok_Handler(t *testing.T) {
	//Arrange
	dataCopy := make([]domain.Employee, len(data))
	copy(dataCopy, data)
	myMockS := employee.MockServiceEmployee{DataMock: data}
	server := createServerEmployee(&myMockS)
	req, rec := createRequestTestEmployee(http.MethodDelete, "/employees/1", "")
	//Act
	server.ServeHTTP(rec, req)
	//Assert
	t.Log(dataCopy)
	t.Log(myMockS.DataMock)
	assert.True(t, myMockS.MethodCalled)
	assert.Equal(t, 204, rec.Code)
	assert.NotEqual(t, dataCopy, myMockS.DataMock)
}

func Test_Report_InboundOrders_ok_Handler(t *testing.T) {
	//All reports
	t.Run("All Reports", func(t *testing.T) {
		//Arrange
		data := []domain.ReportInBO{
			{
				ID:                   1,
				CardNumberID:         "402323",
				FirstName:            "Jhon",
				LastName:             "Doe",
				WarehouseID:          1,
				Inbound_orders_count: 2,
			},
			{
				ID:                   2,
				CardNumberID:         "402324",
				FirstName:            "Jhon",
				LastName:             "Doe",
				WarehouseID:          1,
				Inbound_orders_count: 0,
			},
		}
		var reportsResult map[string][]domain.ReportInBO
		myMockS := employee.MockServiceEmployee{DatamockIBO: data}
		server := createServerEmployee(&myMockS)
		req, rec := createRequestTestEmployee(http.MethodGet, "/employees/reportInboundOrders", "")

		//Act
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &reportsResult)

		//Assert
		assert.NoError(t, err)
		assert.True(t, myMockS.MethodCalled)
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, data, reportsResult["data"])
	})

	//Reports by Id
	t.Run("Reports by Id", func(t *testing.T) {
		//Arrange
		data := []domain.ReportInBO{
			{
				ID:                   1,
				CardNumberID:         "402323",
				FirstName:            "Jhon",
				LastName:             "Doe",
				WarehouseID:          1,
				Inbound_orders_count: 2,
			},
		}
		var reportsResult map[string][]domain.ReportInBO
		myMockS := employee.MockServiceEmployee{DatamockIBO: data}
		server := createServerEmployee(&myMockS)
		req, rec := createRequestTestEmployee(http.MethodGet, "/employees/reportInboundOrders?id=1", "")

		//Act
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &reportsResult)

		//Assert
		assert.NoError(t, err)
		assert.True(t, myMockS.MethodCalled)
		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, data, reportsResult["data"])
	})
}

func Test_Report_InboundOrders_fail_Handler(t *testing.T) {
	//error caused
	t.Run("error caused", func(t *testing.T) {
		//Assert
		errExpected := "Error caused"
		var errorResult map[string]string
		myMockS := employee.MockServiceEmployee{Err: errExpected}
		server := createServerEmployee(&myMockS)
		req, rec := createRequestTestEmployee(http.MethodGet, "/employees/reportInboundOrders", "")

		//Act
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &errorResult)

		//Assert
		assert.NoError(t, err)
		assert.True(t, myMockS.MethodCalled)
		assert.Equal(t, 500, rec.Code)
		assert.Equal(t, errExpected, errorResult["message"])
	})

	//Employee not found
	t.Run("Employee not found", func(t *testing.T) {
		//Arrange
		errExpected := "employee not found"
		var errorResult map[string]string
		myMockS := employee.MockServiceEmployee{}
		server := createServerEmployee(&myMockS)
		req, rec := createRequestTestEmployee(http.MethodGet, "/employees/reportInboundOrders?id=1", "")

		//Act
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &errorResult)

		//Assert
		assert.NoError(t, err)
		assert.True(t, myMockS.MethodCalled)
		assert.Equal(t, 404, rec.Code)
		assert.Equal(t, errExpected, errorResult["message"])

	})
	// id not applicable
	t.Run("id not applicable", func(t *testing.T) {
		//Arrange
		errExpected := "invalid id format"
		var errorResult map[string]string
		myMockS := employee.MockServiceEmployee{}
		server := createServerEmployee(&myMockS)
		req, rec := createRequestTestEmployee(http.MethodGet, "/employees/reportInboundOrders?id=ls", "")

		//Act
		server.ServeHTTP(rec, req)
		err := json.Unmarshal(rec.Body.Bytes(), &errorResult)

		//Assert
		assert.NoError(t, err)
		assert.True(t, myMockS.MethodCalled)
		assert.Equal(t, 422, rec.Code)
		assert.Equal(t, errExpected, errorResult["message"])
	})
}
