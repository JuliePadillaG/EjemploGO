package buyer

import (
	"context"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockRepository struct {
	DataMock       []domain.Buyer
	GetWasCalled   bool
	ExistWasCalled bool
	Error          string
}

func (m *MockRepository) GetAll(ctx context.Context) ([]domain.Buyer, error) {

	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}

	return m.DataMock, nil
}

func (m *MockRepository) Get(ctx context.Context, id int) (domain.Buyer, error) {

	m.GetWasCalled = true

	if m.Error != "" {
		return domain.Buyer{}, fmt.Errorf(m.Error)
	}

	var buyer domain.Buyer
	for i := range m.DataMock {
		if m.DataMock[i].ID == id {
			buyer = m.DataMock[i]
			break
		}
	}

	return buyer, nil
}

func (m *MockRepository) Exists(ctx context.Context, cardNumberID string) bool {

	exist := false
	m.ExistWasCalled = true

	for i := range m.DataMock {
		if m.DataMock[i].CardNumberID == cardNumberID {
			exist = true
		}
	}

	return exist
}

func (m *MockRepository) Save(ctx context.Context, b domain.Buyer) (int, error) {

	if m.Error != "" {
		return 0, fmt.Errorf(m.Error)
	}

	var lastID int

	if len(m.DataMock) != 0 {
		lastID = m.DataMock[len(m.DataMock)-1].ID
	}

	lastID++
	b.ID = lastID
	m.DataMock = append(m.DataMock, b)

	return lastID, nil
}

func (m *MockRepository) Delete(ctx context.Context, id int) error {

	if m.Error != "" {
		return fmt.Errorf(m.Error)
	}

	var pos int
	for i := range m.DataMock {
		if m.DataMock[i].ID == id {
			pos = i
			break
		}
	}

	// fmt.Printf("%+v before %p\n", m.DataMock, m.DataMock)

	m.DataMock = append(m.DataMock[:pos], m.DataMock[pos+1:]...)

	// fmt.Printf("%+v after %p\n", m.DataMock, m.DataMock)

	return nil
}

func (m *MockRepository) Update(ctx context.Context, b domain.Buyer) error {

	m.GetWasCalled = true

	if m.Error != "" {
		return fmt.Errorf(m.Error)
	}

	for i := range m.DataMock {
		if m.DataMock[i].ID == b.ID {
			m.DataMock[i] = b
			break
		}
	}

	return nil
}
