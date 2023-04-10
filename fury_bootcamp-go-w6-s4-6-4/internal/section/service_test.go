package section

import (
	"context"
	"fmt"

	"testing"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/section"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {

	t.Run("Get all sections", func(t *testing.T) {
		//Arrange
		//Preparar
		mockRepository := section.MockRepository{}
		mockRepository.DataMock = []domain.Section{
			{
				ID:                 1,
				SectionNumber:      1,
				CurrentTemperature: 1,
				MinimumTemperature: 1,
				CurrentCapacity:    1,
				MinimumCapacity:    1,
				MaximumCapacity:    1,
				WarehouseID:        1,
				ProductTypeID:      1,
			},
			{
				ID:                 2,
				SectionNumber:      2,
				CurrentTemperature: 2,
				MinimumTemperature: 2,
				CurrentCapacity:    2,
				MinimumCapacity:    2,
				MaximumCapacity:    2,
				WarehouseID:        2,
				ProductTypeID:      2,
			},
		}

		//Act
		//Actuar
		service := NewService(&mockRepository)
		result, err := service.GetAll(context.Background())

		//Assert
		//Afirmar
		assert.NoError(t, err)
		assert.Equal(t, 2, len(result))
	})
	t.Run("Fail to get all sections", func(t *testing.T) {
		//Arrange
		//Preparar
		mockRepository := section.MockRepository{Error: "section not found"}
		mockRepository.DataMock = []domain.Section{}
		mockRepository.Error = "section not found"

		//Act
		//Actuar
		service := NewService(&mockRepository)
		_, err := service.GetAll(context.Background())

		//Assert
		//Afirmar
		assert.Error(t, err)

	})
}

func TestFind_by_id_existent(t *testing.T) {
	//Arrange
	//Preparar
	mockRepository := section.MockRepository{}
	mockRepository.DataMock = []domain.Section{
		{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
		{
			ID:                 2,
			SectionNumber:      2,
			CurrentTemperature: 2,
			MinimumTemperature: 2,
			CurrentCapacity:    2,
			MinimumCapacity:    2,
			MaximumCapacity:    2,
			WarehouseID:        2,
			ProductTypeID:      2,
		},
	}

	//Act
	//Actuar

	service := NewService(&mockRepository)
	result, err := service.Get(context.Background(), 2)

	//Assert
	//Afirmar

	assert.NoError(t, err)
	assert.Equal(t, mockRepository.DataMock[1], result)
	fmt.Println(result)

}

func TestFind_by_id_not_existent(t *testing.T) {
	//Arrange
	//Preparar
	mockRepository := section.MockRepository{}
	mockRepository.DataMock = []domain.Section{
		{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
		{
			ID:                 2,
			SectionNumber:      2,
			CurrentTemperature: 2,
			MinimumTemperature: 2,
			CurrentCapacity:    2,
			MinimumCapacity:    2,
			MaximumCapacity:    2,
			WarehouseID:        2,
			ProductTypeID:      2,
		},
	}

	//Act
	//Actuar

	service := NewService(&mockRepository)
	result, err := service.Get(context.Background(), 3)

	//Assert
	//Afirmar

	assert.Error(t, err)
	assert.Equal(t, domain.Section{}, result)

}

func TestDelete_ok(t *testing.T) {

	mocKRepository := section.MockRepository{}
	mocKRepository.DataMock = []domain.Section{
		{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
		{
			ID:                 2,
			SectionNumber:      2,
			CurrentTemperature: 2,
			MinimumTemperature: 2,
			CurrentCapacity:    2,
			MinimumCapacity:    2,
			MaximumCapacity:    2,
			WarehouseID:        2,
			ProductTypeID:      2,
		},
	}
	expectedResult := []domain.Section{
		{
			ID:                 2,
			SectionNumber:      2,
			CurrentTemperature: 2,
			MinimumTemperature: 2,
			CurrentCapacity:    2,
			MinimumCapacity:    2,
			MaximumCapacity:    2,
			WarehouseID:        2,
			ProductTypeID:      2,
		},
	}

	service := NewService(&mocKRepository)
	err := service.Delete(context.Background(), 1)

	//Assert
	//Afirmar

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, mocKRepository.DataMock)

}

func TestDelete_non_existent(t *testing.T) {
	//Iniciar el repo fake
	//arrange
	//preparar
	mocKRepository := section.MockRepository{}
	mocKRepository.DataMock = []domain.Section{
		{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	//Act
	//Actuar
	service := NewService(&mocKRepository)
	err := service.Delete(context.Background(), 2)

	//Assert
	//Afirmar

	assert.Error(t, err)

}

func TestCreate_ok(t *testing.T) {
	//Arrange
	//Preparar

	mockRepository := section.MockRepository{}
	mockRepository.DataMock = []domain.Section{
		{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	newSection := domain.Section{

		ID:                 2,
		SectionNumber:      2,
		CurrentTemperature: 2,
		MinimumTemperature: 2,
		CurrentCapacity:    2,
		MinimumCapacity:    2,
		MaximumCapacity:    2,
		WarehouseID:        2,
		ProductTypeID:      2,
	}

	//Act
	//Actuar

	service := NewService(&mockRepository)
	result, err := service.Save(context.Background(), newSection)

	//Assert
	//Afimar

	assert.NoError(t, err)
	assert.Equal(t, newSection.ID, result)

}

func TestCreate_conflict(t *testing.T) {
	//Arrange
	//Preparar
	mockRepository := section.MockRepository{}
	mockRepository.DataMock = []domain.Section{
		{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	newSection := domain.Section{

		ID:                 1,
		SectionNumber:      1,
		CurrentTemperature: 1,
		MinimumTemperature: 1,
		CurrentCapacity:    1,
		MinimumCapacity:    1,
		MaximumCapacity:    1,
		WarehouseID:        1,
		ProductTypeID:      1,
	}

	//Act
	//Actuar
	service := NewService(&mockRepository)
	result, err := service.Save(context.Background(), newSection)

	assert.Error(t, err)
	assert.Equal(t, 0, result)
}

func TestUpdate_Existent(t *testing.T) {
	//Arrange
	//Preparar
	//Iniciamos el repo fake
	mockRepository := section.MockRepository{}

	mockRepository.DataMock = []domain.Section{
		{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	expectedUpdate := domain.Section{

		ID:                 1,
		SectionNumber:      2,
		CurrentTemperature: 2,
		MinimumTemperature: 2,
		CurrentCapacity:    2,
		MinimumCapacity:    2,
		MaximumCapacity:    2,
		WarehouseID:        2,
		ProductTypeID:      2,
	}

	//Act
	//Actuar
	// iniciamos el newService para poder usar los metodos de service
	service := NewService(&mockRepository)
	result, err := service.Update(context.Background(), expectedUpdate.ID, expectedUpdate.SectionNumber, expectedUpdate.CurrentTemperature, expectedUpdate.MinimumTemperature, expectedUpdate.CurrentCapacity, expectedUpdate.MinimumCapacity, expectedUpdate.MaximumCapacity, expectedUpdate.WarehouseID, expectedUpdate.ProductTypeID)

	//Assert
	//Afirmar

	assert.NoError(t, err)
	assert.Equal(t, expectedUpdate, mockRepository.DataMock[0])
	assert.Equal(t, expectedUpdate.ID, result.ID)
	t.Log(mockRepository.DataMock)

}

func TestUpdate_non_existent(t *testing.T) {
	//Arrange
	//Preparar
	//Iniciamos el repo fake
	mockRepository := section.MockRepository{}

	mockRepository.DataMock = []domain.Section{
		{
			ID:                 1,
			SectionNumber:      1,
			CurrentTemperature: 1,
			MinimumTemperature: 1,
			CurrentCapacity:    1,
			MinimumCapacity:    1,
			MaximumCapacity:    1,
			WarehouseID:        1,
			ProductTypeID:      1,
		},
	}

	//Act
	//Actuar
	// iniciamos el newService para poder usar los metodos de service
	service := NewService(&mockRepository)
	result, err := service.Update(context.Background(), mockRepository.DataMock[0].ID+1, 2, 2, 2, 2, 2, 2, 2, 2)

	//Assert
	//Afirmar

	assert.Error(t, err)
	assert.Empty(t, result)
	fmt.Println(domain.Section{})

}
