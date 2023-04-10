package buyer

import (
	"context"
	"errors"
	"reflect"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("buyer not found")
        ErrDuplicateCardNumberID = errors.New("duplicate cardNumberID")
)

type Service interface {
        Save(ctx context.Context, cardNumberID, firstName, lastName string) (domain.Buyer, error)
	GetAll(ctx context.Context) ([]domain.Buyer, error)
	Get(ctx context.Context, id int) (domain.Buyer, error)
        Delete(ctx context.Context, id int) error
        // Update(ctx context.Context, id int, cardNumberID, firstName, lastName string) error
        Update(ctx context.Context, id int, firstName, lastName string) (domain.Buyer, error) 
}

type service struct{
        repository Repository
}

func NewService(r Repository) Service {
	return &service{
                repository: r,
        }
}

func (s *service) GetAll(ctx context.Context) ([]domain.Buyer, error) {
        return s.repository.GetAll(ctx)
}

func (s *service) Save(ctx context.Context, cardNumberID, firstName, lastName string) (domain.Buyer, error) {

        if s.repository.Exists(ctx, cardNumberID) {
		return domain.Buyer{}, ErrDuplicateCardNumberID
        }

        buyer := domain.Buyer{}
        buyer.CardNumberID = cardNumberID
        buyer.FirstName = firstName
        buyer.LastName = lastName

        id, err := s.repository.Save(ctx, buyer)
	if err != nil {
		return domain.Buyer{}, err
	}

        buyer.ID = id

       return buyer, nil
}

func (s *service) Get(ctx context.Context, id int) (domain.Buyer, error) {
        return s.repository.Get(ctx, id)
}

func (s *service) Delete(ctx context.Context, id int) error {

        _, err := s.repository.Get(ctx, id)
	if err != nil {
		return ErrNotFound
	}

	return s.repository.Delete(ctx, id)
}

func (s *service) Update(ctx context.Context, id int, firstName, lastName string) (domain.Buyer, error) {

        currentBuyer, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Buyer{}, ErrNotFound
	}

        req := domain.Buyer{}
        req.FirstName = firstName
        req.LastName = lastName

	values := reflect.ValueOf(req)

	for i := 0; i < values.NumField(); i++ {
		if !values.Field(i).IsZero() {
			value := reflect.ValueOf(req).Field(i)
			reflect.ValueOf(&currentBuyer).Elem().Field(i).Set(value)
		}
	}
        
        if err = s.repository.Update(ctx, currentBuyer); err != nil {
                return domain.Buyer{}, err
        }

        return currentBuyer, nil
}

