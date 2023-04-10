package seller

import (
	"context"
	"errors"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Errors
var (
	ErrNotFound         = errors.New("seller not found")
	ErrRequired         = errors.New("field required")
	ErrCidExists        = errors.New("cid already exists")
	ErrLocalityNotFound = errors.New("locality id not found")
	ErrRequest          = errors.New("incorrect field content")
)

type Service interface {
	GetAll() ([]domain.Seller, error)
	Get(id int) (domain.Seller, error)
	Exists(cid int) bool
	Save(cid, locality int, companyName, address, telephone string) (int, error)
	Update(new domain.Seller) (domain.Seller, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) GetAll() ([]domain.Seller, error) {
	return s.repository.GetAll(context.Background())
}

func (s *service) Get(id int) (domain.Seller, error) {
	return s.repository.Get(context.Background(), id)
}

func (s *service) Exists(cid int) bool {
	return s.repository.Exists(context.Background(), cid)
}

func (s *service) Save(cid, locality int, companyName, address, telephone string) (int, error) {
	var seller domain.Seller

	if s.repository.Exists(context.Background(), cid) {
		return 0, ErrCidExists
	}

	if !s.repository.LocalityExists(context.Background(), locality) {
		return 0, ErrLocalityNotFound
	}

	seller.CID = cid
	seller.Address = address
	seller.CompanyName = companyName
	seller.Telephone = telephone
	seller.LocalityID = locality

	if seller.Address == "" {
		return 0, ErrRequired
	}
	if seller.CompanyName == "" {
		return 0, ErrRequired
	}
	if seller.Telephone == "" {
		return 0, ErrRequired
	}
	if seller.CID == 0 {
		return 0, ErrRequired
	}

	if seller.CID < 0 {
		return 0, ErrRequest
	}

	return s.repository.Save(context.Background(), seller)
}

func (s *service) Update(new domain.Seller) (domain.Seller, error) {
	anterior, err := s.repository.Get(context.Background(), new.ID)
	if err != nil {
		return domain.Seller{}, ErrNotFound
	}

	if new.Address == "" {
		new.Address = anterior.Address
	}
	if new.CompanyName == "" {
		new.CompanyName = anterior.CompanyName
	}
	if new.Telephone == "" {
		new.Telephone = anterior.Telephone
	}
	if new.CID <= 0 {
		new.CID = anterior.CID
	}
	if new.LocalityID <= 0 {
		new.LocalityID = anterior.LocalityID
	}

	return new, s.repository.Update(context.Background(), new)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(context.Background(), id)
}
