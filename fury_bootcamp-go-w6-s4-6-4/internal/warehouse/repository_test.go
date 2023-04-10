package warehouse

import (
	"context"
	"errors"
	"testing"

	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_Repository_Exists(t *testing.T) {

	t.Run("should return true if warehouse exists", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		columns := []string{"warehouse_code"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow("1")

		mock.ExpectQuery(regexp.QuoteMeta("SELECT warehouse_code FROM warehouses WHERE warehouse_code=?;")).WithArgs("1").WillReturnRows(rows)
		repository := NewRepository(db)
		ctx := context.TODO()

		exists := repository.Exists(ctx, "1")
		assert.True(t, exists)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should return false if warehouse does not exist", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		columns := []string{"warehouse_code"}
		rows := sqlmock.NewRows(columns)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT warehouse_code FROM warehouses WHERE warehouse_code=?;")).WithArgs("1").WillReturnRows(rows)
		repository := NewRepository(db)
		ctx := context.TODO()

		exists := repository.Exists(ctx, "1")
		assert.False(t, exists)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func Test_Repository_Save(t *testing.T) {

	t.Run("should save a warehouse", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO warehouses"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO warehouses")).WillReturnResult(sqlmock.NewResult(1, 1))
		repository := NewRepository(db)
		defaultNumber := 1
		warehouse := domain.Warehouse{
			Address:            "address",
			Telephone:          "telephone",
			WarehouseCode:      "warehouseCode",
			MinimumCapacity:    &defaultNumber,
			MinimumTemperature: &defaultNumber,
		}
		ctx := context.TODO()

		id, err := repository.Save(ctx, warehouse)
		assert.NoError(t, err)
		assert.Equal(t, 1, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should return exec error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO warehouses"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO warehouses")).WillReturnError(errors.New("insert exec error"))
		repository := NewRepository(db)
		defaultNumber := 1
		warehouse := domain.Warehouse{
			Address:            "address",
			Telephone:          "telephone",
			WarehouseCode:      "warehouseCode",
			MinimumCapacity:    &defaultNumber,
			MinimumTemperature: &defaultNumber,
		}
		ctx := context.TODO()

		_, err = repository.Save(ctx, warehouse)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should return prepare error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO warehouses")).WillReturnError(errors.New("prepare error"))
		repository := NewRepository(db)
		defaultNumber := 1
		warehouse := domain.Warehouse{
			Address:            "address",
			Telephone:          "telephone",
			WarehouseCode:      "warehouseCode",
			MinimumCapacity:    &defaultNumber,
			MinimumTemperature: &defaultNumber,
		}
		ctx := context.TODO()

		_, err = repository.Save(ctx, warehouse)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

}

func Test_Repository_Get(t *testing.T) {

	t.Run("should return a warehouse", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		columns := []string{"id", "address", "telephone", "warehouse_code", "minimum_capacity", "minimum_temperature"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(1, "address", "telephone", "warehouseCode", 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM warehouses WHERE id=?;")).WithArgs(1).WillReturnRows(rows)
		repository := NewRepository(db)
		ctx := context.TODO()

		warehouse, err := repository.Get(ctx, 1)
		assert.NoError(t, err)
		assert.Equal(t, 1, warehouse.ID)
		assert.Equal(t, "address", warehouse.Address)
		assert.Equal(t, "telephone", warehouse.Telephone)
		assert.Equal(t, "warehouseCode", warehouse.WarehouseCode)
		assert.Equal(t, 1, *warehouse.MinimumCapacity)
		assert.Equal(t, 1, *warehouse.MinimumTemperature)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should return query error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM warehouses WHERE id=?;")).WithArgs(1).WillReturnError(errors.New("query error"))
		repository := NewRepository(db)
		ctx := context.TODO()

		_, err = repository.Get(ctx, 1)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func Test_Repository_GetALL(t *testing.T) {

	t.Run("should return a list of warehouses", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		columns := []string{"id", "address", "telephone", "warehouse_code", "minimum_capacity", "minimum_temperature"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(1, "address", "telephone", "warehouseCode", 1, 1)
		rows.AddRow(2, "address", "telephone", "warehouseCode", 1, 1)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM warehouses")).WillReturnRows(rows)
		repository := NewRepository(db)
		ctx := context.TODO()

		warehouses, err := repository.GetAll(ctx)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(warehouses))
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should return query error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM warehouses")).WillReturnError(errors.New("query error"))
		repository := NewRepository(db)
		ctx := context.TODO()

		_, err = repository.GetAll(ctx)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func Test_Repository_Update(t *testing.T) {

	t.Run("should update a warehouse", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE warehouses"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE warehouses")).WithArgs("address", "telephone", "warehouseCode", 1, 1, 1).WillReturnResult(sqlmock.NewResult(1, 1))
		repository := NewRepository(db)
		defaultNumber := 1
		warehouse := domain.Warehouse{
			ID:                 1,
			Address:            "address",
			Telephone:          "telephone",
			WarehouseCode:      "warehouseCode",
			MinimumCapacity:    &defaultNumber,
			MinimumTemperature: &defaultNumber,
		}
		ctx := context.TODO()

		err = repository.Update(ctx, warehouse)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should return exec error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE warehouses"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE warehouses")).WithArgs("address", "telephone", "warehouseCode", 1, 1, 1).WillReturnError(errors.New("update exec error"))
		repository := NewRepository(db)
		defaultNumber := 1
		warehouse := domain.Warehouse{
			ID:                 1,
			Address:            "address",
			Telephone:          "telephone",
			WarehouseCode:      "warehouseCode",
			MinimumCapacity:    &defaultNumber,
			MinimumTemperature: &defaultNumber,
		}
		ctx := context.TODO()

		err = repository.Update(ctx, warehouse)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should return prepare error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE warehouses")).WillReturnError(errors.New("prepare error"))
		repository := NewRepository(db)
		defaultNumber := 1
		warehouse := domain.Warehouse{
			ID:                 1,
			Address:            "address",
			Telephone:          "telephone",
			WarehouseCode:      "warehouseCode",
			MinimumCapacity:    &defaultNumber,
			MinimumTemperature: &defaultNumber,
		}
		ctx := context.TODO()

		err = repository.Update(ctx, warehouse)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func Test_Repository_Delete(t *testing.T) {

	t.Run("should delete a warehouse", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM warehouses"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM warehouses")).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
		repository := NewRepository(db)
		ctx := context.TODO()

		err = repository.Delete(ctx, 1)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should return exec error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM warehouses"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM warehouses")).WithArgs(1).WillReturnError(errors.New("delete exec error"))
		repository := NewRepository(db)
		ctx := context.TODO()

		err = repository.Delete(ctx, 1)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("should return prepare error", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM warehouses")).WillReturnError(errors.New("prepare error"))
		repository := NewRepository(db)
		ctx := context.TODO()

		err = repository.Delete(ctx, 1)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
