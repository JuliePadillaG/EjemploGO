package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/buyer"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/pkg/web"
)

var (
        ErrNotFoundBuyers = errors.New("there is no buyers")
)

type RequestBuyer struct {
        ID           int    `json:"id"`
	CardNumberID string `json:"card_number_id" binding:"required"`
	FirstName    string `json:"first_name"     binding:"required"`
	LastName     string `json:"last_name"      binding:"required"`
}

type RequestPatchBuyer struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}


type Buyer struct {
	buyerService buyer.Service
}

func NewBuyer(b buyer.Service) *Buyer {
	return &Buyer{
		buyerService: b,
	}
}

func (b *Buyer) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
                        web.Error(ctx, http.StatusBadRequest, err.Error())
			return
		}

                buyer, err := b.buyerService.Get(ctx, id)
                if err != nil {
                        web.Error(ctx, http.StatusNotFound, ErrNotFoundBuyers.Error())
                        return
                }

	        var buyers []domain.Buyer

		buyers = append(buyers, buyer)
                web.Success(ctx, http.StatusOK, buyers)
        }
}

func (b *Buyer) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

                buyers, err := b.buyerService.GetAll(ctx)
	        if err != nil {
                        web.Error(ctx, http.StatusUnprocessableEntity, err.Error())
	        	return
	        }

	        if len(buyers) == 0 {
                        web.Error(ctx, http.StatusNotFound, ErrNotFoundBuyers.Error())
	        	return
	        }

                web.Success(ctx, http.StatusOK, buyers)
        }
}

func (b *Buyer) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
                var req RequestBuyer

                if err := ctx.ShouldBindJSON(&req); err != nil {
                        web.Error(ctx, http.StatusUnprocessableEntity, err.Error())
                        return
		}

		buyer, err := b.buyerService.Save(ctx, req.CardNumberID, req.FirstName, req.LastName)
		if err != nil {
                        web.Error(ctx, http.StatusConflict, err.Error())
			return
		}

                web.Success(ctx, http.StatusCreated, buyer)
        }
}

func (b *Buyer) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

                id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, err.Error())
			return
		}

		req := RequestPatchBuyer{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Error(ctx, http.StatusBadRequest, err.Error())
			return
		}

                buyerUpdated, err := b.buyerService.Update(ctx, id, req.FirstName, req.LastName)
                if err != nil {
			web.Error(ctx, http.StatusNotFound, err.Error())
			return
		}

                web.Success(ctx, http.StatusOK, buyerUpdated)
        }
}

func (b *Buyer) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

                id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Error(ctx, http.StatusBadRequest, err.Error())
			return
		}

		if err := b.buyerService.Delete(ctx, id); err != nil {
			web.Error(ctx, http.StatusNotFound, err.Error())
			return
		}

		web.Success(ctx, http.StatusNoContent, nil)
	}
}