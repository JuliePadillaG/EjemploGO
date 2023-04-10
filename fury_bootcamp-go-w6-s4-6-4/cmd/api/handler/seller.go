package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/seller"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/pkg/web"
)

type requestSeller struct {
	ID          int    `json:"id"`
	CID         int    `json:"cid"`
	CompanyName string `json:"company_name"`
	Address     string `json:"address"`
	Telephone   string `json:"telephone"`
	LocalityID  int    `json:"locality_id"`
}

type Seller struct {
	sellerService seller.Service
}

func NewSeller(s seller.Service) *Seller {
	return &Seller{
		sellerService: s,
	}
}

func (s *Seller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		s, err := s.sellerService.GetAll()
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}
		web.Success(c, http.StatusOK, s)
	}
}

func (s *Seller) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusExpectationFailed, err.Error())
			return
		}
		s, err := s.sellerService.Get(id)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}
		web.Success(c, http.StatusOK, s)
	}
}

func (s *Seller) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestSeller
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}

		id, err := s.sellerService.Save(req.CID, req.LocalityID, req.CompanyName, req.Address, req.Telephone)
		if err != nil {
			if err.Error() == seller.ErrCidExists.Error() || err.Error() == seller.ErrLocalityNotFound.Error() {
				web.Error(c, http.StatusConflict, err.Error())
				return
			}
			if err.Error() == seller.ErrRequired.Error() {
				web.Error(c, http.StatusUnprocessableEntity, err.Error())
				return
			}
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		req.ID = id

		web.Success(c, http.StatusCreated, req)
	}
}

func (s *Seller) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		var req domain.Seller
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}

		req.ID = id

		req, err = s.sellerService.Update(req)
		if err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}
		web.Success(c, http.StatusOK, req)
	}
}

func (s *Seller) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		if err := s.sellerService.Delete(id); err != nil {
			web.Error(c, http.StatusNotFound, err.Error())
			return
		}
		web.Success(c, http.StatusOK, s)
	}
}
