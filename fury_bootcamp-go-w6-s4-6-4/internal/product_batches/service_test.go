package productbatches

import (
	"context"
	//"fmt"
	"testing"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	productbatches "github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/product_batches"
	"github.com/stretchr/testify/assert"
)

func TestCreatePBService(t *testing.T) {
	
	t.Run("Create a new product batch", func(t *testing.T) {
		//Arrange
		//Preparar
		mockRepository := productbatches.MockRepository{}
		mockRepository.DataMockPB = []domain.Product_batches{
			{
				ID:             		1,
				BatchNumber:    		1,
				CurrentQuantity:  		1,
				CurrentTemperature:		1,
				DueDate:        		"2021-01-01",
				InitialQuantity:    	1,
				ManufacturingDate:  	"2021-01-01",
				ManufacturingHour:  	1,
				MinimumTemperature:     1,
				ProductId:      		1,
				SectionId:      		1,

			},
		}
		

		//mockRepository.ID = 1
		mockRepository.Error = ""
		//mockRepository.ExistsID = false

		//Act
		//Actuar
		service := NewService(&mockRepository)
		result, err := service.CreatePB(context.Background(), domain.Product_batches{
				
			ID:             		2,
			BatchNumber:    		2,
			CurrentQuantity:  		2,
			CurrentTemperature:		2,
			DueDate:        		"2021-01-01",
			InitialQuantity:    	2,
			ManufacturingDate:  	"2021-01-01",
			ManufacturingHour:  	2,
			MinimumTemperature:     2,
			ProductId:      		1,
			SectionId:      		1,
		})

		//Assert
		//Afirmar
		assert.NoError(t, err)
		assert.Equal(t, 2, result)
})
	t.Run("Fail to create a new product batch", func(t *testing.T) {
	//Arrange
	//Preparar
	mockRepository := productbatches.MockRepository{}
	mockRepository.DataMockPB = []domain.Product_batches{}
	mockRepository.Error = "product batch not found"
	//Act
	//Actuar
	service := NewService(&mockRepository)
	result, err := service.CreatePB(context.Background(), domain.Product_batches{
		ID:             		1,
		BatchNumber:    		1,
		CurrentQuantity:  		1,
		CurrentTemperature:		1,
		DueDate:        		"2021-01-01",
		InitialQuantity:    	1,
		ManufacturingDate:  	"2021-01-01",
		ManufacturingHour:  	1,
		MinimumTemperature:     1,
		ProductId:      		1,
		SectionId:      		1,
	})

	//Assert
	//Afirmar
	assert.Error(t, err)
	assert.Equal(t, 0, result)
})

}


func TestReadPBService(t *testing.T){
	
	t.Run("Fail to read a product batch", func(t *testing.T) {
	//Arrange
	//Preparar
	mockRepository := productbatches.MockRepository{}
	mockRepository.DataMockPB = []domain.Product_batches{}
	mockRepository.Error = "product batch not found"
	//Act
	//Actuar
	service := NewService(&mockRepository)
	_, err := service.ReadPB(context.Background(), 1)

	//Assert
	//Afirmar
	assert.Error(t, err)
})
}

