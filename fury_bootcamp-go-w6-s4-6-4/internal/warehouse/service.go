package warehouse

import (
	"context"
	"errors"
	"reflect"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("warehouse not found")
)

type Service interface {
	Get(id int) (domain.Warehouse, error)
	GetAll() ([]domain.Warehouse, error)
	Save(w domain.Warehouse) (int, error)
	Update(w domain.Warehouse, id int) (domain.Warehouse, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Get(id int) (domain.Warehouse, error) {
	return s.repository.Get(context.Background(), id)
}

func (s *service) GetAll() ([]domain.Warehouse, error) {
	return s.repository.GetAll(context.Background())
}

func (s *service) Save(w domain.Warehouse) (int, error) {
	if s.repository.Exists(context.Background(), w.WarehouseCode) {
		return 0, errors.New("warehouse code already exists")
	}

	return s.repository.Save(context.Background(), w)
}

func (s *service) Update(w domain.Warehouse, id int) (domain.Warehouse, error) {
	originalWarehouse, err := s.Get(id)
	if err != nil {
		return domain.Warehouse{}, err
	}
	values := reflect.ValueOf(w)
	for i := 0; i < values.NumField(); i++ {
		if values.Field(i).IsZero() || values.Type().Field(i).Name == "ID" {
			value := reflect.ValueOf(originalWarehouse).Field(i)
			reflect.ValueOf(&w).Elem().Field(i).Set(value)
		}
	}
	return w, s.repository.Update(context.Background(), w)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(context.Background(), id)
}
