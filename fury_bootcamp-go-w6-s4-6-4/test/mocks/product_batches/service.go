package productbatches

import (
	"context"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockService struct {
	Db MockRepository
}

func (s *MockService) CreatePB(ctx context.Context, pb domain.Product_batches) (int, error) {
	if s.Db.ExistsID {
		return 0, fmt.Errorf(s.Db.Error)
	}
	//id, err := MockRepository. .Save(ctx, section)
	if s.Db.Error != "" {
		return 0, fmt.Errorf(s.Db.Error)
	}
	s.Db.ID++
	return s.Db.ID, nil
}

func (s *MockService) ReadPB(ctx context.Context, id int) (domain.ReportProduct, error) {
	if s.Db.Error != "" {
		return domain.ReportProduct{}, fmt.Errorf(s.Db.Error)
	}
	return s.Db.ReadPB(ctx, id)
}

func (s *MockService) GetPB(ctx context.Context, id int) (domain.Product_batches, error) {
	if s.Db.Error != "" {
		return domain.Product_batches{}, fmt.Errorf(s.Db.Error)
	}
	return s.Db.GetPB(ctx, id)
}

func (s *MockService) ExistenceSectionId(ctx context.Context, section_id int) bool {
	return s.Db.ExistenceSectionId(ctx, section_id)
}

func (s *MockService) ExistsProductBatches(ctx context.Context, batch_number int) bool {
	return s.Db.ExistsProductBatches(ctx, batch_number)
}

func (s *MockService) ExistenceProductId(ctx context.Context, product_id int) bool {
	return s.Db.ExistenceProductId(ctx, product_id)
}
