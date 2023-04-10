package seller

import (
	"testing"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/sellers"
	"github.com/stretchr/testify/assert"
)

func TestCreateOK(t *testing.T) {
	//Arrange
	database := []domain.Seller{}

	id := 1

	mockRepository := sellers.MockRepository{
		DataMock: database,
		Error:    "",
	}

	service := NewService(&mockRepository)

	//Act
	results, err := service.Save(89, 6700, "Empresa", "Direccion", "92388372")

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, id, results)
}

func TestCreateConflict(t *testing.T) {
	//Arrange
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}

	mockRepository := sellers.MockRepository{
		DataMock: database,
		Error:    "",
	}

	service := NewService(&mockRepository)

	//Act
	_, err := service.Save(19, 6700, "Empresa", "Direccion", "92388372")
	_, err2 := service.Save(0, 6700, "Empresa", "Direccion", "92388372")
	_, err3 := service.Save(11, 6700, "", "Direccion", "92388372")
	_, err4 := service.Save(11, 6700, "Empresa", "", "92388372")
	_, err5 := service.Save(11, 6700, "Empresa", "Direccion", "")
	_, err6 := service.Save(-10, 6700, "Empresa", "Direccion", "92388372")

	//Assert
	assert.ErrorContains(t, err, "cid already exists")
	assert.ErrorContains(t, err2, "field required")
	assert.ErrorContains(t, err3, "field required")
	assert.ErrorContains(t, err4, "field required")
	assert.ErrorContains(t, err5, "field required")
	assert.ErrorContains(t, err6, "incorrect field content")
}

func TestServiceGetAll(t *testing.T) {
	// Arrange.
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
		{
			ID:          2,
			CID:         23,
			CompanyName: "Motrola",
			Address:     "Calle real 11",
			Telephone:   "91284346",
		},
	}

	mockRepository := sellers.MockRepository{
		DataMock: database,
		Error:    "",
	}

	service := NewService(&mockRepository)

	// Act.
	results, err := service.GetAll()

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, 2, len(results))
}

func TestFindByIdNonExistent(t *testing.T) {
	// Arrange.
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}

	mockRepository := sellers.MockRepository{
		DataMock: database,
		Error:    "",
	}

	service := NewService(&mockRepository)

	// Act.
	results, err := service.Get(2)

	// Assert.
	assert.NotNil(t, err)
	assert.Empty(t, results)
}

func TestFindByIdExistent(t *testing.T) {
	// Arrange.
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}

	mockRepository := sellers.MockRepository{
		DataMock: database,
		Error:    "",
	}

	service := NewService(&mockRepository)

	// Act.
	results, err := service.Get(1)

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, database[0], results)
}

func TestUpdateOK(t *testing.T) {
	// Arrange.
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}

	mockRepository := sellers.MockRepository{
		DataMock: database,
		Error:    "",
	}

	service := NewService(&mockRepository)

	new := domain.Seller{ID: 1, CID: 34}
	esperado := domain.Seller{
		ID:          1,
		CID:         34,
		CompanyName: "LG",
		Address:     "Avenida 11122",
		Telephone:   "0303456",
	}

	// Act.
	results, err := service.Update(new)

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, esperado, results)
}

func TestUpdateNonExistent(t *testing.T) {
	// Arrange.
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}

	mockRepository := sellers.MockRepository{
		DataMock: database,
		Error:    "",
	}

	service := NewService(&mockRepository)

	new := domain.Seller{ID: 2}

	// Act.
	results, err := service.Update(new)

	// Assert.
	assert.NotNil(t, err)
	assert.Empty(t, results)
}

func TestDeleteNonExistent(t *testing.T) {
	// Arrange.
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}

	mockRepository := sellers.MockRepository{
		DataMock: database,
		Error:    "",
	}

	service := NewService(&mockRepository)

	// Act.
	err := service.Delete(2)

	// Assert.
	assert.NotNil(t, err)
}

func TestDeleteOK(t *testing.T) {
	// Arrange.
	database := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
		{
			ID:          2,
			CID:         23,
			CompanyName: "Motrola",
			Address:     "Calle real 11",
			Telephone:   "91284346",
		},
	}

	expected := []domain.Seller{
		{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		},
	}

	mockRepository := sellers.MockRepository{
		DataMock: database,
		Error:    "",
	}

	service := NewService(&mockRepository)

	// Act.
	err := service.Delete(2)

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, expected, mockRepository.DataMock)
}
