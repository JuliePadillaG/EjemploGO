package section

import (
	"context"
	"errors"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Errors
var (
	ErrNotFound = errors.New("section not found")
	ErrExists   = errors.New("section already exists")
)

// Paso 1. Se debe generar la interface Service con todos sus métodos.
type Service interface {
	GetAll(ctx context.Context) ([]domain.Section, error)
	Get(ctx context.Context, id int) (domain.Section, error)
	Exists(ctx context.Context, sectionNumber int) bool
	Save(ctx context.Context, s domain.Section) (int, error)
	Delete(ctx context.Context, id int) error
	Update(ctx context.Context, ID int, SectionNumber int, CurrentTemperature int, MinimumTemperature int, CurrentCapacity int, MinimumCapacity int, MaximumCapacity int, WarehouseID int, ProductTypeID int) (domain.Section, error)
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
func (s *service) GetAll(ctx context.Context) ([]domain.Section, error) {

	sections, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return sections, nil
}

func (s *service) Get(ctx context.Context, id int) (domain.Section, error) {

	sections, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Section{}, err
	}
	return sections, nil

}

func (s *service) Exists(ctx context.Context, sectionNumber int) bool {
	return s.repository.Exists(ctx, sectionNumber)
}

func (r *service) Save(ctx context.Context, s domain.Section) (int, error) {

	exists := r.Exists(ctx, s.SectionNumber)

	if exists {
		return 0, ErrExists //"section already exists"
	}
	return r.repository.Save(ctx, s)

}

func (r *service) Delete(ctx context.Context, id int) error {
	sections, err := r.repository.Get(ctx, id)
	if err != nil {
		return err
	}
	if sections.ID == 0 {
		return ErrNotFound
	}
	return r.repository.Delete(ctx, id)
}

func (r *service) Update(ctx context.Context, ID int, SectionNumber int, CurrentTemperature int, MinimumTemperature int, CurrentCapacity int, MinimumCapacity int, MaximumCapacity int, WarehouseID int, ProductTypeID int) (domain.Section, error) {

	sect, err := r.repository.Get(ctx, ID)
	if err != nil {
		return domain.Section{}, err
	}
	if SectionNumber != 0 {
		sect.SectionNumber = SectionNumber
	}
	if CurrentTemperature != 0 {
		sect.CurrentTemperature = CurrentTemperature
	}
	if MinimumTemperature != 0 {
		sect.MinimumTemperature = MinimumTemperature
	}
	if CurrentCapacity != 0 {
		sect.CurrentCapacity = CurrentCapacity
	}
	if MinimumCapacity != 0 {
		sect.MinimumCapacity = MinimumCapacity
	}
	if MaximumCapacity != 0 {
		sect.MaximumCapacity = MaximumCapacity
	}
	if WarehouseID != 0 {
		sect.WarehouseID = WarehouseID
	}
	if ProductTypeID != 0 {
		sect.ProductTypeID = ProductTypeID
	}
	return sect, r.repository.Update(ctx, sect)

}
