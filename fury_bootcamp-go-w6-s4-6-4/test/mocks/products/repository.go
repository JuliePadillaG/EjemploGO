package products

import (
	"context"
	"errors"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockRepositoryProduct struct {
	DataMock []domain.Product
	Error    string
}

func (m *MockRepositoryProduct) GetAll(ctx context.Context) ([]domain.Product, error) {

	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil
}

func (m *MockRepositoryProduct) Get(ctx context.Context, id int) (domain.Product, error) {

	if m.Error != "" {
		return domain.Product{}, fmt.Errorf(m.Error)
	}

	for i := 0; i < 2; i++ {

		if m.DataMock[i].ID == id {
			p := domain.Product{
				ID:             m.DataMock[i].ID,
				Description:    m.DataMock[i].Description,
				ExpirationRate: m.DataMock[i].ExpirationRate,
				FreezingRate:   m.DataMock[i].FreezingRate,
				Height:         m.DataMock[i].Height,
				Length:         m.DataMock[i].Length,
				Netweight:      m.DataMock[i].Netweight,
				ProductCode:    m.DataMock[i].ProductCode,
				RecomFreezTemp: m.DataMock[i].RecomFreezTemp,
				Width:          m.DataMock[i].Width,
				ProductTypeID:  m.DataMock[i].ProductTypeID,
				SellerID:       m.DataMock[i].SellerID,
			}
			return p, nil
		}

		return domain.Product{}, fmt.Errorf(m.Error)
	}
	return domain.Product{}, fmt.Errorf(m.Error)
}

func (m *MockRepositoryProduct) Exists(ctx context.Context, productCode string) bool {

	for i, _ := range m.DataMock {
		if productCode == m.DataMock[i].ProductCode {
			return true
		}
	}

	return false

}

func (m *MockRepositoryProduct) Save(ctx context.Context, p domain.Product) (int, error) {
	if m.Error != "" {
		return 0, fmt.Errorf(m.Error)
	}

	id := 1

	if len(m.DataMock) > 0 {
		id = (m.DataMock[len(m.DataMock)-1].ID)
	}

	p.ID = id

	m.DataMock = append(m.DataMock, p)

	return id, nil
}

func (m *MockRepositoryProduct) Update(ctx context.Context, p domain.Product) error {
	for i := range m.DataMock {
		if m.DataMock[i].ID == p.ID {
			m.DataMock[i].Description = p.Description
			m.DataMock[i].ExpirationRate = p.ExpirationRate
			m.DataMock[i].FreezingRate = p.FreezingRate
			m.DataMock[i].Height = p.Height
			m.DataMock[i].Length = p.Length
			m.DataMock[i].Netweight = p.Netweight
			m.DataMock[i].ProductCode = p.ProductCode
			m.DataMock[i].RecomFreezTemp = p.RecomFreezTemp
			m.DataMock[i].Width = p.Width
			m.DataMock[i].ProductTypeID = p.ProductTypeID
			m.DataMock[i].SellerID = p.SellerID
			return nil
		}
	}
	return errors.New("not updated")
}

func (m *MockRepositoryProduct) Delete(ctx context.Context, id int) error {

	if m.Error != "" {
		return fmt.Errorf(m.Error)
	}

	for i, _ := range m.DataMock {

		if m.DataMock[i].ID == id {
			m.DataMock = append(m.DataMock[:i], m.DataMock[i+1:]...)
			return nil
		}
	}
	return errors.New("not eliminated")

}

func (m *MockRepositoryProduct) GetProductRecords(ctx context.Context, id string) (product_records_report []domain.ProductRecordsReport, err error) {

	return []domain.ProductRecordsReport{}, nil

}
