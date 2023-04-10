package productbatches

import (
	"context"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockRepository struct {
	DataMockPB []domain.Product_batches
	DataMockRP []domain.ReportProduct
	Error      string
	ExistsID   bool
	ID         int
}

func (m *MockRepository) CreatePB(ctx context.Context, pb domain.Product_batches) (int, error) {
	if m.Error != "" {
		return 0, fmt.Errorf(m.Error)
	}

	m.DataMockPB = append(m.DataMockPB, pb)
	return pb.ID, nil
}

func (m *MockRepository) ReadPB(ctx context.Context, id int) (domain.ReportProduct, error) {
	if m.Error != "" {
		return domain.ReportProduct{}, fmt.Errorf(m.Error)
	}
	for _, pb := range m.DataMockRP {
		if pb.SectionNumber == id {
			return pb, nil
		}
	}
	return domain.ReportProduct{}, fmt.Errorf("product batch not found")
}

func (m *MockRepository) GetPB(ctx context.Context, id int) (domain.Product_batches, error) {
	if m.Error != "" {
		return domain.Product_batches{}, fmt.Errorf(m.Error)
	}
	for _, pb := range m.DataMockPB {
		if pb.ID == id {
			return pb, nil
		}
	}
	return domain.Product_batches{}, fmt.Errorf("product batch not found")
}

func (m *MockRepository) ExistenceSectionId(ctx context.Context, section_id int) bool {
	for _, pb := range m.DataMockPB {
		if pb.SectionId == section_id {
			return true
		}
	}
	return false
}

func (m *MockRepository) ExistenceProductId(ctx context.Context, product_id int) bool {
	for _, pb := range m.DataMockPB {
		if pb.ProductId == product_id {
			return true
		}
	}
	return false
}

func (m *MockRepository) ExistsProductBatches(ctx context.Context, batch_number int) bool {
	for _, pb := range m.DataMockPB {
		if pb.BatchNumber == batch_number {
			return true
		}
	}
	return false
}
