package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/locality"
	"github.com/stretchr/testify/assert"
)

func createServerLocality(mockService locality.MockService) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	p := NewLocality(&mockService)
	r := gin.Default()
	pr := r.Group("api/v1/localities")
	pr.GET("/reportSellers", p.GetAllSellersByLocality())
	pr.GET("/reportCarries", p.GetReport())
	pr.POST("", p.Create())
	return r
}
func createRequestTestLocality(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}
func TestCreateLocality(t *testing.T) {
	// arrange
	mockService := locality.MockService{
		DataMock: []domain.Locality{{
			ID:           1759,
			LocalityName: "Gonzalez Catan",
			ProvinceName: "Buenos Aires",
			CountryName:  "Argentina",
		}},
		Error: "",
	}
	var resp domain.Locality
	r := createServerLocality(mockService)

	t.Run("create ok", func(t *testing.T) {
		req, rr := createRequestTestSeller(http.MethodPost, "/api/v1/localities", `{
		"locality_id": 1755,
		"locality_name": "Moron",
		"province_name": "Buenos Aires",
		"country_name": "Argentina"
	  }s`)
		// act
		r.ServeHTTP(rr, req)
		// assert
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, 201, rr.Code)
	})

	t.Run("create id exists", func(t *testing.T) {
		req, rr := createRequestTestSeller(http.MethodPost, "/api/v1/localities", `{
		"locality_id": 1759,
		"locality_name": "Moron",
		"province_name": "Buenos Aires",
		"country_name": "Argentina"
	  }s`)
		// act
		r.ServeHTTP(rr, req)
		// assert
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, 409, rr.Code)
	})

	t.Run("create fail", func(t *testing.T) {
		req, rr := createRequestTestSeller(http.MethodPost, "/api/v1/localities", `{
		"locality_id": ,
		"locality_name": "Moron",
		"province_name": "Buenos Aires",
		"country_name": "Argentina"
	  }s`)
		// act
		r.ServeHTTP(rr, req)
		// assert
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, 404, rr.Code)
	})

	t.Run("create field locality required", func(t *testing.T) {
		req, rr := createRequestTestSeller(http.MethodPost, "/api/v1/localities", `{
		"locality_id": 12,
		"locality_name": "",
		"province_name": "Buenos Aires",
		"country_name": "Argentina"
	  }s`)
		// act
		r.ServeHTTP(rr, req)
		// assert
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, 422, rr.Code)
	})

	t.Run("create field province required", func(t *testing.T) {
		req, rr := createRequestTestSeller(http.MethodPost, "/api/v1/localities", `{
		"locality_id": 12,
		"locality_name": "San Justo",
		"province_name": "",
		"country_name": "Argentina"
	  }s`)
		// act
		r.ServeHTTP(rr, req)
		// assert
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, 422, rr.Code)
	})

	t.Run("create field country required", func(t *testing.T) {
		req, rr := createRequestTestSeller(http.MethodPost, "/api/v1/localities", `{
		"locality_id": 12,
		"locality_name": "San Justo",
		"province_name": "Buenos Aires",
		"country_name": ""
	  }s`)
		// act
		r.ServeHTTP(rr, req)
		// assert
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, 422, rr.Code)
	})

	t.Run("create field id required", func(t *testing.T) {
		req, rr := createRequestTestSeller(http.MethodPost, "/api/v1/localities", `{
		"locality_id": 0,
		"locality_name": "San Justo",
		"province_name": "Buenos Aires",
		"country_name": "Argentis"
	  }s`)
		// act
		r.ServeHTTP(rr, req)
		// assert
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, 422, rr.Code)
	})
}

func TestGetSellersLocality(t *testing.T) {
	// arrange
	mockService := locality.MockService{
		DataMock: []domain.Locality{{
			ID:           1759,
			LocalityName: "Gonzalez Catan",
			ProvinceName: "Buenos Aires",
			CountryName:  "Argentina",
		}},
		Error: "",
	}
	var resp domain.Locality
	r := createServerLocality(mockService)

	t.Run("get ok", func(t *testing.T) {
		req, rr := createRequestTestSeller(http.MethodGet, "/api/v1/localities/reportSellers", "")
		// act
		r.ServeHTTP(rr, req)
		// assert
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, 200, rr.Code)
	})

	t.Run("get fail", func(t *testing.T) {
		req, rr := createRequestTestSeller(http.MethodGet, "/api/v1/localities/reportSellers?id=1759", "")
		// act
		r.ServeHTTP(rr, req)
		// assert
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, 404, rr.Code)
	})
}

func TestGetCarriesLocality(t *testing.T) {
	// arrange
	mockService := locality.MockService{
		DataMock: []domain.Locality{{
			ID:           1759,
			LocalityName: "Gonzalez Catan",
			ProvinceName: "Buenos Aires",
			CountryName:  "Argentina",
		}},
		Error: "",
	}
	var resp domain.Locality
	r := createServerLocality(mockService)

	t.Run("get ok", func(t *testing.T) {
		req, rr := createRequestTestSeller(http.MethodGet, "/api/v1/localities/reportCarries", "")
		// act
		r.ServeHTTP(rr, req)
		// assert
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, 200, rr.Code)
	})

	t.Run("get fail", func(t *testing.T) {
		req, rr := createRequestTestSeller(http.MethodGet, "/api/v1/localities/reportCarries?id=1759", "")
		// act
		r.ServeHTTP(rr, req)
		// assert
		err := json.Unmarshal(rr.Body.Bytes(), &resp)
		assert.Nil(t, err)
		assert.Equal(t, 404, rr.Code)
	})
}
