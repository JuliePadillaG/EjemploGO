package carry

import (
	"testing"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/carry"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := carry.NewRepositoryCarry([]domain.Carry{})
	service := NewService(repo)
	t.Run("should create a new carry", func(t *testing.T) {
		// Arrange
		carry := domain.Carry{
			CID:         "DHD",
			Locality_id: 1,
		}

		// Act
		id, err := service.Save(carry)

		// Assert
		assert.Equal(t, err, nil)
		assert.Equal(t, id, 1)
	})
	t.Run("should not create a new carry with same carrycode", func(t *testing.T) {
		// Arrange
		carry := domain.Carry{
			CID:         "DHD",
			Locality_id: 1,
		}

		// Act
		id, err := service.Save(carry)

		// Assert
		assert.Equal(t, "carry code already exists", err.Error())
		assert.Equal(t, 0, id)
	})
	t.Run("should not create a new carry with non existent locality", func(t *testing.T) {
		// Arrange
		carry := domain.Carry{
			CID:         "DHK",
			Locality_id: 2,
		}

		// Act
		id, err := service.Save(carry)

		// Assert
		assert.Equal(t, "locality code doesn't exists", err.Error())
		assert.Equal(t, 0, id)
	})
}
