package product_records

import (
	"context"
	"errors"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type Service interface {
	Save(ctx context.Context, last_update_date string, purchase_price float64, sale_price float64, products_id int) (domain.ProductRecords, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Save(ctx context.Context, last_update_date string, purchase_price float64, sale_price float64, products_id int) (domain.ProductRecords, error) {
	
	var newProductRecord domain.ProductRecords

	if s.repo.ExistsProductRecord(ctx, newProductRecord.ID) {
		return domain.ProductRecords{}, errors.New("error: product_records id already exists")
	} 
	
	if !s.repo.UniqueProduct(ctx, products_id) {
		return domain.ProductRecords{}, errors.New("error: product id doesn't exists")
	}
	
	newProductRecord = domain.ProductRecords{
		LastUpdateDate: last_update_date,
		PurchasePrice: purchase_price,
		SalePrice: sale_price,
		ProductID: products_id,
	}

	product_records_id, err := s.repo.Save(ctx, newProductRecord)
	if err != nil {
		return domain.ProductRecords{}, err
	}

	newProductRecord.ID = int(product_records_id)
	return newProductRecord, nil
}
