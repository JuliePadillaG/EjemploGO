package sellers

import (
	"errors"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("seller not found")
	ErrRequired = errors.New("field required")
	ErrExists   = errors.New("cid already exists")
	ErrRequest  = errors.New("incorrec field content")
)

type MockService struct {
	DataMock []domain.Seller
	Error    string
}

func (m *MockService) GetAll() ([]domain.Seller, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil
}

func (m *MockService) Get(id int) (domain.Seller, error) {
	for i, elemento := range m.DataMock {
		if elemento.ID == id {
			return m.DataMock[i], nil
		}
	}

	return domain.Seller{}, fmt.Errorf(m.Error)
}

func (m *MockService) Exists(cid int) bool {
	for _, elemento := range m.DataMock {
		if elemento.CID == cid {
			return true
		}
	}
	return false
}

func (m *MockService) Save(cid, locality int, companyName, address, telephone string) (int, error) {
	var seller domain.Seller

	if m.Exists(cid) {
		return 0, ErrExists
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
	if seller.LocalityID <= 0 {
		return 0, ErrRequired
	}

	if seller.CID < 0 {
		return 0, ErrRequest
	}

	m.DataMock = append(m.DataMock, seller)

	return 1, nil
}

func (m *MockService) Update(new domain.Seller) (domain.Seller, error) {
	anterior, err := m.Get(new.ID)
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
	if new.LocalityID <= 0 {
		new.LocalityID = anterior.LocalityID
	}
	if new.CID <= 0 {
		new.CID = anterior.CID
	}
	return new, nil
}

func (m *MockService) Delete(id int) error {
	_, err := m.Get(id)
	if err != nil {
		return ErrNotFound
	}
	return nil
}
