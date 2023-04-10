package productbatches

import (
	"context"
	"errors"
	"log"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Errors
var (
	ErrNotFoundSectionID = errors.New("section_id not found")
	ErrNotFoundProductID = errors.New("product_id not found")
	ErrExists            = errors.New("product_batches already exists")
)

type Service interface {
	CreatePB(ctx context.Context, pb domain.Product_batches) (int, error)
	ReadPB(ctx context.Context, id int) (domain.ReportProduct, error)
	ExistenceSectionId(ctx context.Context, section_id int) bool
	ExistenceProductId(ctx context.Context, product_id int) bool
	ExistsProductBatches(ctx context.Context, batch_number int) bool
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreatePB(ctx context.Context, pb domain.Product_batches) (int, error) {
	existsSectionId := s.repository.ExistenceSectionId(ctx, pb.SectionId)
	log.Println("Section_id", pb.SectionId)
	log.Println("Product_id", pb.ProductId)

	if !existsSectionId {
		return 0, ErrNotFoundSectionID
	}

	existsProductId := s.repository.ExistenceProductId(ctx, pb.ProductId)

	if !existsProductId {
		return 0, ErrNotFoundProductID
	}

	exists := s.repository.ExistsProductBatches(ctx, pb.BatchNumber)
	if exists {
		return 0, ErrExists
	}

	return s.repository.CreatePB(ctx, pb)
}

func (s *service) ReadPB(ctx context.Context, id int) (domain.ReportProduct, error) {
	return s.repository.ReadPB(ctx, id)
}

func (s *service) ExistenceSectionId(ctx context.Context, section_id int) bool {
	return s.repository.ExistenceSectionId(ctx, section_id)
}

func (s *service) ExistenceProductId(ctx context.Context, product_id int) bool {
	return s.repository.ExistenceProductId(ctx, product_id)
}

func (s *service) ExistsProductBatches(ctx context.Context, batch_number int) bool {
	return s.repository.ExistsProductBatches(ctx, batch_number)
}
