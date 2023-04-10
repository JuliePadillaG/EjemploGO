package product

import (
	"context"
	"errors"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("product not found")
)

// Paso 1. Se debe generar la interface Service con todos sus métodos.
type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Save(ctx context.Context, description string, expiration_rate int, freezing_rate int, height float32, length float32, netweight float32, product_code string, recommended_freezing_temperature float32, width float32, product_type_id int, seller_id int) (domain.Product, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, id int, description string, expiration_rate *int, freezing_rate *int, height *float32, length *float32, netweight *float32, product_code string, recommended_freezing_temperature *float32, width *float32, product_type_id *int, seller_id *int) (domain.Product, error)
	GetProductRecords(ctx context.Context, id string) (product_record_report []domain.ProductRecordsReport, err error)
}

// Paso 2. Se debe generar la estructura service que contenga el repositorio.
type service struct {
	repository Repository
}

// Paso 3. Se debe generar una función que devuelva el Servicio.
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Paso 4. Se deben implementar todos los métodos correspondientes a las operaciones a realizar.
func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {

	products, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *service) Save(ctx context.Context, description string, expiration_rate int, freezing_rate int, height float32, length float32, netweight float32, product_code string, recommended_freezing_temperature float32, width float32, product_type_id int, seller_id int) (domain.Product, error) {

	if !s.repository.Exists(ctx, product_code) {
		newProduct := domain.Product{
			Description: description,
			ExpirationRate: expiration_rate,
			FreezingRate: freezing_rate,
			Height: height,
			Length: length,
			Netweight: netweight,
			ProductCode: product_code,
			RecomFreezTemp: recommended_freezing_temperature,
			Width: width,
			ProductTypeID: product_type_id,
			SellerID: seller_id,
		}

		id, err := s.repository.Save(ctx, newProduct)
		if err != nil {
			return domain.Product{}, err
		}
		newProduct.ID = id
		return newProduct, nil
	}

	return domain.Product{}, errors.New("product_code already exists")

}


func (s *service) Get(ctx context.Context, id int) (domain.Product, error) {
	pr, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Product{}, ErrNotFound
	}
	return pr, nil
}


func (s *service) Delete(ctx context.Context, id int) error {
	_, err := s.repository.Get(ctx, id)
	if err != nil {
		return ErrNotFound
	}

	return s.repository.Delete(ctx, id)
}


func (s *service) Update(ctx context.Context, id int, description string, expiration_rate *int, freezing_rate *int, height *float32, length *float32, netweight *float32, product_code string, recommended_freezing_temperature *float32, width *float32, product_type_id *int, seller_id *int) (domain.Product, error) {
	p, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Product{}, ErrNotFound
	}
	if description != "" {
		p.Description = description
	}
	if expiration_rate != nil {
		p.ExpirationRate = *expiration_rate
	}
	if freezing_rate != nil {
		p.FreezingRate = *freezing_rate
	}
	if height != nil {
		p.Height = *height
	}
	if length != nil {
		p.Length = *length
	}
	if netweight != nil {
		p.Netweight = *netweight
	}
	if product_code != "" {
		p.ProductCode = product_code
	}
	if recommended_freezing_temperature != nil {
		p.RecomFreezTemp = *recommended_freezing_temperature
	}
	if width != nil {
		p.Width = *width
	}
	if product_type_id != nil {
		p.ProductTypeID = *product_type_id
	}
	if seller_id != nil {
		p.SellerID = *seller_id
	}

	return p, s.repository.Update(ctx, p)
}

func (s *service) GetProductRecords(ctx context.Context, id string) (product_record_report []domain.ProductRecordsReport, err error) {
	
	product_record_report, err = s.repository.GetProductRecords(ctx, id)
	if err != nil {
		return nil, err
	}

	return product_record_report, nil
}
