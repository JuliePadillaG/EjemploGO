package employee

import (
	"context"
	"errors"
	"strconv"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("employee not found")
)

type Service interface {
	Save(ctx context.Context, cardNumberId string, name string, lastname string, wharehouseId int) (domain.Employee, error)
	GetAllEmployees(ctx context.Context) ([]domain.Employee, error)
	GetEmployeeByID(ctx context.Context, id int) (domain.Employee, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, name string, lastname string, wharehouseId *int) (domain.Employee, error)
	Report_BO(ctx context.Context, id string) ([]domain.ReportInBO, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}
func (s *service) Save(ctx context.Context, cardNumberId string, name string, lastname string, wharehouseId int) (domain.Employee, error) {
	if !s.repository.Exists(ctx, cardNumberId) {
		newemp := domain.Employee{
			CardNumberID: cardNumberId,
			FirstName:    name,
			LastName:     lastname,
			WarehouseID:  wharehouseId}
		id, err := s.repository.Save(ctx, newemp)
		if err != nil {
			return domain.Employee{}, err
		}
		newemp.ID = id
		return newemp, nil
	}
	return domain.Employee{}, errors.New("The card_number_id already exists")
}

func (s *service) GetAllEmployees(ctx context.Context) ([]domain.Employee, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) GetEmployeeByID(ctx context.Context, id int) (domain.Employee, error) {
	employeedfound, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Employee{}, ErrNotFound
	}
	return employeedfound, nil
}
func (s *service) Delete(ctx context.Context, id int) error {
	_, err := s.repository.Get(ctx, id)
	if err != nil {
		return ErrNotFound
	}
	return s.repository.Delete(ctx, id)
}

func (s *service) Update(ctx context.Context, id int, name string, lastname string, wharehouseId *int) (domain.Employee, error) {
	e, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Employee{}, ErrNotFound
	}
	if name != "" {
		e.FirstName = name
	}
	if lastname != "" {
		e.LastName = lastname
	}
	if wharehouseId != nil {
		e.WarehouseID = *wharehouseId
	}
	return e, s.repository.Update(ctx, e)
}

func (s *service) Report_BO(ctx context.Context, id string) ([]domain.ReportInBO, error) {
	if id != "" {
		id_e, err := strconv.Atoi(id)
		if err != nil {
			return []domain.ReportInBO{}, errors.New("invalid id format")
		}
		_, err = s.repository.Get(ctx, id_e)
		if err != nil {
			return []domain.ReportInBO{}, ErrNotFound
		}
		return s.repository.ReportInboundOrdersByID(ctx, id_e)
	}
	return s.repository.ReportInboundOrders(ctx)
}
