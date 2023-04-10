package buyer

import (
	"context"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockService struct {
	DataMock                []domain.Buyer
	GetAllBySellerWasCalled bool
	Error                   string
}

func (m *MockService) GetAll(ctx context.Context) ([]domain.Buyer, error) {

	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}

	return m.DataMock, nil
}

func (m *MockService) Get(ctx context.Context, id int) (domain.Buyer, error) {

	if m.Error != "" {
		return domain.Buyer{}, fmt.Errorf(m.Error)
	}

	var buyer domain.Buyer
	for i := range m.DataMock {
		if m.DataMock[i].ID == id {
			buyer = m.DataMock[i]
		}
	}

	return buyer, nil
}

func (m *MockService) Save(ctx context.Context, cardNumberID, firstName, lastName string) (domain.Buyer, error) {

	if m.Error != "" {
		return domain.Buyer{}, fmt.Errorf(m.Error)
	}

	var lastID int
	if len(m.DataMock) == 0 {
		lastID = 0
	} else {
		lastID = m.DataMock[len(m.DataMock)-1].ID
	}

	lastID++
	buyer := domain.Buyer{
		ID:           lastID,
		CardNumberID: cardNumberID,
		FirstName:    firstName,
		LastName:     lastName,
	}

	m.DataMock = append(m.DataMock, buyer)

	return buyer, nil
}

func (m *MockService) Delete(ctx context.Context, id int) error {

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

	fmt.Printf("%+v before %p\n", m.DataMock, m.DataMock)

	m.DataMock = append(m.DataMock[:pos], m.DataMock[pos+1:]...)

	fmt.Printf("%+v after %p\n", m.DataMock, m.DataMock)

	return nil
}

func (m *MockService) Update(ctx context.Context, id int, firstName, lastName string) (domain.Buyer, error) {

	if m.Error != "" {
		return domain.Buyer{}, fmt.Errorf(m.Error)
	}

	// no patch method
	buyer := domain.Buyer{
		FirstName: firstName,
		LastName:  lastName,
	}

	for i := range m.DataMock {
		if m.DataMock[i].ID == id {
			buyer.ID = m.DataMock[i].ID
			buyer.CardNumberID = m.DataMock[i].CardNumberID
			m.DataMock[i] = buyer
		}
	}

	return buyer, nil
}
