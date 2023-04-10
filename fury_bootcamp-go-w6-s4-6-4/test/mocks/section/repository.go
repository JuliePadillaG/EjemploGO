package section

import (
	"context"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockRepository struct {
	DataMock []domain.Section
	Error    string
	ExistsID bool
	ID       int
}

func (m *MockRepository) GetAll(ctx context.Context) ([]domain.Section, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil
}

func (m *MockRepository) Get(ctx context.Context, id int) (domain.Section, error) {

	if m.Error != "" {
		return domain.Section{}, fmt.Errorf(m.Error)
	}
	for _, section := range m.DataMock {
		if section.ID == id {
			return section, nil
		}
	}
	return domain.Section{}, fmt.Errorf("section not found")
}

func (m *MockRepository) Exists(ctx context.Context, sectionNumber int) bool {
	for _, section := range m.DataMock {
		if section.SectionNumber == sectionNumber {
			return true
		}
	}
	return false
}

func (m *MockRepository) Save(ctx context.Context, s domain.Section) (int, error) {
	if m.Error != "" {
		return 0, fmt.Errorf(m.Error)
	}
	if m.Exists(context.Background(), s.SectionNumber) {
		return 0, fmt.Errorf(m.Error)
	}
	m.DataMock = append(m.DataMock, s)
	return s.ID, nil
}

func (m *MockRepository) Update(ctx context.Context, s domain.Section) error {
	if m.Error != "" {
		return fmt.Errorf(m.Error)
	}
	for i, section := range m.DataMock {
		if section.ID == s.ID {
			m.DataMock[i] = s
			return nil
		}
	}
	return fmt.Errorf("section not found")
}

func (m *MockRepository) Delete(ctx context.Context, id int) error {

	if m.Error != "" {
		return fmt.Errorf(m.Error)
	}

	for i, data := range m.DataMock {
		if data.ID == id {
			m.DataMock = append(m.DataMock[:i], m.DataMock[i+1:]...)
			return nil
		}
	}
	return nil
}
