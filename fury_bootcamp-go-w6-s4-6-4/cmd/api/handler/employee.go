package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/employee"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/pkg/web"
)

type request struct {
	CardNumberID string `json:"card_number_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	WarehouseID  *int   `json:"warehouse_id"`
}

type Employee struct {
	employeeService employee.Service
}

func NewEmployee(e employee.Service) *Employee {
	return &Employee{
		employeeService: e,
	}
}

func (e *Employee) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, 400, "%s", err)
			return
		}
		employeebyId, err := e.employeeService.GetEmployeeByID(c, int(id))
		if err != nil {
			web.Error(c, 404, "%s", err)
			return
		}
		web.Success(c, 200, employeebyId)
	}
}

func (e *Employee) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		employees, err := e.employeeService.GetAllEmployees(c)
		if err != nil {
			web.Error(c, 404, "%s", err)
			return
		}
		if len(employees) == 0 {
			web.Success(c, 200, "No existing employees")
			return
		}
		web.Success(c, 200, employees)
	}
}

func (e *Employee) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request

		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 422, "%s", err)
			return
		}
		if req.CardNumberID == "" {
			web.Error(c, 422, "%s", "card_number_id is required")
			return
		}
		if req.FirstName == "" {
			web.Error(c, 422, "%s", "first_name is required")
			return
		}
		if req.LastName == "" {
			web.Error(c, 422, "%s", "last_name is required")
			return
		}
		if req.WarehouseID == nil {
			defaultint := 0
			req.WarehouseID = &defaultint
		}
		if *req.WarehouseID < 0 {
			web.Error(c, 422, "%s", "WareHouseID cannot be negative")
			return
		}
		emp, err := e.employeeService.Save(c, req.CardNumberID, req.FirstName, req.LastName, *req.WarehouseID)
		if err != nil {
			if err.Error() == "The card_number_id already exists" {
				web.Error(c, 409, "%s", err)
			} else {
				web.Error(c, 500, "%s", err)
			}
			return
		}
		web.Success(c, 201, emp)
	}
}

func (e *Employee) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, 400, "%s", err)
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Error(c, 422, "%s", err)
			return
		}
		if req.CardNumberID != "" {
			web.Error(c, 400, "%s", "card_number_id field cannot be updated")
			return
		}
		employeeUpdate, err := e.employeeService.Update(c, int(id), req.FirstName, req.LastName, req.WarehouseID)
		if err != nil {
			web.Error(c, 404, "%s", err)
			return
		}
		web.Success(c, 200, employeeUpdate)
	}
}

func (e *Employee) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			web.Error(c, 400, "%s", err)
			return
		}
		err = e.employeeService.Delete(c, int(id))
		if err != nil {
			if err.Error() == "employee not found" {
				web.Error(c, 404, "%s", err)
			} else {
				web.Error(c, 500, "%s", err)
			}
			return
		}
		web.Success(c, 204, "employee deleted")
	}
}

func (e *Employee) Report_InboundOrders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")
		reports, err := e.employeeService.Report_BO(ctx, id)
		if err != nil {
			if err.Error() == "employee not found" {
				web.Error(ctx, 404, "%s", err)
			} else if err.Error() == "invalid id format" {
				web.Error(ctx, 422, "%s", err)
			} else {
				web.Error(ctx, 500, "%s", err)
			}
			return
		}
		if len(reports) == 0 {
			web.Success(ctx, 200, "No existing reports inbound_orders")
			return
		}
		web.Success(ctx, 200, reports)
	}
}
