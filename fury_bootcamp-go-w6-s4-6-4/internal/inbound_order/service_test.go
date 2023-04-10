package inboundorder

import (
	"context"
	"testing"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	inboundorder "github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/inbound_order"
	"github.com/stretchr/testify/assert"
)

var data = []domain.Inbound_order{
	{
		ID:               1,
		Order_date:       "2021-04-04",
		Order_number:     "order#1",
		Employee_id:      4,
		Product_batch_id: 1,
		Warehouse_id:     1,
	},
	{
		ID:               2,
		Order_date:       "2022-01-09",
		Order_number:     "order#1",
		Employee_id:      4,
		Product_batch_id: 1,
		Warehouse_id:     1,
	},
}

func Test_GetAll_Ok_Service(t *testing.T) {
	//Arrange
	var ibos []domain.Inbound_order
	ibos = append(ibos, data...)

	myMockR := inboundorder.MockRepositoryIBO{DataMock: ibos}
	service := NewService(&myMockR)

	//Act
	results, err := service.GetAll_inboundOrders(context.TODO())

	//Assert
	assert.True(t, myMockR.MethodCalled)
	assert.NoError(t, err)
	assert.Equal(t, ibos, results)
}

func Test_GetAll_Fail_Service(t *testing.T) {
	//Arrange
	var ibos []domain.Inbound_order
	ibos = append(ibos, data...)
	errorExpected := "Error inGetAll_inboundOrders"

	myMockR := inboundorder.MockRepositoryIBO{DataMock: ibos, Err: errorExpected}
	service := NewService(&myMockR)

	//Act
	results, err := service.GetAll_inboundOrders(context.TODO())

	//Assert
	assert.True(t, myMockR.MethodCalled)
	assert.NotNil(t, err)
	assert.Equal(t, errorExpected, err.Error())
	assert.Empty(t, results)

}

func Test_Save_ok_Service(t *testing.T) {
	//Arrange
	newIBO := domain.Inbound_order{
		ID:               1,
		Order_date:       "2021-04-04",
		Order_number:     "order#1",
		Employee_id:      1,
		Product_batch_id: 1,
		Warehouse_id:     1,
	}
	emp := []domain.Employee{
		{ID: 1,
			CardNumberID: "402323",
			FirstName:    "Jhon",
			LastName:     "Doe",
			WarehouseID:  1,
		},
	}

	myMockR := inboundorder.MockRepositoryIBO{DataMockEmp: emp}
	service := NewService(&myMockR)

	//Act
	result, err := service.Save(context.TODO(), newIBO.Order_date, newIBO.Order_number, newIBO.Employee_id, newIBO.Product_batch_id, newIBO.Warehouse_id)

	//Assert
	assert.True(t, myMockR.MethodCalled)
	assert.NoError(t, err)
	assert.Equal(t, newIBO, result)
}

func Test_Save_Fail_Service(t *testing.T) {
	newIBO := domain.Inbound_order{
		ID:               1,
		Order_date:       "2021-04-04",
		Order_number:     "order#1",
		Employee_id:      4,
		Product_batch_id: 1,
		Warehouse_id:     1,
	}

	t.Run("employee does not exist", func(t *testing.T) {
		myMockR := inboundorder.MockRepositoryIBO{}
		service := NewService(&myMockR)

		//Act
		result, err := service.Save(context.TODO(), newIBO.Order_date, newIBO.Order_number, newIBO.Employee_id, newIBO.Product_batch_id, newIBO.Warehouse_id)

		//Assert
		assert.True(t, myMockR.MethodCalled)
		assert.NotNil(t, err)
		assert.Equal(t, ErrEmployeeNotExist, err)
		assert.Empty(t, result)
	})
	t.Run("already exists", func(t *testing.T) {
		//Arrange
		myMockR := inboundorder.MockRepositoryIBO{DataMock: []domain.Inbound_order{newIBO}}
		service := NewService(&myMockR)

		//Act
		result, err := service.Save(context.TODO(), newIBO.Order_date, newIBO.Order_number, newIBO.Employee_id, newIBO.Product_batch_id, newIBO.Warehouse_id)

		//Assert
		assert.True(t, myMockR.MethodCalled)
		assert.NotNil(t, err)
		assert.Equal(t, ErrAlreadyExist, err)
		assert.Empty(t, result)
	})
}
