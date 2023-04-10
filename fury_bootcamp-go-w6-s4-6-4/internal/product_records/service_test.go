package product_records

import (
	"testing"
	//"errors"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/productRecords"
	"github.com/stretchr/testify/assert"
)

func TestCreateOk(t *testing.T) {
	mockService := productRecords.MockProductRecordsRepository{
		DataMock: []domain.ProductRecords{},
		Error:    "",
	}

	//Arrange
	product_record_test := domain.ProductRecords{
		ID:             1,
		LastUpdateDate: "2022-12-04",
		PurchasePrice:  20.9,
		SalePrice:      90.8,
		ProductID:      1,
	}

	//Act
	service := NewService(&mockService)
	result, err := service.Save(ctx, product_record_test.LastUpdateDate, product_record_test.PurchasePrice, product_record_test.SalePrice, product_record_test.ProductID)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, product_record_test, result)
}

// func Test_CreateFail(t *testing.T) {
// 	mockService := productRecords.MockProductRecordsRepository{
// 		DataMock: []domain.ProductRecords{},
// 		Error:    "",
// 	}

// 	product_record_test := domain.ProductRecords{
// 		ID:             1,
// 		LastUpdateDate: "2022-12-04",
// 		PurchasePrice:  20.9,
// 		SalePrice:      90.8,
// 		ProductID:      1,
// 	}

// 	product_record_test_fail := domain.ProductRecords{
// 		ID:             1,
// 		LastUpdateDate: "2022-12-04",
// 		PurchasePrice:  20.9,
// 		SalePrice:      90.8,
// 		ProductID:      1,
// 	}

// 	expectedError := errors.New("product_records id already exists")
// 	service := NewService(&mockService)
// 	result, err := service.Save(ctx, product_record_test.LastUpdateDate, product_record_test.PurchasePrice, product_record_test.SalePrice, product_record_test.ProductID)
// 	assert.Nil(t, err)
// 	assert.NotEqual(t, product_record_test_fail, result)
// 	assert.Equal(t, expectedError, err)
// }
