package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/buyer"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createServer(mockService buyer.MockService) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	handler := NewBuyer(&mockService)
	router := gin.Default()

	rg := router.Group("/api/v1")
	pr := rg.Group("buyers")
	pr.GET("", handler.GetAll())
	pr.GET("/:id", handler.Get())
	pr.POST("", handler.Create())
	pr.PATCH("/:id", handler.Update())
	pr.DELETE("/:id", handler.Delete())

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestHandlerFindAll(t *testing.T) {
	// arrange
	database := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "232345",
			FirstName:    "hello",
			LastName:     "World",
		},
	}

	mockService := buyer.MockService{
		DataMock: database,
	}

	var resp map[string][]domain.Buyer

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/api/v1/buyers", "")

	// Act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)

	// fmt.Printf("--- %+v\n", rr)
	// fmt.Printf("--- %+v\n", err)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, mockService.DataMock, resp["data"])
}

func TestHandlerFindAllFail(t *testing.T) {
	// arrange
	errMessage := "There is no buyers"
	database := []domain.Buyer{}

	mockService := buyer.MockService{
		DataMock: database,
		Error:    errMessage,
	}

	var resp map[string]string

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/api/v1/buyers", "")

	// Act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)

	t.Logf("t --- %+v\n", rr)
	t.Logf("t --- %+v\n", err)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, mockService.Error, resp["message"])
}

func TestHandlerFindByIdExistent(t *testing.T) {
	// arrange
	database := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "232345",
			FirstName:    "hello",
			LastName:     "World",
		},
		{
			ID:           2,
			CardNumberID: "232346",
			FirstName:    "hello_2",
			LastName:     "World_2",
		},
	}

	mockService := buyer.MockService{
		DataMock: database,
	}

	var resp map[string][]domain.Buyer

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/api/v1/buyers/2", "")

	// Act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)

	// fmt.Printf("--- %+v\n", rr)
	// fmt.Printf("--- %+v\n", err)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, mockService.DataMock[1], resp["data"][0])
}

func TestHandlerFindByIdNotExistent(t *testing.T) {
	// arrange
	errMessage := "there is no buyers"
	database := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "232345",
			FirstName:    "hello",
			LastName:     "World",
		},
		{
			ID:           2,
			CardNumberID: "232346",
			FirstName:    "hello_2",
			LastName:     "World_2",
		},
	}

	mockService := buyer.MockService{
		DataMock: database,
		Error:    errMessage,
	}

	var resp map[string]string

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/api/v1/buyers/5", "")

	// Act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)

	// fmt.Printf("--- %+v\n", rr)
	// fmt.Printf("--- %+v\n", err)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, mockService.Error, resp["message"])
}

func TestHandlerFindByIdBadID(t *testing.T) {
	// arrange
	errMessage := "strconv.Atoi: parsing \"bar\": invalid syntax"
	database := []domain.Buyer{}

	mockService := buyer.MockService{
		DataMock: database,
		Error:    errMessage,
	}

	var resp map[string]string

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodGet, "/api/v1/buyers/bar", "")

	// Act
	router.ServeHTTP(rr, req)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)

	// fmt.Printf("--- %+v\n", rr)
	// fmt.Printf("--- %+v\n", err)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, mockService.Error, resp["message"])
}

func TestHandlerCreateOk(t *testing.T) {
	// Arrange
	database := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "232345",
			FirstName:    "hello",
			LastName:     "World",
		},
	}

	requestBodyJSON := domain.Buyer{
		CardNumberID: "232347",
		FirstName:    "hello_2",
		LastName:     "World_2",
	}

	productToSave, err := json.Marshal(&requestBodyJSON)
	if err != nil {
		panic(err)
	}

	mockService := buyer.MockService{
		DataMock: database,
	}

	var resp map[string]domain.Buyer

	// Act
	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodPost, "/api/v1/buyers", string(productToSave))
	router.ServeHTTP(rr, req)
	err = json.Unmarshal(rr.Body.Bytes(), &resp)

	// fmt.Printf("--- %+v\n", rr)
	// fmt.Printf("--- %+v\n", err)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Equal(t, len(mockService.DataMock)+1, resp["data"].ID)
}

func TestHandlerCreateFail(t *testing.T) {
	// Arrange
	errMessage := "Key: 'RequestBuyer.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag"
	database := []domain.Buyer{}

	// Missing "FirstName" field
	requestBodyJSON := domain.Buyer{
		CardNumberID: "232347",
		LastName:     "World_2",
	}

	productToSave, err := json.Marshal(&requestBodyJSON)
	if err != nil {
		panic(err)
	}

	mockService := buyer.MockService{
		DataMock: database,
		Error:    errMessage,
	}

	var resp map[string]string

	// Act
	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodPost, "/api/v1/buyers", string(productToSave))
	router.ServeHTTP(rr, req)
	err = json.Unmarshal(rr.Body.Bytes(), &resp)

	// fmt.Printf("--- %+v\n", rr)
	// fmt.Printf("--- %+v\n", err)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)
	assert.Equal(t, errMessage, resp["message"])
}

func TestHandlerCreateConflict(t *testing.T) {
	// Arrange
	errMessage := "duplicate cardNumberID"
	cardNumberIDRepeated := "232345"
	database := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: cardNumberIDRepeated,
			FirstName:    "hello",
			LastName:     "World",
		},
	}

	// Missing "FirstName" field
	requestBodyJSON := domain.Buyer{
		CardNumberID: cardNumberIDRepeated,
		FirstName:    "hello_2",
		LastName:     "World_2",
	}

	productToSave, err := json.Marshal(&requestBodyJSON)
	if err != nil {
		panic(err)
	}

	mockService := buyer.MockService{
		DataMock: database,
		Error:    errMessage,
	}

	var resp map[string]string

	// Act
	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodPost, "/api/v1/buyers", string(productToSave))
	router.ServeHTTP(rr, req)
	err = json.Unmarshal(rr.Body.Bytes(), &resp)

	// fmt.Printf("--- %+v\n", rr)
	// fmt.Printf("--- %+v\n", err)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusConflict, rr.Code)
	assert.Equal(t, errMessage, resp["message"])
}

func TestHandlerUpdateOk(t *testing.T) {
	// Arrange
	beforeBuyer := domain.Buyer{
		ID:           1,
		CardNumberID: "23472398",
		FirstName:    "hello",
		LastName:     "World",
	}

	afterUpdate := domain.Buyer{
		ID:           1,
		CardNumberID: "23472398",
		FirstName:    "hello_2",
		LastName:     "World",
	}

	mockService := buyer.MockService{
		DataMock: []domain.Buyer{
			beforeBuyer,
		},
	}

	var resp map[string]domain.Buyer

	productToUpdate, err := json.Marshal(&afterUpdate)
	if err != nil {
		panic(err)
	}

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodPatch, "/api/v1/buyers/1", string(productToUpdate))

	// Act
	router.ServeHTTP(rr, req)
	err = json.Unmarshal(rr.Body.Bytes(), &resp)

	// fmt.Printf("--- %+v\n", rr)
	// fmt.Printf("--- %+v\n", err)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
	assert.Equal(t, afterUpdate, resp["data"])
}

func TestHandlerUpdateNotExisting(t *testing.T) {
	// Arrange
	errMessage := "buyer not found"

	afterUpdate := domain.Buyer{
		CardNumberID: "23472398",
		FirstName:    "hello_2",
		LastName:     "World",
	}

	mockService := buyer.MockService{
		DataMock: []domain.Buyer{},
		Error:    errMessage,
	}

	var resp map[string]string

	productToUpdate, err := json.Marshal(&afterUpdate)
	if err != nil {
		panic(err)
	}

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodPatch, "/api/v1/buyers/1", string(productToUpdate))

	// Act
	router.ServeHTTP(rr, req)
	err = json.Unmarshal(rr.Body.Bytes(), &resp)

	// fmt.Printf("--- %+v\n", rr)
	// fmt.Printf("--- %+v\n", err)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, rr.Code)
	assert.Equal(t, errMessage, resp["message"])
}

func TestHandlerUpdateBadId(t *testing.T) {
	// Arrange
	errMessage := "strconv.Atoi: parsing \"bar\": invalid syntax"

	afterUpdate := domain.Buyer{
		CardNumberID: "23472398",
		FirstName:    "hello_2",
		LastName:     "World",
	}

	mockService := buyer.MockService{
		DataMock: []domain.Buyer{},
		Error:    errMessage,
	}

	var resp map[string]string

	productToUpdate, err := json.Marshal(&afterUpdate)
	if err != nil {
		panic(err)
	}

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodPatch, "/api/v1/buyers/bar", string(productToUpdate))

	// Act
	router.ServeHTTP(rr, req)
	err = json.Unmarshal(rr.Body.Bytes(), &resp)

	// fmt.Printf("--- %+v\n", rr)
	// fmt.Printf("--- %+v\n", err)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, errMessage, resp["message"])
}

func TestHandlerDeleteOk(t *testing.T) {
	// arrange
	database := []domain.Buyer{
		{
			ID:           1,
			CardNumberID: "232345",
			FirstName:    "hello",
			LastName:     "World",
		},
	}

	// lenBeforeDatabase := len(database)

	dbCopy := make([]domain.Buyer, len(database))
	copy(dbCopy, database)
	// dbCopy[0].ID = 0

	// t.Log(database)
	// t.Log(dbCopy)

	mockService := buyer.MockService{
		DataMock: dbCopy,
	}

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodDelete, "/api/v1/buyers/1", "")

	// Act
	router.ServeHTTP(rr, req)

	t.Logf(" %v database %p\n", database, database)
	t.Logf(" %v mockService DataMock %p\n", mockService.DataMock, mockService.DataMock)

	// Assert
	assert.Equal(t, http.StatusNoContent, rr.Code)
	// assert.NotEqual(t, database, mockService.DataMock)
	// assert.NotEqual(t, lenBeforeDatabase, len(mockService.DataMock))
}

func TestHandlerDeleteNotExistent(t *testing.T) {
	// arrange
	errMessage := "buyer not found"
	database := []domain.Buyer{}

	mockService := buyer.MockService{
		DataMock: database,
		Error:    errMessage,
	}

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodDelete, "/api/v1/buyers/1", "")

	// Act
	router.ServeHTTP(rr, req)

	// fmt.Printf("--- %+v\n", rr)
	// fmt.Printf("--- %+v\n", err)

	// Assert
	assert.Equal(t, http.StatusNotFound, rr.Code)
	// assert.NotEqual(t, len(database), len(mockService.DataMock))
}

func TestHandlerDeleteBadId(t *testing.T) {
	// arrange
	errMessage := "strconv.Atoi: parsing \"bar\": invalid syntax"
	database := []domain.Buyer{}

	mockService := buyer.MockService{
		DataMock: database,
		Error:    errMessage,
	}

	router := createServer(mockService)
	req, rr := createRequestTest(http.MethodDelete, "/api/v1/buyers/bar", "")

	// Act
	router.ServeHTTP(rr, req)

	// fmt.Printf("--- %+v\n", rr)
	// fmt.Printf("--- %+v\n", err)

	// Assert
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	// assert.NotEqual(t, len(database), len(mockService.DataMock))
}
