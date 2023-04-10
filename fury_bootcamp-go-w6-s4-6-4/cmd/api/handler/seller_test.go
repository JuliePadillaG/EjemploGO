package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/sellers"
	"github.com/stretchr/testify/assert"
)

func createServerSeller(mockService sellers.MockService) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	p := NewSeller(&mockService)
	r := gin.Default()
	pr := r.Group("api/v1/sellers")
	pr.GET("", p.GetAll())
	pr.GET("/:id", p.Get())
	pr.POST("", p.Create())
	pr.PATCH("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	return r
}
func createRequestTestSeller(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}
func TestCreateOk(t *testing.T) {
	// arrange
	mockService := sellers.MockService{
		DataMock: []domain.Seller{},
		Error:    "",
	}
	var resp domain.Seller
	r := createServerSeller(mockService)
	req, rr := createRequestTestSeller(http.MethodPost, "/api/v1/sellers", `{
        "cid": 34,"company_name": "LG","address": "Avenida 11122","telephone": "0303456", "locality_id": 1759
        }`)
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 201, rr.Code)
}
func TestCreateBadRequest(t *testing.T) {
	// arrange
	mockService := sellers.MockService{
		DataMock: []domain.Seller{},
		Error:    "",
	}
	var resp domain.Seller
	r := createServerSeller(mockService)
	req, rr := createRequestTestSeller(http.MethodPost, "/api/v1/sellers", `{
        "cid": -34,"company_name": "as","address": "Avenida 11122","telephone": "0303456", "locality_id": 1759
        }`)
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 400, rr.Code)
}
func TestCreateFail(t *testing.T) {
	// arrange
	mockService := sellers.MockService{
		DataMock: []domain.Seller{},
		Error:    "",
	}
	var resp domain.Seller
	r := createServerSeller(mockService)
	req, rr := createRequestTestSeller(http.MethodPost, "/api/v1/sellers", `{
        "cid":34,"address":"Avenida 11122","telephone": "0303456"
        }`)
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 422, rr.Code)
}
func TestCreateConflict(t *testing.T) {
	// arrange
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}
	mockService := sellers.MockService{
		DataMock: database,
		Error:    "",
	}
	var resp domain.Seller
	r := createServerSeller(mockService)
	req, rr := createRequestTestSeller(http.MethodPost, "/api/v1/sellers", `{
        "cid":19,"address":"Avenida 11122","telephone": "0303456"
        }`)
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 409, rr.Code)
}
func TestFindAll(t *testing.T) {
	// arrange
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}
	expected := database
	mockService := sellers.MockService{
		DataMock: database,
		Error:    "",
	}
	var resp map[string][]domain.Seller
	r := createServerSeller(mockService)
	req, rr := createRequestTestSeller(http.MethodGet, "/api/v1/sellers", "")
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, expected, resp["data"])
}
func TestFindAllFail(t *testing.T) {
	// arrange
	database := []domain.Seller{}
	mockService := sellers.MockService{
		DataMock: database,
		Error:    "empty",
	}
	var resp domain.Seller
	r := createServerSeller(mockService)
	req, rr := createRequestTestSeller(http.MethodGet, "/api/v1/sellers", "")
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 404, rr.Code)
}
func TestFindByIdNonExistent(t *testing.T) {
	// arrange
	database := []domain.Seller{}
	mockService := sellers.MockService{
		DataMock: database,
		Error:    "",
	}
	var resp domain.Seller
	r := createServerSeller(mockService)
	req, rr := createRequestTestSeller(http.MethodGet, "/api/v1/sellers/1", "")
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 404, rr.Code)
}
func TestFindByIdExistent(t *testing.T) {
	// arrange
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}
	var expected = domain.Seller{
		ID:          1,
		CID:         19,
		CompanyName: "LG",
		Address:     "Avenida 11122",
		Telephone:   "0303456",
	}
	mockService := sellers.MockService{
		DataMock: database,
		Error:    "",
	}
	var resp map[string]domain.Seller
	r := createServerSeller(mockService)
	req, rr := createRequestTestSeller(http.MethodGet, "/api/v1/sellers/1", "")
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, expected, resp["data"])
}
func TestUpdateOK(t *testing.T) {
	// arrange
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}
	var expected = domain.Seller{
		ID:          1,
		CID:         19,
		CompanyName: "Samsung",
		Address:     "Avenida 11122",
		Telephone:   "0303456",
	}
	mockService := sellers.MockService{
		DataMock: database,
		Error:    "",
	}
	var resp map[string]domain.Seller
	r := createServerSeller(mockService)
	req, rr := createRequestTestSeller(http.MethodPatch, "/api/v1/sellers/1", `{
        "company_name": "Samsung"
        }`)
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 200, rr.Code)
	assert.Equal(t, expected, resp["data"])
}
func TestUpdateNonExistent(t *testing.T) {
	// arrange
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}
	mockService := sellers.MockService{
		DataMock: database,
		Error:    "",
	}
	var resp domain.Seller
	r := createServerSeller(mockService)
	req, rr := createRequestTestSeller(http.MethodPatch, "/api/v1/sellers/2", `{
        "company_name": "Samsung"
        }`)
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 404, rr.Code)
}
func TestDeleteNonExistent(t *testing.T) {
	// arrange
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}
	mockService := sellers.MockService{
		DataMock: database,
		Error:    "",
	}
	var resp domain.Seller
	r := createServerSeller(mockService)
	req, rr := createRequestTestSeller(http.MethodDelete, "/api/v1/sellers/2", "")
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 404, rr.Code)
}
func TestDeleteOK(t *testing.T) {
	// arrange
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}
	mockService := sellers.MockService{
		DataMock: database,
		Error:    "",
	}
	var resp domain.Seller
	r := createServerSeller(mockService)
	req, rr := createRequestTestSeller(http.MethodDelete, "/api/v1/sellers/1", "")
	// act
	r.ServeHTTP(rr, req)
	// assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, 200, rr.Code)
	assert.Empty(t, resp)
}
