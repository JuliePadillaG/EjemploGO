package products

import (
	"context"
	"errors"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockServiceProduct struct {
	DataMock []domain.Product
	Error    string
}

func (m *MockServiceProduct) GetAll(ctx context.Context) ([]domain.Product, error) {
	if m.Error != "" {
		return []domain.Product{}, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil
}

func (m *MockServiceProduct) Get(ctx context.Context, id int) (domain.Product, error) {
	for _, value := range m.DataMock {
		if value.ID == id {
			return value, nil
		}
	}

	return domain.Product{}, errors.New("product not found")
}

func (m *MockServiceProduct) Save(ctx context.Context, description string, expiration_rate int, freezing_rate int, height float32, length float32, netweight float32, product_code string, recommended_freezing_temperature float32, width float32, product_type_id int, seller_id int) (domain.Product, error) {
	if m.Error != "" {
		return domain.Product{}, errors.New(m.Error)
	}

	for _, value := range m.DataMock {
		if value.ProductCode == product_code {
			return domain.Product{}, errors.New("product_code already exists")
		}
	}

	id := 1

	if len(m.DataMock) > 0 {
		id = m.DataMock[len(m.DataMock)-1].ID + 1
	}

	var newProduct domain.Product

	newProduct.ID = id
	newProduct.Description = description
	newProduct.ExpirationRate = expiration_rate
	newProduct.FreezingRate = freezing_rate
	newProduct.Height = height
	newProduct.Length = length
	newProduct.Netweight = netweight
	newProduct.ProductCode = product_code
	newProduct.RecomFreezTemp = recommended_freezing_temperature
	newProduct.Width = width
	newProduct.ProductTypeID = product_type_id
	newProduct.SellerID = seller_id

	m.DataMock = append(m.DataMock, newProduct)

	return newProduct, nil
}

func (m *MockServiceProduct) Delete(ctx context.Context, id int) error {
	idx := 0
	found := false

	for i := range m.DataMock {
		if m.DataMock[i].ID == id {
			idx = i
			found = true
		}
	}

	if !found {
		return errors.New("product not found")
	}

	m.DataMock = append(m.DataMock[:idx], m.DataMock[idx+1:]...)
	return nil
}

func (m *MockServiceProduct) Update(ctx context.Context, id int, description string, expiration_rate *int, freezing_rate *int, height *float32, length *float32, netweight *float32, product_code string, recommended_freezing_temperature *float32, width *float32, product_type_id *int, seller_id *int) (domain.Product, error) {
	idx := 0
	found := false

	for i := range m.DataMock {
		if m.DataMock[i].ID == id {
			idx = i
			found = true
		}
	}

	if !found {
		return domain.Product{}, errors.New("product not found")
	}
	if description != "" {
		m.DataMock[idx].Description = description
	}
	if expiration_rate != nil {
		m.DataMock[idx].ExpirationRate = *expiration_rate
	}
	if freezing_rate != nil {
		m.DataMock[idx].FreezingRate = *freezing_rate
	}
	if height != nil {
		m.DataMock[idx].Height = *height
	}
	if length != nil {
		m.DataMock[idx].Length = *length
	}
	if netweight != nil {
		m.DataMock[idx].Netweight = *netweight
	}
	if product_code != "" {
		m.DataMock[idx].ProductCode = product_code
	}
	if recommended_freezing_temperature != nil {
		m.DataMock[idx].RecomFreezTemp = *recommended_freezing_temperature
	}
	if width != nil {
		m.DataMock[idx].Width = *width
	}
	if product_type_id != nil {
		m.DataMock[idx].ProductTypeID = *product_type_id
	}
	if seller_id != nil {
		m.DataMock[idx].SellerID = *seller_id
	}

	return m.DataMock[idx], nil

}

func (m *MockServiceProduct) GetProductRecords(ctx context.Context, id string) (product_records_report []domain.ProductRecordsReport, err error) {

	return []domain.ProductRecordsReport{}, nil

}
