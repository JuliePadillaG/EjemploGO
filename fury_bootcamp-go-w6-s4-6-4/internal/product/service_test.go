package product

import (
	"context"
	"errors"
	"testing"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/products"
	"github.com/stretchr/testify/assert"
)

//Declaración de variables globales
var ctx context.Context
var database = []domain.Product{
    {
        ID: 1,
        Description: "Producto congelado",
        ExpirationRate: 10,
        FreezingRate: 10,
        Height: 10,
        Length: 10,
        Netweight: 10,
        ProductCode: "567",
        RecomFreezTemp: 5,
        Width: 10,
        ProductTypeID: 12345,
        SellerID: 6,
    },
    {
        ID: 2,
        Description: "Producto refrigerado",
        ExpirationRate: 20,
        FreezingRate: 9,
        Height: 20,
        Length: 20,
        Netweight: 100,
        ProductCode: "980",
        RecomFreezTemp: 6,
        Width: 20,
        ProductTypeID: 12345,
        SellerID: 7,
    },
}

// User Story asociada: CREATE
// Caso borde: create_ok
// Si contiene los campos necesarios se creará
func TestCreateOk(t *testing.T) {
    
    //Arrange
    p := domain.Product{
        ID: 1,
        Description: "Producto refrigerado",
        ExpirationRate: 20,
        FreezingRate: 9,
        Height: 20,
        Length: 20,
        Netweight: 100,
        ProductCode: "980",
        RecomFreezTemp: 6,
        Width: 20,
        ProductTypeID: 12345,
        SellerID: 7,
    }

    mockRepository := products.MockRepositoryProduct{}

    //Act
    service := NewService(&mockRepository)
    result, err := service.Save(ctx, p.Description, p.ExpirationRate, p.FreezingRate, p.Height, p.Length, p.Netweight, p.ProductCode, p.RecomFreezTemp, p.Width, p.ProductTypeID, p.SellerID)
    
    //Assert
    assert.Nil(t, err)
    assert.Equal(t, p, result)

}

// User Story asociada: CREATE
// Caso borde: create_conflict
// Si el product_code ya existe no podrá ser creado
func TestCreateConflict(t *testing.T) {
    
    //Arrange
    p := domain.Product {
        ID: 1,
        Description: "Producto refrigerado",
        ExpirationRate: 20,
        FreezingRate: 9,
        Height: 20,
        Length: 20,
        Netweight: 100,
        ProductCode: "980",
        RecomFreezTemp: 6,
        Width: 20,
        ProductTypeID: 12345,
        SellerID: 7,
    }
    expectedError := errors.New("product_code already exists")
    mockRepository := products.MockRepositoryProduct{DataMock: database}

    //Act
    service := NewService(&mockRepository)
    result, err := service.Save(ctx, p.Description, p.ExpirationRate, p.FreezingRate, p.Height, p.Length, p.Netweight, p.ProductCode, p.RecomFreezTemp, p.Width, p.ProductTypeID, p.SellerID)

    //Assert
    assert.NotNil(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, result)

}

// User Story asociada: READ
// Caso borde: find_all
// Si la lista posee “n” elementos devolverá un cantidad de los elementos totales
func TestFindAll(t *testing.T) {

	// Arrange
    mockRepository := products.MockRepositoryProduct{
        DataMock: database,
        //ErrWrite: "",
        //ErrRead: "",
    }
	
	// Act
    service := NewService(&mockRepository)
    result, err := service.GetAll(ctx)

	// Assert
    assert.Nil(t, err)
    assert.Equal(t, mockRepository.DataMock, result)
}

// User Story asociada: READ
// Caso borde: find_by_id_non_existent
// Si el elemento buscado por id no existe retorna nulo
func TestFindByIdNonExistent(t *testing.T) {
    
    //Arrange
    expectedError := errors.New("product not found")
    
    mockRepository := products.MockRepositoryProduct{
        DataMock: database,
    }

    //Act
    service := NewService(&mockRepository)
    result, err := service.Get(ctx, 5)

    //Assert
    assert.NotNil(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, result)
}

// User Story asociada: READ
// Caso borde: find_by_id_existent
// Si el elemento buscado por id existe devolverá la información del elemento solicitado
func TestFindByIdExistent(t *testing.T) {
    
    //Arrange
    pr := database[0]
    mockRepository := products.MockRepositoryProduct{
        DataMock: database,
    }

    //Act
    service := NewService(&mockRepository)
    result, err := service.Get(ctx, pr.ID)

    //Assert
    assert.Nil(t, err)
    assert.Equal(t, pr, result)
}

// User Story asociada: UPDATE
// Caso borde: update_existent
// Cuando la actualización de datos sea exitosa se devolverá el producto con la información actualizada
func TestUpdateExistent(t *testing.T) {
    
    //Arrange
    p := domain.Product {
        ID: 1,
        Description: "Producto refrigerado",
        ExpirationRate: 30,
        FreezingRate: 7,
        Height: 30,
        Length: 30,
        Netweight: 100,
        ProductCode: "980",
        RecomFreezTemp: 6,
        Width: 20,
        ProductTypeID: 12345,
        SellerID: 7,
    }

    mockRepository := products.MockRepositoryProduct{
        DataMock: database,
    }

    //Act
    service := NewService(&mockRepository)
    result, err := service.Update(ctx, p.ID, p.Description, &p.ExpirationRate, &p.FreezingRate, &p.Height, &p.Length, &p.Netweight, p.ProductCode, &p.RecomFreezTemp, &p.Width, &p.ProductTypeID, &p.SellerID)

    //Assert
    assert.Nil(t, err)
    assert.Equal(t, p, result)
    assert.Equal(t, p, mockRepository.DataMock[0])
}

// User Story asociada: UPDATE
// Caso borde: update_non_existent
// Si el product que se desea actualizar no existe se devolverá null.
func TestUpdateNonExistent(t *testing.T) {
    
    //Arrange
    expectedError := errors.New("product not found")

    p := domain.Product{
        ID: 8,
        Description: "Producto refrigerado",
        ExpirationRate: 20,
        FreezingRate: 8,
        Height: 50,
        Length: 50,
        Netweight: 100,
        ProductCode: "687997",
        RecomFreezTemp: 6,
        Width: 20,
        ProductTypeID: 12345,
        SellerID: 7,
    }

    mockRepository := products.MockRepositoryProduct{
        DataMock: database,
    }
    //Act
    service := NewService(&mockRepository)
    result, err := service.Update(ctx, p.ID, p.Description, &p.ExpirationRate, &p.FreezingRate, &p.Height, &p.Length, &p.Netweight, p.ProductCode, &p.RecomFreezTemp, &p.Width, &p.ProductTypeID, &p.SellerID)

    //Assert
    assert.NotNil(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, result)
}

// User Story asociada: DELETE
// Caso borde: delete_non_existent
// Cuando el producto no existe se devolverá null.
func TestDeleteNonExistent(t *testing.T) {
    
    //Arrange
    expectedError := errors.New("product not found")

    p := domain.Product {
        ID: 8,
        Description: "Producto refrigerado",
        ExpirationRate: 20,
        FreezingRate: 8,
        Height: 50,
        Length: 50,
        Netweight: 100,
        ProductCode: "687997",
        RecomFreezTemp: 6,
        Width: 20,
        ProductTypeID: 12345,
        SellerID: 7,
    }
    mockRepository := products.MockRepositoryProduct{
        DataMock: database,
    }

    //Act
    service := NewService(&mockRepository)
    err := service.Delete(ctx, p.ID)

    //Assert
    assert.NotNil(t, err)
    assert.Equal(t, expectedError, err)

}

// User Story asociada: DELETE
// Caso borde: delete_ok
// Si la eliminación es exitosa el elemento no aparecerá en la lista.
func TestDeleteOk(t *testing.T) {
    
    //Arrange
    p := domain.Product{
        ID: 1,
        Description: "Producto congelado",
        ExpirationRate: 10,
        FreezingRate: 10,
        Height: 10,
        Length: 10,
        Netweight: 10,
        ProductCode: "567",
        RecomFreezTemp: 5,
        Width: 10,
        ProductTypeID: 12345,
        SellerID: 6,
    }

    expectedError := errors.New("product not found")

    mockRepository := products.MockRepositoryProduct{
        DataMock: database,
    }

    //Act
    service := NewService(&mockRepository)
    err := service.Delete(ctx, p.ID)
    result, errService := service.Get(ctx, p.ID)

    //Assert
    assert.Nil(t, err)
    assert.Equal(t, expectedError, errService)
    assert.Empty(t, result)
}

