package handler

import (
	"bytes"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/section"
	sectionmock "github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/section"

	"github.com/stretchr/testify/assert"
)

func ServiceMockRepository(mockdb sectionmock.MockRepository) section.Service {
	return &sectionmock.MockService{Db: mockdb}
}

var DataStruck = []domain.Section{
	{
		ID:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseID:        1,
		ProductTypeID:      1,
	},
}

func createServerSection(mockService sectionmock.MockRepository) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	service := ServiceMockRepository(mockService)
	handler := NewSection(service)
	r := gin.Default()
	er := r.Group("/sections")
	er.GET("/", handler.GetAll())
	er.GET("/:id", handler.Get())
	er.DELETE("/:id", handler.Delete())
	er.PATCH("/:id", handler.Update())
	er.POST("/", handler.Create())

	return r
}

func createRequestTestSection(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application-json")
	return req, httptest.NewRecorder()
}

func TestCreate_okSection(t *testing.T) {
	//Arrange
	//Preparar
	mockService := sectionmock.MockRepository{DataMock: DataStruck, ExistsID: false, ID: 1}
	server := createServerSection(mockService)
	req, respons := createRequestTestSection(http.MethodPost, "/sections/",
		`{
			"section_number": 2, 
			"current_temperature": 1, 
			"minimum_temperature": 1, 
			"current_capacity": 2, 
			"minimum_capacity": 2, 
			"maximum_capacity": 2, 
			"warehouse_id": 1, 
			"product_type_id": 1
		}`)

	data := struct {
		Data domain.Section
	}{}

	server.ServeHTTP(respons, req)

	err := json.Unmarshal(respons.Body.Bytes(), &data)

	t.Log("Hi ", respons.Body.String())

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, respons.Code)
	assert.Equal(t, 2, data.Data.ID)

}

func TestCreate_conflictSection(t *testing.T) {
	mockService := sectionmock.MockRepository{DataMock: DataStruck, Error: "section already exists"}
	server := createServerSection(mockService)
	req, respons := createRequestTestSection(http.MethodPost, "/sections/", `{
		"section_number": 1, 
		"current_temperature": 2, 
		"minimum_temperature": 2, 
		"current_capacity": 2, 
		"minimum_capacity": 2, 
		"maximum_capacity": 1, 
		"warehouse_id": 1, 
		"product_type_id": 1
	}`)
	data := struct {
		Message string
		Error   string
	}{}

	server.ServeHTTP(respons, req)
	t.Log("----->", respons.Body.String())
	err := json.Unmarshal(respons.Body.Bytes(), &data)

	t.Log("Error", respons.Code)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusConflict, respons.Code)
	assert.Equal(t, "section already exists", data.Message)

}

func TestCreate_failSection(t *testing.T) {
	mockService := sectionmock.MockRepository{DataMock: DataStruck, Error: "section already exists"}
	server := createServerSection(mockService)
	req, respons := createRequestTestSection(http.MethodPost, "/sections/", `{
		"section_number": 1, 
		"current_temperature": 2, 
		"minimum_temperature": 2, 
		"current_capacity": 2, 
		"minimum_capacity": 2, 
		"maximum_capacity": 1, 
		"warehouse_id": 1, 
		"product_type_id": 1
	}`)
	data := struct {
		Message string
		Error   string
	}{}

	server.ServeHTTP(respons, req)
	t.Log("----->", respons.Body.String())
	err := json.Unmarshal(respons.Body.Bytes(), &data)

	t.Log("Error", respons.Code)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusConflict, respons.Code)

}

// TestFind_AllSection
func TestFind_AllSection(t *testing.T) {

	var DataStruck2 = []domain.Section{
		{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
		{
			ID:                 2,
			SectionNumber:      2,
			CurrentTemperature: 2,
			MinimumTemperature: 2,
			CurrentCapacity:    2,
			MinimumCapacity:    2,
			MaximumCapacity:    2,
			WarehouseID:        2,
			ProductTypeID:      2,
		},
	}

	mockService := sectionmock.MockRepository{DataMock: DataStruck2}
	server := createServerSection(mockService)
	req, respons := createRequestTestSection(http.MethodGet, "/sections/", "")

	server.ServeHTTP(respons, req)
	//t.Log("Lo que tiene data: ",data)
	//t.Log("----->", respons.Body.String())

	var resp map[string][]domain.Section

	err := json.Unmarshal(respons.Body.Bytes(), &resp)

	//t.Log("respons: ",respons.Body)
	//t.Log("data: ",data)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, respons.Code)
	assert.Equal(t, 2, len(resp["data"]))
	assert.Equal(t, mockService.DataMock, resp["data"])
	//t.Log("DATA MOCKSERVICE",mockService.DataMock)
	t.Log("DATA RESPONSE", resp["data"])
}

// TestFind_by_id_non_existentSection
func TestFind_by_id_non_existentSection(t *testing.T) {

	mockService := sectionmock.MockRepository{DataMock: DataStruck}
	server := createServerSection(mockService)
	req, respons := createRequestTestSection(http.MethodGet, "/sections/2", "")

	data := struct {
		Data domain.Section
	}{}

	server.ServeHTTP(respons, req)
	t.Log("Lo que tiene data: ", data)
	t.Log("----->", respons.Body.String())

	err := json.Unmarshal(respons.Body.Bytes(), &data)

	t.Log("respons: ", respons.Code)
	t.Log("data: ", data.Data)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusConflict, respons.Code)

}

// TestFind_by_id_existentSection
func TestFind_by_id_existentSection(t *testing.T) {

	mockService := sectionmock.MockRepository{DataMock: DataStruck}
	server := createServerSection(mockService)
	req, respons := createRequestTestSection(http.MethodGet, "/sections/1", "")

	data := struct {
		Data domain.Section
	}{}

	server.ServeHTTP(respons, req)
	t.Log("Lo que tiene data: ", data)
	t.Log("----->", respons.Body.String())

	err := json.Unmarshal(respons.Body.Bytes(), &data)

	t.Log("respons: ", respons.Code)
	t.Log("data: ", data.Data)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, respons.Code)
}

// TestUpdate_okSection
func TestUpdate_okSection(t *testing.T) {

	var expectedStruckUpdate = []domain.Section{
		{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 2,
			MinimumTemperature: 3,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseID:        1,
			ProductTypeID:      88,
		},
	}

	mockService := sectionmock.MockRepository{DataMock: DataStruck}
	server := createServerSection(mockService)
	req, respons := createRequestTestSection(http.MethodPatch, "/sections/1", `{

 		"current_temperature":2,
 		"minimum_temperature":3,
		"product_type_id":88
	}`)

	t.Log("dataStruck sin update: ", DataStruck)

	data := struct {
		Data domain.Section
	}{}

	server.ServeHTTP(respons, req)

	err := json.Unmarshal(respons.Body.Bytes(), &data)

	t.Log("Código de respuesta: ", respons.Code)

	t.Log("DataStruck esperada: ", expectedStruckUpdate)
	t.Log("DataStruck Obtenida: ", DataStruck)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, respons.Code)
	assert.Equal(t, expectedStruckUpdate, DataStruck)
}

// TestUpdate_non_existentSection
func TestUpdate_non_existentSection(t *testing.T) {
	mockService := sectionmock.MockRepository{DataMock: DataStruck}
	server := createServerSection(mockService)
	req, respons := createRequestTestSection(http.MethodPatch, "/sections/2", `{

 		"current_temperature":2,
 		"minimum_temperature":3,
		"product_type_id":88
	}`)
	codigoExpected := http.StatusNotFound

	t.Log("dataStruck sin update: ", DataStruck)

	data := struct {
		Data domain.Section
	}{}

	server.ServeHTTP(respons, req)

	err := json.Unmarshal(respons.Body.Bytes(), &data)

	t.Log("Código de respuesta: ", respons.Code)

	assert.Nil(t, err)
	assert.Equal(t, codigoExpected, respons.Code)
}

// TestDelete_non_existentSection
func TestDelete_non_existentSection(t *testing.T) {

	mockService := sectionmock.MockRepository{DataMock: DataStruck}
	server := createServerSection(mockService)
	req, respons := createRequestTestSection(http.MethodDelete, "/sections/2", "")

	errorExpected := "section not found"

	t.Log("dataStruck", DataStruck)

	data := struct {
		Data    domain.Section
		Message string
	}{}

	server.ServeHTTP(respons, req)

	err := json.Unmarshal(respons.Body.Bytes(), &data)
	t.Log("Código de respuesta: ", respons.Code)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, respons.Code)
	assert.Equal(t, errorExpected, data.Message)
}

// TestDelete_existentSection
func TestDelete_existentSection(t *testing.T) {

	var struck = []domain.Section{
		{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 2,
			MinimumTemperature: 3,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseID:        1,
			ProductTypeID:      88,
		},
		{
			ID:                 2,
			SectionNumber:      2,
			CurrentTemperature: 2,
			MinimumTemperature: 3,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseID:        1,
			ProductTypeID:      88,
		},
	}

	mockService := sectionmock.MockRepository{DataMock: struck}
	server := createServerSection(mockService)
	req, respons := createRequestTestSection(http.MethodDelete, "/sections/1", "")
	data := struct {
		Data domain.Section
	}{}

	server.ServeHTTP(respons, req)

	t.Log("Struct antes de ser eliminado ", struck)

	err := json.Unmarshal(respons.Body.Bytes(), &data)

	t.Log("Código de respuesta: ", respons.Code)

	assert.Error(t, err)
	assert.Equal(t, http.StatusNoContent, respons.Code)
}
