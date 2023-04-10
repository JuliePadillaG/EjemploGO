package warehouse

import (
	"testing"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/test/mocks/warehouse"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := warehouse.NewRepositoryWarehouse([]domain.Warehouse{})
	service := NewService(repo)
	t.Run("should create a new warehouse", func(t *testing.T) {
		// Arrange
		minCapacity := 10
		minTem := 10
		warehouse := domain.Warehouse{
			Address:            "Monroe 400",
			Telephone:          "44470000",
			WarehouseCode:      "DHD",
			MinimumCapacity:    &minCapacity,
			MinimumTemperature: &minTem,
		}

		// Act
		id, err := service.Save(warehouse)

		// Assert
		assert.Equal(t, err, nil)
		assert.Equal(t, id, 1)
	})
	t.Run("should not create a new warehouse with same warehousecode", func(t *testing.T) {
		// Arrange
		minCapacity := 10
		minTem := 10
		warehouse := domain.Warehouse{
			Address:            "Monroe 400",
			Telephone:          "44470000",
			WarehouseCode:      "DHD",
			MinimumCapacity:    &minCapacity,
			MinimumTemperature: &minTem,
		}

		// Act
		id, err := service.Save(warehouse)

		// Assert
		assert.Equal(t, "warehouse code already exists", err.Error())
		assert.Equal(t, 0, id)
	})
}

func TestRead(t *testing.T) {
	// Arrange
	warehouses := []domain.Warehouse{
		{
			ID:                 1,
			Address:            "Monroe 400",
			Telephone:          "44470000",
			WarehouseCode:      "DHD",
			MinimumCapacity:    nil,
			MinimumTemperature: nil,
		},
		{
			ID:                 2,
			Address:            "Monroe 400",
			Telephone:          "44470000",
			WarehouseCode:      "DHK",
			MinimumCapacity:    nil,
			MinimumTemperature: nil,
		},
	}

	repo := warehouse.NewRepositoryWarehouse(warehouses)
	service := NewService(repo)
	t.Run("should get a warehouse", func(t *testing.T) {

		// Act
		result, err := service.GetAll()

		// Assert
		assert.Equal(t, err, nil)
		assert.Equal(t, warehouses, result)
	})
	t.Run("should get a warehouse by id", func(t *testing.T) {

		// Act
		result, err := service.Get(1)

		// Assert
		assert.Equal(t, err, nil)
		assert.Equal(t, warehouses[0], result)
	})
	t.Run("should not get a warehouse by id", func(t *testing.T) {

		// Act
		result, err := service.Get(3)

		// Assert
		assert.Equal(t, "warehouse not found", err.Error())
		assert.Equal(t, domain.Warehouse{}, result)
	})
}

func TestUpdate(t *testing.T) {
	// Arrange
	warehouses := []domain.Warehouse{
		{
			ID:                 1,
			Address:            "Monroe 400",
			Telephone:          "44470000",
			WarehouseCode:      "DHD",
			MinimumCapacity:    nil,
			MinimumTemperature: nil,
		},
		{
			ID:                 2,
			Address:            "Monroe 400",
			Telephone:          "44470000",
			WarehouseCode:      "DHK",
			MinimumCapacity:    nil,
			MinimumTemperature: nil,
		},
	}

	repo := warehouse.NewRepositoryWarehouse(warehouses)
	service := NewService(repo)
	t.Run("should update a warehouse", func(t *testing.T) {
		// Arrange
		minCapacity := 10
		minTem := 10
		warehouse := domain.Warehouse{
			ID:                 1,
			Address:            "Monroe 300",
			Telephone:          "44470000",
			WarehouseCode:      "DHD",
			MinimumCapacity:    &minCapacity,
			MinimumTemperature: &minTem,
		}

		// Act
		result, err := service.Update(warehouse, 1)

		// Assert
		assert.Equal(t, err, nil)
		assert.Equal(t, warehouse, result)
	})
	t.Run("should update a warehouse with some fields", func(t *testing.T) {
		// Arrange
		minCapacity := 10
		minTem := 10
		warehouse := domain.Warehouse{
			MinimumCapacity:    &minCapacity,
			MinimumTemperature: &minTem,
		}

		excepted := domain.Warehouse{
			ID:                 2,
			Address:            "Monroe 400",
			Telephone:          "44470000",
			WarehouseCode:      "DHK",
			MinimumCapacity:    &minCapacity,
			MinimumTemperature: &minTem,
		}
		// Act
		result, err := service.Update(warehouse, 2)

		// Assert
		assert.Equal(t, err, nil)
		assert.Equal(t, excepted, result)
	})
	t.Run("should not update a warehouse", func(t *testing.T) {
		// Arrange
		minCapacity := 10
		minTem := 10
		warehouse := domain.Warehouse{
			ID:                 3,
			Address:            "Monroe 400",
			Telephone:          "44470000",
			WarehouseCode:      "DHD",
			MinimumCapacity:    &minCapacity,
			MinimumTemperature: &minTem,
		}

		// Act
		result, err := service.Update(warehouse, 3)

		// Assert
		assert.Equal(t, "warehouse not found", err.Error())
		assert.Equal(t, domain.Warehouse{}, result)
	})
}

func TestDelete(t *testing.T) {
	// Arrange
	warehouses := []domain.Warehouse{
		{
			ID:                 1,
			Address:            "Monroe 400",
			Telephone:          "44470000",
			WarehouseCode:      "DHD",
			MinimumCapacity:    nil,
			MinimumTemperature: nil,
		},
	}

	repo := warehouse.NewRepositoryWarehouse(warehouses)
	service := NewService(repo)
	t.Run("should not delete a non-existent warehouse", func(t *testing.T) {
		// Act
		err := service.Delete(3)

		// Assert
		assert.Equal(t, "warehouse not found", err.Error())
	})
	t.Run("should delete a warehouse", func(t *testing.T) {
		// Arrange
		expected := []domain.Warehouse{}

		// Act
		err := service.Delete(1)
		result, _ := service.GetAll()
		// Assert
		assert.Equal(t, err, nil)
		assert.Equal(t, expected, result)
	})
}
