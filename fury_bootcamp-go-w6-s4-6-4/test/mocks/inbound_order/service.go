package inboundorder

import (
	"context"
	"errors"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockServiceIBO struct {
	DataMock     []domain.Inbound_order
	Err          string
	MethodCalled bool
}

func (ms *MockServiceIBO) GetAll_inboundOrders(ctx context.Context) ([]domain.Inbound_order, error) {
	ms.MethodCalled = true
	if ms.Err != "" {
		return []domain.Inbound_order{}, errors.New(ms.Err)
	}
	return ms.DataMock, nil
}

func (ms *MockServiceIBO) Save(ctx context.Context, order_date string, order_number string, employee_id int, product_batch_id int, warehouse_id int) (domain.Inbound_order, error) {
	ms.MethodCalled = true
	if ms.Err != "" {
		return domain.Inbound_order{}, errors.New(ms.Err)
	}
	var newIBO domain.Inbound_order
	var id int = 1
	if len(ms.DataMock) > 0 {
		id = ms.DataMock[len(ms.DataMock)-1].ID + 1
	}
	newIBO.ID = id
	newIBO.Order_date = order_date
	newIBO.Order_number = order_number
	newIBO.Employee_id = employee_id
	newIBO.Product_batch_id = product_batch_id
	newIBO.Warehouse_id = warehouse_id

	ms.DataMock = append(ms.DataMock, newIBO)

	return newIBO, nil
}
