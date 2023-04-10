package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/locality"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/pkg/web"
)

type requestLocality struct {
	ID           int    `json:"locality_id"`
	LocalityName string `json:"locality_name"`
	ProvinceName string `json:"province_name"`
	CountryName  string `json:"country_name"`
}

type Locality struct {
	localityService locality.Service
}

func NewLocality(l locality.Service) *Locality {
	return &Locality{
		localityService: l,
	}
}

// CreateLocality godoc
// @Summary Create locality
// @Tags Localities
// @Description create locality
// @Accept  json
// @Produce  json
// @Param locality body domain.Locality true "Locality"
// @Success 201 {object} web.response
// @Router /api/v1/localities [post]
func (l *Locality) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestLocality
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		if req.ID <= 0 {
			web.Error(c, http.StatusUnprocessableEntity, locality.ErrRequired.Error()+": id")
			return
		}
		if req.LocalityName == "" {
			web.Error(c, http.StatusUnprocessableEntity, locality.ErrRequired.Error()+": locality_name")
			return
		}
		if req.ProvinceName == "" {
			web.Error(c, http.StatusUnprocessableEntity, locality.ErrRequired.Error()+": province_name")
			return
		}
		if req.CountryName == "" {
			web.Error(c, http.StatusUnprocessableEntity, locality.ErrRequired.Error()+": country_name")
			return
		}

		new, err := l.localityService.Create(c, domain.Locality(req))
		if err != nil {
			if err.Error() == "id already exists" {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
			web.Error(c, 422, err.Error())
			return
		}

		web.Success(c, http.StatusCreated, new)
	}
}

func (l *Locality) GetAllSellersByLocality() gin.HandlerFunc {
	return func(c *gin.Context) {
		localityID := c.Query("id")

		reports, err := l.localityService.GetAllSellersByLocality(c, localityID)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web.Success(c, http.StatusOK, reports)
	}
}

// GetLocalityReport godoc
// @Summary Get locality report
// @Tags Localities
// @Description get locality report
// @Accept  json
// @Produce  json
// @Param id query string true "Locality id"
// @Success 200 {object} web.response
// @Router /api/v1/localities/reportCarries [get]
func (l *Locality) GetReport() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		reports, err := l.localityService.GetCarriesReport(c, id)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		web.Success(c, http.StatusOK, reports)
	}
}
