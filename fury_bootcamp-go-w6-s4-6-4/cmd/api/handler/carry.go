package handler

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/carry"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/pkg/web"
)

type Carry struct {
	carryService carry.Service
}

func NewCarry(c carry.Service) *Carry {
	return &Carry{
		carryService: c,
	}
}

// CreateCarry godoc
// @Summary Create a carry
// @Tags Carries
// @Description create a carry
// @Accept  json
// @Produce  json
// @Param carry body domain.Carry true "Carry"
// @Success 201 {object} web.response
// @Router /api/v1/carries [post]
func (c *Carry) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		carry := domain.Carry{}
		if err := ctx.ShouldBindJSON(&carry); err != nil {
			web.Error(ctx, http.StatusUnprocessableEntity, err.Error())
			return
		}

		emptyFields := []string{}
		values := reflect.ValueOf(carry)
		for i := 0; i < values.NumField(); i++ {
			if values.Type().Field(i).Name == "ID" || values.Type().Field(i).Name == "Locality_id" {
				continue
			}
			if values.Field(i).IsZero() {
				emptyFields = append(emptyFields, values.Type().Field(i).Name)
			}
		}

		if len(emptyFields) > 0 {
			web.Error(ctx, http.StatusUnprocessableEntity, "empty fields: "+strings.Join(emptyFields, ", "))
			return
		}

		id, err := c.carryService.Save(carry)
		if err != nil {
			web.Error(ctx, http.StatusConflict, err.Error())
			return
		}

		carry.ID = id

		web.Success(ctx, http.StatusCreated, carry)
	}
}
