package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/products"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var database = []domain.Product{
	{
		ID:             1,
		Description:    "Producto congelado",
		ExpirationRate: 10,
		FreezingRate:   10,
		Height:         10,
		Length:         10,
		Netweight:      10,
		ProductCode:    "567",
		RecomFreezTemp: 5,
		Width:          10,
		ProductTypeID:  12345,
		SellerID:       6,
	},
	{
		ID:             2,
		Description:    "Producto refrigerado",
		ExpirationRate: 20,
		FreezingRate:   9,
		Height:         20,
		Length:         20,
		Netweight:      100,
		ProductCode:    "980",
		RecomFreezTemp: 6,
		Width:          20,
		ProductTypeID:  12345,
		SellerID:       7,
	},
}

func createServerProduct(mockService *products.MockServiceProduct) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	handler := NewProduct(mockService)
	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", handler.Create())
	pr.GET("/", handler.GetAll())
	pr.GET("/:id", handler.Get())
	pr.PATCH("/:id", handler.Update())
	pr.DELETE("/:id", handler.Delete())

	return r
}

func createRequestTestProduct(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application-json")
	return req, httptest.NewRecorder()
}

// User Story asociada: CREATE
// Caso borde: create_ok
// Cuando el ingreso de datos sea exitoso se devolverá un código 201 junto con el objeto ingresado.
// Status code: 201
func TestCreateOkProduct(t *testing.T) {

	//Arrange
	pr := domain.Product{
		ID:             3,
		Description:    "Producto de aseo personal",
		ExpirationRate: 30,
		FreezingRate:   90,
		Height:         20.0,
		Length:         20.0,
		Netweight:      100.0,
		ProductCode:    "j4j5j7",
		RecomFreezTemp: 0.0,
		Width:          109.0,
		ProductTypeID:  123845,
		SellerID:       9,
	}
	var productNew map[string]domain.Product

	mockService := products.MockServiceProduct{
		DataMock: database,
	}

	//Act
	server := createServerProduct(&mockService)
	request, received := createRequestTestProduct(http.MethodPost, "/products/",
		`{
            "description": "Producto de aseo personal",
            "expiration_rate": 30,
            "freezing_rate": 90,
            "height": 20.0,
            "length": 20.0,
            "netweight": 100.0,
            "product_code": "j4j5j7",
            "recommended_freezing_temperature": 0.0,
            "width": 109.0,
            "product_type_id": 123845,
            "seller_id": 9
        }`)

	server.ServeHTTP(received, request)
	t.Log(received.Body.String())
	err := json.Unmarshal(received.Body.Bytes(), &productNew)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 201, received.Code)
	assert.Equal(t, pr, productNew["data"])
}

// User Story asociada: CREATE
// Caso borde: create_fail
// Si el objeto JSON no contiene los campos necesarios se devolverá un código 422
// Status code: 422
func TestCreateFailProduct(t *testing.T) {

	//Arrange
	expectedError := "description is required"
	var errorResult map[string]string

	mockService := products.MockServiceProduct{
		DataMock: database,
	}

	//Act
	server := createServerProduct(&mockService)
	request, received := createRequestTestProduct(http.MethodPost, "/products/",
		`{
            "expiration_rate": 30,
            "freezing_rate": 90,
            "height": 20,
            "length": 20,
            "netweight": 100,
            "product_code": "j4j5j6",
            "recommended_freezing_temperature": 0,
            "width": 109,
            "product_type_id": 123845,
            "seller_id": 9
        }`)
	server.ServeHTTP(received, request)
	err := json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 422, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])

	//Arrange
	expectedError = "product code is required"

	//Act
	request, received = createRequestTestProduct(http.MethodPost, "/products/",
		`{
            "description": "Producto congelado",
            "expiration_rate": 10,
            "freezing_rate": 10,
            "height": 10,
            "length": 10,
            "netweight": 10,
            "recommended_freezing_temperature": 5,
            "width": 10,
            "product_type_id": 12345,
            "seller_id": 6
        }`)
	server.ServeHTTP(received, request)
	err = json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 422, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])

	//Arrange
	expectedError = "expiration_rate is required"

	//Act
	request, received = createRequestTestProduct(http.MethodPost, "/products/",
		`{
            "description": "Producto congelado",
            "freezing_rate": 10,
            "height": 10,
            "length": 10,
            "netweight": 10,
            "product_code": "j4j5j6",
            "recommended_freezing_temperature": 5,
            "width": 10,
            "product_type_id": 12345,
            "seller_id": 6
        }`)
	server.ServeHTTP(received, request)
	err = json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 422, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 422, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])

	//Arrange
	expectedError = "freezing_rate is required"

	//Act
	request, received = createRequestTestProduct(http.MethodPost, "/products/",
		`{
            "description": "Producto congelado",
            "expiration_rate": 10,
            "height": 10,
            "length": 10,
            "netweight": 10,
            "product_code": "j4j5j6",
            "recommended_freezing_temperature": 5,
            "width": 10,
            "product_type_id": 12345,
            "seller_id": 6
        }`)
	server.ServeHTTP(received, request)
	err = json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 422, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])

	//Arrange
	expectedError = "height is required"

	//Act
	request, received = createRequestTestProduct(http.MethodPost, "/products/",
		`{
            "description": "Producto congelado",
            "expiration_rate": 10,
            "freezing_rate": 10,
            "length": 10,
            "netweight": 10,
            "product_code": "j4j5j6",
            "recommended_freezing_temperature": 5,
            "width": 10,
            "product_type_id": 12345,
            "seller_id": 6
        }`)
	server.ServeHTTP(received, request)
	err = json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 422, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 422, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])

	//Arrange
	expectedError = "length is required"

	//Act
	request, received = createRequestTestProduct(http.MethodPost, "/products/",
		`{
            "description": "Producto congelado",
            "expiration_rate": 10,
            "freezing_rate": 10,
            "height": 10,
            "netweight": 10,
            "product_code": "j4j5j6",
            "recommended_freezing_temperature": 5,
            "width": 10,
            "product_type_id": 12345,
            "seller_id": 6
        }`)
	server.ServeHTTP(received, request)
	err = json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 422, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])

	//Arrange
	expectedError = "netweight is required"

	//Act
	request, received = createRequestTestProduct(http.MethodPost, "/products/",
		`{
            "description": "Producto congelado",
            "expiration_rate": 10,
            "freezing_rate": 10,
            "height": 10,
            "length": 10,
            "product_code": "j4j5j6",
            "recommended_freezing_temperature": 5,
            "width": 10,
            "product_type_id": 12345,
            "seller_id": 6
        }`)
	server.ServeHTTP(received, request)
	err = json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 422, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])

	//Arrange
	expectedError = "product code is required"

	//Act
	request, received = createRequestTestProduct(http.MethodPost, "/products/",
		`{
            "description": "Producto congelado",
            "expiration_rate": 10,
            "freezing_rate": 10,
            "height": 10,
            "length": 10,
            "netweight": 10,
            "recommended_freezing_temperature": 5,
            "width": 10,
            "product_type_id": 12345,
            "seller_id": 6
        }`)
	server.ServeHTTP(received, request)
	err = json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 422, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])

	//Arrange
	expectedError = "recommended_freezing_temperature is required"

	//Act
	request, received = createRequestTestProduct(http.MethodPost, "/products/",
		`{
            "description": "Producto congelado",
            "expiration_rate": 10,
            "freezing_rate": 10,
            "height": 10,
            "length": 10,
            "netweight": 10,
            "product_code": "j4j5j6",
            "width": 10,
            "product_type_id": 12345,
            "seller_id": 6
        }`)
	server.ServeHTTP(received, request)
	err = json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 422, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])

	//Arrange
	expectedError = "width is required"

	//Act
	request, received = createRequestTestProduct(http.MethodPost, "/products/",
		`{
            "description": "Producto congelado",
            "expiration_rate": 10,
            "freezing_rate": 10,
            "height": 10,
            "length": 10,
            "netweight": 10,
            "product_code": "j4j5j6",
            "recommended_freezing_temperature": 5,
            "product_type_id": 12345,
            "seller_id": 6
        }`)
	server.ServeHTTP(received, request)
	err = json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 422, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])

}

// User Story asociada: CREATE, Caso borde: create_conflict, Si el product_code ya existe devuelve un error 409 Conflict
func TestCreateConflictProduct(t *testing.T) {

	//Arrange
	expectedError := "product_code already exists"
	var errorResult map[string]string

	mockService := products.MockServiceProduct{
		DataMock: database,
	}

	//Act
	server := createServerProduct(&mockService)
	request, received := createRequestTestProduct(http.MethodPost, "/products/",
		`{
            "description": "Producto congelado",
            "expiration_rate": 10,
            "freezing_rate": 10,
            "height": 10,
            "length": 10,
            "netweight": 10,
            "product_code": "567",
            "recommended_freezing_temperature": 5,
            "width": 10,
            "product_type_id": 12345,
            "seller_id": 6
        }`)
	server.ServeHTTP(received, request)
	err := json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 409, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])
}

// User Story asociada: READ, Caso borde: find_all, Cuando la petición sea exitosa el backend devolverá un listado de todas los productos existentes
func TestFindAllProduct(t *testing.T) {
	//Arrange
	productsExpected := database
	var productsResult map[string][]domain.Product

	mockService := products.MockServiceProduct{
		DataMock: database,
	}

	//Act
	server := createServerProduct(&mockService)
	request, received := createRequestTestProduct(http.MethodGet, "/products/", "")
	server.ServeHTTP(received, request)
	err := json.Unmarshal(received.Body.Bytes(), &productsResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 200, received.Code)
	assert.Equal(t, productsExpected, productsResult["data"])

}

func TestFindAllProductFail(t *testing.T) {
	expectedError := "get all error"
	var errorResult map[string]string

	mockService := products.MockServiceProduct{DataMock: database, Error: expectedError}
	server := createServerProduct(&mockService)
	request, received := createRequestTestProduct(http.MethodGet, "/products/", "")

	server.ServeHTTP(received, request)

	err := json.Unmarshal(received.Body.Bytes(), &errorResult)

	assert.Nil(t, err)
	assert.Equal(t, 404, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])
}

// User Story asociada: READ, Caso borde: find_by_id_non_existent, Cuando el producto no exista se devolverá un código 404
func TestFindByIdNonExistentProduct(t *testing.T) {
	//Arrange
	expectedError := "product not found"
	var errorResult map[string]string
	mockService := products.MockServiceProduct{
		DataMock: database,
	}

	//Act
	server := createServerProduct(&mockService)
	request, received := createRequestTestProduct(http.MethodGet, "/products/5", "")
	server.ServeHTTP(received, request)
	err := json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 404, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])

	//Output:

}

// User Story asociada: READ
// Caso borde: find_by_id_existent
// Cuando la petición sea exitosa el backend devolverá la información del producto solicitado
// Status code: 200
func TestFindByIdExistentProduct(t *testing.T) {
	//Arrange
	productExpected := domain.Product{
		ID:             1,
		Description:    "Producto congelado",
		ExpirationRate: 10,
		FreezingRate:   10,
		Height:         10,
		Length:         10,
		Netweight:      10,
		ProductCode:    "567",
		RecomFreezTemp: 5,
		Width:          10,
		ProductTypeID:  12345,
		SellerID:       6,
	}

	var productResult map[string]domain.Product
	mockService := products.MockServiceProduct{
		DataMock: database,
	}

	//Act
	server := createServerProduct(&mockService)
	request, received := createRequestTestProduct(http.MethodGet, "/products/1", "")

	server.ServeHTTP(received, request)
	err := json.Unmarshal(received.Body.Bytes(), &productResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 200, received.Code)
	assert.Equal(t, productExpected, productResult["data"])
}

// User Story asociada: UPDATE
// Caso borde: update_ok
// Cuando la actualización de datos sea exitosa se devolverá el producto con la información
// actualizada junto con un código 200
// Status code: 200
func TestUpdateOkProduct(t *testing.T) {
	//Arrange
	productExpected := domain.Product{
		ID:             1,
		Description:    "Producto congelado",
		ExpirationRate: 10,
		FreezingRate:   10,
		Height:         20.0,
		Length:         20.0,
		Netweight:      10,
		ProductCode:    "567",
		RecomFreezTemp: 5,
		Width:          10,
		ProductTypeID:  12345,
		SellerID:       6,
	}

	var productResult map[string]domain.Product

	mockService := products.MockServiceProduct{
		DataMock: database,
	}

	//Act
	server := createServerProduct(&mockService)
	request, received := createRequestTestProduct(http.MethodPatch, "/products/1",
		`{
            "height": 20.0,
            "length": 20.0
        }`)
	server.ServeHTTP(received, request)
	t.Log(received.Body.String())
	err := json.Unmarshal(received.Body.Bytes(), &productResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 200, received.Code)
	assert.Equal(t, productExpected, productResult["data"])
}

// User Story asociada: UPDATE
// Caso borde: update_non_existent
// Si el producto que se desea actualizar no existe se devolverá un código 404
// Status code: 404
func TestUpdateNonExistentProduct(t *testing.T) {
	//Arrange
	expectedError := "product not found"
	var errorResult map[string]string

	mockService := products.MockServiceProduct{
		DataMock: database,
	}

	//Act
	server := createServerProduct(&mockService)
	request, received := createRequestTestProduct(http.MethodPatch, "/products/6",
		`{
        "height": 20.0,
        "length": 20.0
    }`)

	server.ServeHTTP(received, request)

	err := json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 404, received.Code)
	assert.Equal(t, expectedError, errorResult["message"])
}

// User Story asociada: DELETE
// Caso borde: delete_non_existent
// Cuando el producto no existe se devolverá un código 404
// Status code: 404
func TestDeleteNonExistentProduct(t *testing.T) {
	//Arrange
	errorExpected := "product not found"
	var errorResult map[string]string

	mockService := products.MockServiceProduct{
		DataMock: database,
	}

	//Act
	server := createServerProduct(&mockService)
	request, received := createRequestTestProduct(http.MethodDelete, "/products/9", "")
	server.ServeHTTP(received, request)
	err := json.Unmarshal(received.Body.Bytes(), &errorResult)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, 404, received.Code)
	assert.Equal(t, errorExpected, errorResult["message"])
}

// User Story asociada: DELETE
// Caso borde: delete_ok
// Cuando la eliminación sea exitosa se devolverá un código 204
// Status code: 204
func TestDeleteOkProduct(t *testing.T) {

	//append([]domain.Product{}, database...)

	//Arrange
	dataCopy := make([]domain.Product, len(database))
	copy(dataCopy, database)

	mockService := products.MockServiceProduct{
		DataMock: database,
	}

	//Act
	server := createServerProduct(&mockService)
	request, received := createRequestTestProduct(http.MethodDelete, "/products/1", "")
	server.ServeHTTP(received, request)
	t.Log(dataCopy)
	t.Log(mockService.DataMock)

	//Assert
	assert.Equal(t, 204, received.Code)
	assert.NotEqual(t, dataCopy, mockService.DataMock)
}
