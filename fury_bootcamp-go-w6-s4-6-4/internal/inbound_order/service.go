package inboundorder

import (
	"context"
	"errors"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

var (
	ErrNotFound         = errors.New("Inbound_order not found")
	ErrAlreadyExist     = errors.New("The order_number already exists")
	ErrEmployeeNotExist = errors.New("The employee not exists")
)

type Service interface {
	GetAll_inboundOrders(ctx context.Context) ([]domain.Inbound_order, error)
	Save(ctx context.Context, order_date string, order_number string, employee_id int, product_batch_id int, wharehouse_id int) (domain.Inbound_order, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll_inboundOrders(ctx context.Context) ([]domain.Inbound_order, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Save(ctx context.Context, order_date string, order_number string, employee_id int, product_batch_id int, warehouse_id int) (domain.Inbound_order, error) {
	if !s.repository.ExistsInboundOrder(ctx, order_number) {
		if s.repository.ExistsEmployee(ctx, employee_id) {
			newInBoundOrder := domain.Inbound_order{
				Order_date:       order_date,
				Order_number:     order_number,
				Employee_id:      employee_id,
				Product_batch_id: product_batch_id,
				Warehouse_id:     warehouse_id,
			}
			id, err := s.repository.Save(ctx, newInBoundOrder)
			if err != nil {
				return domain.Inbound_order{}, err
			}
			newInBoundOrder.ID = id
			return newInBoundOrder, nil
		}
		return domain.Inbound_order{}, ErrEmployeeNotExist
	}
	return domain.Inbound_order{}, ErrAlreadyExist
}
