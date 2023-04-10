package buyer

import (
	"context"
	"errors"
	"testing"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/buyer"
	"github.com/stretchr/testify/assert"
)


func TestServiceFindAll(t *testing.T) {
	// Arrange.
        database := []domain.Buyer{
                {
        		ID:           1,
                        CardNumberID: "232345",
                        FirstName:    "hello",
                        LastName:     "World",
                },
                {
        		ID:           2,
                        CardNumberID: "232346",
                        FirstName:    "hello_2",
                        LastName:     "World_2",
                },
        }

	mockRepository := buyer.MockRepository{
		DataMock: database,
	}

	service := NewService(&mockRepository)
        ctx := context.Background()

	// Act.
	buyers, err := service.GetAll(ctx)

	// Assert.
	assert.Nil(t, err)
	assert.Equal(t, len(mockRepository.DataMock), len(buyers))
	assert.Equal(t, mockRepository.DataMock, buyers)
}

func TestServiceFindByIdExistent(t *testing.T) {
	// Arrange.
        id := 1
        database := []domain.Buyer{
                {
                        ID:           id,
                        CardNumberID: "232345",
                        FirstName:    "hello",
                        LastName:     "World",
                },
                {
                        ID:           2,
                        CardNumberID: "232346",
                        FirstName:    "hello_2",
                        LastName:     "World_2",
                },
        }

	mockRepository := buyer.MockRepository{
		DataMock: database,
	}

	service := NewService(&mockRepository)
        ctx := context.Background()

	// Act.
	buyer, err := service.Get(ctx, id)

	// Assert.
	assert.Nil(t, err)
        assert.Equal(t, mockRepository.DataMock[0].ID, buyer.ID)
        assert.Equal(t, mockRepository.DataMock[0], buyer)
}


func TestServiceFindByIdNotExistent(t *testing.T) {
	// Arrange.
        id := 1
        errMessage := "buyer not found"
        database := []domain.Buyer{}

	mockRepository := buyer.MockRepository{
		DataMock: database,
                Error: errMessage,
	}

	service := NewService(&mockRepository)
        ctx := context.Background()

	// Act.
	buyer, err := service.Get(ctx, id)

	// Assert.
	assert.NotNil(t, err)
        assert.Empty(t, buyer)
}

func TestServiceCreateOk(t *testing.T) {
        // Arrange
        newBuyer := domain.Buyer{
                CardNumberID: "232345",
                FirstName:    "hello",
                LastName:     "World",
        }

        mockRepository := buyer.MockRepository{
                DataMock: []domain.Buyer{},
                ExistWasCalled: false,
                GetWasCalled: false,
        }

        productIDExpected := 1

        // Act
        service := NewService(&mockRepository)
        ctx := context.Background()

        result, err := service.Save(ctx, newBuyer.CardNumberID, newBuyer.FirstName, newBuyer.LastName)

        // Assert
        assert.Nil(t, err)
        assert.True(t, mockRepository.ExistWasCalled)
        // assert.True(t, mockRepository.GetWasCalled)
        assert.Equal(t, mockRepository.DataMock[0], result)
        assert.Equal(t, productIDExpected, result.ID)
}

func TestServiceCreateConflict(t *testing.T) {
	// Arrange
        errMessage := errors.New("duplicate cardNumberID")
        cardNumberID := "232345"
        database := []domain.Buyer{
                {
                        ID:           1,
                        CardNumberID: cardNumberID,
                        FirstName:    "hello",
                        LastName:     "World",
                },
                {
                        ID:           2,
                        CardNumberID: "232346",
                        FirstName:    "hello_2",
                        LastName:     "World_2",
                },
        }

        newBuyer := domain.Buyer{
                CardNumberID: cardNumberID,
                FirstName:    "hello_2",
                LastName:     "World_2",
        }

	mockRepository := buyer.MockRepository{
		DataMock: database,
                ExistWasCalled: false,
                GetWasCalled: false,
	}

        productIDExpected := 0

        // Act
	service := NewService(&mockRepository)
        ctx := context.Background()
        result, err := service.Save(ctx, newBuyer.CardNumberID, newBuyer.FirstName, newBuyer.LastName)

        // t.Log(err)

        // Assert
	assert.NotNil(t, err)
        assert.EqualError(t, errMessage, err.Error())
	assert.Empty(t, result)
        assert.True(t, mockRepository.ExistWasCalled)
        assert.False(t, mockRepository.GetWasCalled)
	assert.Equal(t, productIDExpected, result.ID)
}

func TestServiceUpdateExistent(t *testing.T) {
	// Arrange.
        database := []domain.Buyer{
                {
                        ID:           1,
                        CardNumberID: "219784",
                        FirstName:    "hello",
                        LastName:     "World",
                },
                {
                        ID:           2,
                        CardNumberID: "232346",
                        FirstName:    "hello_2",
                        LastName:     "World_2",
                },
        }

        afterBuyerToUpdate := database[1]

        buyerToUpdate := domain.Buyer{
                ID:           2,
                CardNumberID: "232346",
                FirstName:    "hello_3",
                LastName:     "World_3",
        }

	mockRepository := buyer.MockRepository{
		DataMock: database,
                GetWasCalled: false,
	}

	service := NewService(&mockRepository)
        ctx := context.Background()

	// Act.
        buyer, err := service.Update(ctx, buyerToUpdate.ID, buyerToUpdate.FirstName, buyerToUpdate.LastName)

	// Assert.
	assert.Nil(t, err)
        assert.Equal(t, buyerToUpdate, buyer)
        assert.True(t, mockRepository.GetWasCalled)
        assert.NotEqual(t, afterBuyerToUpdate, buyer)
}




func TestServiceUpdateNonExistent(t *testing.T) {
	// Arrange.
        errMessage := errors.New("buyer not found")
        database := []domain.Buyer{
                {
                        ID:           1,
                        CardNumberID: "219784",
                        FirstName:    "hello",
                        LastName:     "World",
                },
        }

        buyerToUpdate := domain.Buyer{
                ID:           2,
                CardNumberID: "232346",
                FirstName:    "hello_3",
                LastName:     "World_3",
        }

	mockRepository := buyer.MockRepository{
		DataMock: database,
                GetWasCalled: false,
                Error: errMessage.Error(),
	}

	service := NewService(&mockRepository)
        ctx := context.Background()

	// Act.
        buyer, err := service.Update(ctx, buyerToUpdate.ID, buyerToUpdate.FirstName, buyerToUpdate.LastName)

	// Assert.
	assert.NotNil(t, err)
        assert.EqualError(t, errMessage, err.Error())
        assert.Empty(t, buyer)
        assert.True(t, mockRepository.GetWasCalled)
}

func TestServiceDeleteOk(t *testing.T) {
	// Arrange.
        buyerIDToDelete := 1
        database := []domain.Buyer{
                {
                        ID:           buyerIDToDelete,
                        CardNumberID: "219784",
                        FirstName:    "hello",
                        LastName:     "World",
                },
                {
                        ID:           2,
                        CardNumberID: "232346",
                        FirstName:    "hello_2",
                        LastName:     "World_2",
                },
        }

        dbCopy := make([]domain.Buyer, len(database))
        copy(dbCopy, database)

	mockRepository := buyer.MockRepository{
		DataMock: dbCopy,
	}

	service := NewService(&mockRepository)
        ctx := context.Background()

	// Act.
        err := service.Delete(ctx, buyerIDToDelete)

        // t.Logf("%p mockRepository.DataMock %+v", mockRepository.DataMock, mockRepository.DataMock)
        // t.Logf("%p database %+v", database, database)

        // Assert.
	assert.Nil(t, err)
        assert.NotEqual(t, database, mockRepository.DataMock)
        assert.NotEqual(t, len(database), len(mockRepository.DataMock))
}

func TestServiceDeleteNonExistent(t *testing.T) {
	// Arrange.
        buyerIDToDelete := 1
        errMessage := errors.New("buyer not found")

	mockRepository := buyer.MockRepository{
		DataMock: []domain.Buyer{},
                Error: errMessage.Error(),
	}

	service := NewService(&mockRepository)
        ctx := context.Background()

	// Act.
        err := service.Delete(ctx, buyerIDToDelete)

        // t.Logf("%p mockRepository.DataMock %+v", mockRepository.DataMock, mockRepository.DataMock)
        // t.Logf("%p database %+v", database, database)

        // Assert.
	assert.NotNil(t, err)
        assert.EqualError(t, errMessage, err.Error())
}


