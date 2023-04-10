package section

import (
	"context"
	"errors"
	"fmt"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type MockService struct {
	Db MockRepository
}

func (s *MockService) Delete(ctx context.Context, id int) error {

	for i, data := range s.Db.DataMock {
		if data.ID == id {
			s.Db.DataMock = append(s.Db.DataMock[:i], s.Db.DataMock[i+1:]...)
			return nil
		}
	}
	return errors.New("section not found")
}
func (s *MockService) GetAll(ctx context.Context) ([]domain.Section, error) {

	if s.Db.Error != "" {
		return nil, fmt.Errorf(s.Db.Error)
	}
	return s.Db.DataMock, nil
}

func (s *MockService) Get(ctx context.Context, id int) (domain.Section, error) {

	if s.Db.Error != "" {
		return domain.Section{}, fmt.Errorf(s.Db.Error)
	}
	for _, section := range s.Db.DataMock {
		if section.ID == id {
			return section, nil
		}
	}
	return domain.Section{}, fmt.Errorf(s.Db.Error)
}

func (s *MockService) Exists(ctx context.Context, sectionNumber int) bool {
	return s.Db.ExistsID
}

func (s *MockService) Save(ctx context.Context, section domain.Section) (int, error) {
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

func (s *MockService) Update(ctx context.Context, ID int, SectionNumber int, CurrentTemperature int, MinimumTemperature int, CurrentCapacity int, MinimumCapacity int, MaximumCapacity int, WarehouseID int, ProductTypeID int) (domain.Section, error) {
	value := 0
	flag := false
	for i := range s.Db.DataMock {
		if s.Db.DataMock[i].ID == ID {
			value = i
			flag = true
		}
	}
	if !flag {
		return domain.Section{}, errors.New("section not found")
	}

	var copyDatamock = s.Db.DataMock[value]

	if CurrentTemperature != 0 {
		copyDatamock.CurrentTemperature = CurrentTemperature
	}
	if MinimumTemperature != 0 {
		copyDatamock.MinimumTemperature = MinimumTemperature
	}
	if CurrentCapacity != 0 {
		copyDatamock.CurrentCapacity = CurrentCapacity
	}
	if MinimumCapacity != 0 {
		copyDatamock.MinimumCapacity = MinimumCapacity
	}
	if MaximumCapacity != 0 {
		copyDatamock.MaximumCapacity = MaximumCapacity
	}
	if WarehouseID != 0 {
		copyDatamock.WarehouseID = WarehouseID
	}
	if ProductTypeID != 0 {
		copyDatamock.ProductTypeID = ProductTypeID
	}

	s.Db.DataMock[value] = copyDatamock

	return copyDatamock, nil
}
