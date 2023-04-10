package employee

import (
	"context"
	"errors"
	"testing"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/employee"
	"github.com/stretchr/testify/assert"
)

var data = []domain.Employee{
	{
		ID:           1,
		CardNumberID: "402323",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  1,
	},
	{
		ID:           2,
		CardNumberID: "402322",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  3,
	},
}

func Test_create_ok_Service(t *testing.T) {
	//Arrange
	empExpec := domain.Employee{
		ID:           1,
		CardNumberID: "402323",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  1,
	}
	myMockR := employee.MockRepositoryEmployee{}
	service := NewService(&myMockR)
	var ctx context.Context
	//Act
	employResult, err := service.Save(ctx, empExpec.CardNumberID, empExpec.FirstName, empExpec.LastName, empExpec.WarehouseID)
	//Assert
	assert.True(t, myMockR.MethodCalled)
	assert.Nil(t, err)
	assert.Equal(t, empExpec, employResult)
}

func Test_create_conflict_Service(t *testing.T) {
	//Arrange
	newEmployee := domain.Employee{
		ID:           0,
		CardNumberID: "402323",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  1,
	}
	errExpected := errors.New("The card_number_id already exists")
	var dat []domain.Employee
	dat = append(dat, data...)
	myMockR := employee.MockRepositoryEmployee{DataMock: dat}
	service := NewService(&myMockR)
	var ctx context.Context
	//Act
	employResult, err := service.Save(ctx, newEmployee.CardNumberID, newEmployee.FirstName, newEmployee.LastName, newEmployee.WarehouseID)
	//Assert
	assert.False(t, myMockR.MethodCalled)
	assert.NotNil(t, err)
	assert.Equal(t, errExpected, err)
	assert.Empty(t, employResult)
}

func Test_find_all_Service(t *testing.T) {
	//Arrange
	var employeesExpected []domain.Employee
	employeesExpected = append(employeesExpected, data...)
	myMockR := employee.MockRepositoryEmployee{DataMock: data}
	service := NewService(&myMockR)
	var ctx context.Context
	//Act
	employeesResult, err := service.GetAllEmployees(ctx)
	//Assert
	assert.True(t, myMockR.MethodCalled)
	assert.Nil(t, err)
	assert.Equal(t, employeesExpected, employeesResult)

	//Fail
	myMockR.Err = "Error in repository"
	employeesResult, err = service.GetAllEmployees(ctx)
	assert.True(t, myMockR.MethodCalled)
	assert.NotNil(t, err)
	assert.Empty(t, employeesResult)
}

func Test_find_by_id_non_existent_Service(t *testing.T) {
	//Arrange
	var dat []domain.Employee
	dat = append(dat, data...)
	myMockR := employee.MockRepositoryEmployee{DataMock: dat}
	service := NewService(&myMockR)
	var ctx context.Context
	//Act
	employeeResult, err := service.GetEmployeeByID(ctx, 3)
	//Assert
	assert.True(t, myMockR.MethodCalled)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNotFound, err)
	assert.Empty(t, employeeResult)
}

func Test_find_by_id_existent_Service(t *testing.T) {
	//Arrange
	var dat []domain.Employee
	dat = append(dat, data...)
	employeeExpected := dat[0]

	myMockR := employee.MockRepositoryEmployee{DataMock: dat}
	service := NewService(&myMockR)
	var ctx context.Context
	//Act
	employeeResult, err := service.GetEmployeeByID(ctx, employeeExpected.ID)
	//Assert
	assert.True(t, myMockR.MethodCalled)
	assert.Nil(t, err)
	assert.Equal(t, employeeExpected, employeeResult)
}

func Test_update_existent_Service(t *testing.T) {
	//Arrange
	employeeExpected := domain.Employee{
		ID:           1,
		CardNumberID: "402323",
		FirstName:    "Michell",
		LastName:     "Doe Jhonson",
		WarehouseID:  1,
	}
	var dat []domain.Employee
	dat = append(dat, data...)
	myMockR := employee.MockRepositoryEmployee{DataMock: dat}
	service := NewService(&myMockR)
	var ctx context.Context
	//Act
	employeeResult, err := service.Update(ctx, employeeExpected.ID, employeeExpected.FirstName, employeeExpected.LastName, &employeeExpected.WarehouseID)
	//Assert
	assert.True(t, myMockR.MethodCalled)
	assert.Nil(t, err)
	assert.Equal(t, employeeExpected, employeeResult)
	assert.Equal(t, employeeExpected, myMockR.DataMock[0])
}

func Test_update_non_existent_Service(t *testing.T) {
	//Arrange
	employeeUpdate := domain.Employee{
		ID:           5,
		CardNumberID: "402323",
		FirstName:    "Michell",
		LastName:     "Doe Jhonson",
		WarehouseID:  1,
	}
	var dat []domain.Employee
	dat = append(dat, data...)
	myMockR := employee.MockRepositoryEmployee{DataMock: dat}
	service := NewService(&myMockR)
	var ctx context.Context
	//Act
	employeeResult, err := service.Update(ctx, employeeUpdate.ID, employeeUpdate.FirstName, employeeUpdate.LastName, &employeeUpdate.WarehouseID)
	//Assert
	assert.True(t, myMockR.MethodCalled)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNotFound, err)
	assert.Empty(t, employeeResult)
}

func Test_delete_non_existent_Service(t *testing.T) {
	//Arrange
	employeeDelete := domain.Employee{
		ID:           5,
		CardNumberID: "402323",
		FirstName:    "Michell",
		LastName:     "Doe Jhonson",
		WarehouseID:  1,
	}
	myMockR := employee.MockRepositoryEmployee{DataMock: data}
	service := NewService(&myMockR)
	var ctx context.Context
	//Act
	err := service.Delete(ctx, employeeDelete.ID)
	//Assert
	assert.True(t, myMockR.MethodCalled)
	assert.NotNil(t, err)
	assert.Equal(t, ErrNotFound, err)
}

func Test_delete_ok_Service(t *testing.T) {
	//Arrange
	employeeDelete := domain.Employee{
		ID:           1,
		CardNumberID: "402323",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  1,
	}
	var dat []domain.Employee
	dat = append(dat, data...)
	myMockR := employee.MockRepositoryEmployee{DataMock: dat}
	service := NewService(&myMockR)
	var ctx context.Context
	//Act
	err := service.Delete(ctx, employeeDelete.ID)

	resultEmployee, errget := service.GetEmployeeByID(ctx, employeeDelete.ID)
	//Assert
	assert.True(t, myMockR.MethodCalled)
	assert.Nil(t, err)
	assert.Equal(t, ErrNotFound, errget)
	assert.Empty(t, resultEmployee)
}

func Test_report_bo_all_ok_Service(t *testing.T) {
	//Arrange
	data := []domain.ReportInBO{
		{
			ID:                   1,
			CardNumberID:         "402323",
			FirstName:            "Jhon",
			LastName:             "Doe",
			WarehouseID:          1,
			Inbound_orders_count: 2,
		},
		{
			ID:                   2,
			CardNumberID:         "402324",
			FirstName:            "Jhon",
			LastName:             "Doe",
			WarehouseID:          1,
			Inbound_orders_count: 0,
		},
	}
	myMockR := employee.MockRepositoryEmployee{DatamockIBO: data}
	service := NewService(&myMockR)
	var ctx context.Context

	//Act
	result, err := service.Report_BO(ctx, "")

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, data, result)
}

func Test_report_bo_all_fail_Service(t *testing.T) {
	//Arrange
	data := []domain.ReportInBO{
		{
			ID:                   1,
			CardNumberID:         "402323",
			FirstName:            "Jhon",
			LastName:             "Doe",
			WarehouseID:          1,
			Inbound_orders_count: 2,
		},
		{
			ID:                   2,
			CardNumberID:         "402324",
			FirstName:            "Jhon",
			LastName:             "Doe",
			WarehouseID:          1,
			Inbound_orders_count: 0,
		},
	}
	errorExpected := "Error report all"
	myMockR := employee.MockRepositoryEmployee{DatamockIBO: data, Err: errorExpected}
	service := NewService(&myMockR)
	var ctx context.Context

	//Act
	result, err := service.Report_BO(ctx, "")

	//Assert
	assert.NotNil(t, err)
	assert.Empty(t, result)
}
func Test_report_bo_byId_ok_Service(t *testing.T) {
	//Arrange
	data := []domain.ReportInBO{
		{
			ID:                   1,
			CardNumberID:         "402323",
			FirstName:            "Jhon",
			LastName:             "Doe",
			WarehouseID:          1,
			Inbound_orders_count: 2,
		},
	}
	datae := []domain.Employee{
		{
			ID:           1,
			CardNumberID: "402323",
			FirstName:    "Jhon",
			LastName:     "Doe",
			WarehouseID:  1,
		},
	}

	myMockR := employee.MockRepositoryEmployee{DatamockIBO: data, DataMock: datae}
	service := NewService(&myMockR)
	var ctx context.Context

	//Act
	result, err := service.Report_BO(ctx, "1")

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, data, result)
}

func Test_report_bo_byId_fail_Service(t *testing.T) {
	data := []domain.ReportInBO{
		{
			ID:                   1,
			CardNumberID:         "402323",
			FirstName:            "Jhon",
			LastName:             "Doe",
			WarehouseID:          1,
			Inbound_orders_count: 2,
		},
	}
	errorExpected := errors.New("employee not found")
	myMockR := employee.MockRepositoryEmployee{DatamockIBO: data}
	service := NewService(&myMockR)
	var ctx context.Context

	//Act
	result, err := service.Report_BO(ctx, "1")

	//Assert
	assert.NotNil(t, err)
	assert.Equal(t, errorExpected, err)
	assert.Empty(t, result)
}
