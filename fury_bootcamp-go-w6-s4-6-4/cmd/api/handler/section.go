package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/section"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/pkg/web"
)

type Section struct {
	sectionService section.Service
}

func NewSection(s section.Service) *Section {
	return &Section{
		sectionService: s,
	}
}

// GetAllSections godoc
// @Summary Get all sections
// @Tags Sections
// @Description get all sections
// @Accept  json
// @Produce  json
// @Success 200 {object} web.response
// @Router /api/v1/sections [get]
func (s *Section) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {

		sect, err := s.sectionService.GetAll(c)
		if err != nil {
			web.Error(c, http.StatusNotFound, "%s", err)
			return
		}

		web.Success(c, http.StatusOK, sect)
	}
}

// GetSection godoc
// @Summary Get section
// @Tags Sections
// @Description get one section
// @Accept  json
// @Produce  json
// @Param id path int true "Section ID"
// @Success 200 {object} web.response
// @Router /api/v1/sections/:id [get]
func (s *Section) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			web.Error(c, http.StatusNotFound, "%s", err)
			return
		}
		sect, err := s.sectionService.Get(c, idInt)
		if err != nil {
			web.Error(c, http.StatusConflict, "%s", err)
			return
		}
		web.Success(c, http.StatusOK, sect)
	}
}

// UpdateSection godoc
// @Summary Update section
// @Tags Sections
// @Description update section
// @Accept  json
// @Produce  json
// @Param id path int true "Section ID"
// @Success 200 {object} web.response
// @Router /api/v1/sections/:id [patch]
func (s Section) Update() gin.HandlerFunc {
	return func(c *gin.Context) {

		var req domain.Section

		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, http.StatusNotFound, "%s", err)
			return
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, http.StatusNotFound, "%s", err)
			return
		}

		sec, err := s.sectionService.Update(c, int(id), req.SectionNumber, req.CurrentTemperature, req.MinimumTemperature, req.CurrentCapacity, req.MinimumCapacity, req.MaximumCapacity, req.WarehouseID, req.ProductTypeID)
		if err != nil {
			web.Error(c, http.StatusNotFound, "%s", err)
			return
		}

		web.Success(c, http.StatusOK, sec)
	}
}

// DeleteSection godoc
// @Summary Delete section
// @Tags Sections
// @Description delete section
// @Accept  json
// @Produce  json
// @Param id path int true "Section ID"
// @Success 204
// @Router /api/v1/sections/:id [delete]
func (s *Section) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			web.Error(c, http.StatusConflict, "%s", err)
			return
		}
		err = s.sectionService.Delete(c, idInt)
		if err != nil {
			web.Error(c, http.StatusNotFound, "%s", err)
			return
		}

		web.Success(c, http.StatusNoContent, "Deleted")
	}
}

// CreateSection godoc
// @Summary Create section
// @Tags Sections
// @Description create section
// @Accept  json
// @Produce  json
// @Param section body domain.Section true "Section"
// @Success 201 {object} web.response
// @Router /api/v1/sections [post]
func (s *Section) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req domain.Section

		id, err := s.sectionService.Save(c, req)
		if err != nil && err.Error() == "section already exists" {
			web.Error(c, http.StatusConflict, "%s", err)
			return
		}
		req.ID = id
	
		web.Success(c, http.StatusCreated, req)
	}
}
