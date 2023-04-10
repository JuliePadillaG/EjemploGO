package locality

import (
	"context"
	"testing"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/locality"
	"github.com/stretchr/testify/assert"
)

func TestCreateLocalityS(t *testing.T) {
	//Arrange
	database := []domain.Locality{{
		ID:           1754,
		LocalityName: "San Justo",
		ProvinceName: "Buenos Aires",
		CountryName:  "Argentina",
	}}

	locality_test := domain.Locality{
		ID:           1759,
		LocalityName: "Gonzalez Catan",
		ProvinceName: "Buenos Aires",
		CountryName:  "Argentina",
	}

	mockRepository := locality.MockRepository{
		DataMock: database,
		Error:    "",
	}

	service := NewService(&mockRepository)

	t.Run("create ok", func(t *testing.T) {
		//Act
		results, err := service.Create(context.TODO(), locality_test)

		//Assert
		assert.Nil(t, err)
		assert.Equal(t, locality_test, results)
	})

	t.Run("create id exist", func(t *testing.T) {
		locality_test.ID = 1754

		//Act
		results, err := service.Create(context.TODO(), locality_test)

		//Assert
		assert.NotNil(t, err)
		assert.Empty(t, results)
	})

}

func TestGetSellerS(t *testing.T) {
	//Arrange
	database := []domain.Locality{}

	mockRepository := locality.MockRepository{
		DataMock: database,
		Error:    "",
	}

	service := NewService(&mockRepository)

	//Act
	_, err := service.GetAllSellersByLocality(context.TODO(), "1759")

	//Assert
	assert.Nil(t, err)
}

func TestGetCarriesS(t *testing.T) {
	//Arrange
	database := []domain.Locality{}

	mockRepository := locality.MockRepository{
		DataMock: database,
		Error:    "",
	}

	service := NewService(&mockRepository)

	//Act
	_, err := service.GetCarriesReport(context.TODO(), "1759")

	//Assert
	assert.Nil(t, err)
}
