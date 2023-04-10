package section

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/stretchr/testify/assert"
)

var (
	errorExec    = errors.New("error in Exec")
	errorPrepare = errors.New("error in Prepare")
)

var FakeSection = []domain.Section{
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

func TestGetAllSections(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Get all sections", func(t *testing.T) {

		colums := []string{"ID", "SectionNumber", "CurrentTemperature", "MinimumTemperature", "CurrentCapacity", "MinimumCapacity", "MaximumCapacity", "WarehouseID", "ProductTypeID"}
		rows := sqlmock.NewRows(colums)

		rows.AddRow(FakeSection[0].ID, FakeSection[0].SectionNumber, FakeSection[0].CurrentTemperature, FakeSection[0].MinimumTemperature, FakeSection[0].CurrentCapacity, FakeSection[0].MinimumCapacity, FakeSection[0].MaximumCapacity, FakeSection[0].WarehouseID, FakeSection[0].ProductTypeID)

		rows.AddRow(FakeSection[1].ID, FakeSection[1].SectionNumber, FakeSection[1].CurrentTemperature, FakeSection[1].MinimumTemperature, FakeSection[1].CurrentCapacity, FakeSection[1].MinimumCapacity, FakeSection[1].MaximumCapacity, FakeSection[1].WarehouseID, FakeSection[1].ProductTypeID)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM sections`)).WillReturnRows(rows)

		repo := NewRepository(db)
		sections, err := repo.GetAll(context.Background())

		assert.Nil(t, err)
		t.Log("total sections", sections)
		assert.Equal(t, 2, len(sections))

	})
}

func TestGetByIdSect(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Get section by id", func(t *testing.T) {
		colums := []string{"ID", "SectionNumber", "CurrentTemperature", "MinimumTemperature", "CurrentCapacity", "MinimumCapacity", "MaximumCapacity", "WarehouseID", "ProductTypeID"}
		rows := sqlmock.NewRows(colums)

		rows.AddRow(FakeSection[0].ID, FakeSection[0].SectionNumber, FakeSection[0].CurrentTemperature, FakeSection[0].MinimumTemperature, FakeSection[0].CurrentCapacity, FakeSection[0].MinimumCapacity, FakeSection[0].MaximumCapacity, FakeSection[0].WarehouseID, FakeSection[0].ProductTypeID)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM sections WHERE id=?;`)).WillReturnRows(rows)

		t.Log("section", FakeSection[0].ID)

		repo := NewRepository(db)
		section, err := repo.Get(context.Background(), 1)

		t.Log("section", section)
		assert.Nil(t, err)
		assert.Equal(t, 1, section.ID)
	})
}

func TestExistsSect(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Get section by id", func(t *testing.T) {
		colums := []string{"ID", "SectionNumber", "CurrentTemperature", "MinimumTemperature", "CurrentCapacity", "MinimumCapacity", "MaximumCapacity", "WarehouseID", "ProductTypeID"}
		rows := sqlmock.NewRows(colums)

		rows.AddRow(FakeSection[0].ID, FakeSection[0].SectionNumber, FakeSection[0].CurrentTemperature, FakeSection[0].MinimumTemperature, FakeSection[0].CurrentCapacity, FakeSection[0].MinimumCapacity, FakeSection[0].MaximumCapacity, FakeSection[0].WarehouseID, FakeSection[0].ProductTypeID)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT section_number FROM sections WHERE section_number=?;`)).WillReturnRows(rows)

		repo := NewRepository(db)
		section := repo.Exists(context.Background(), 1)
		assert.False(t, section)
	})
}

func TestSaveSect(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Save section Ok", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`))
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`)).WillReturnResult(sqlmock.NewResult(1, 1))

		columns := []string{"ID", "SectionNumber", "CurrentTemperature", "MinimumTemperature", "CurrentCapacity", "MinimumCapacity", "MaximumCapacity", "WarehouseID", "ProductTypeID"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(FakeSection[0].ID, FakeSection[0].SectionNumber, FakeSection[0].CurrentTemperature, FakeSection[0].MinimumTemperature, FakeSection[0].CurrentCapacity, FakeSection[0].MinimumCapacity, FakeSection[0].MaximumCapacity, FakeSection[0].WarehouseID, FakeSection[0].ProductTypeID)

		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM sections WHERE id=?;`)).WithArgs(1).WillReturnRows(rows)

		repo := NewRepository(db)
		id, err := repo.Save(context.Background(), FakeSection[0])
		assert.NoError(t, err)

		section, err := repo.Get(context.TODO(), id)
		assert.NoError(t, err)

		assert.NotNil(t, section)
		assert.Equal(t, FakeSection[0].ID, section.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Fail Exec", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`))
		mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`)).WillReturnError(errorExec)

		repo := NewRepository(db)

		id, err := repo.Save(context.Background(), FakeSection[0])

		assert.EqualError(t, err, errorExec.Error())
		assert.Equal(t, 0, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
	t.Run("Fail Prepare", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(`INSERT INTO sections (section_number, current_temperature, minimum_temperature, current_capacity, minimum_capacity, maximum_capacity, warehouse_id, product_type_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?);`)).WillReturnError(errorPrepare)

		repo := NewRepository(db)
		id, err := repo.Save(context.Background(), FakeSection[0])

		assert.EqualError(t, err, errorPrepare.Error())
		assert.Equal(t, 0, id)

	})
}

func TestUpdateSect(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	t.Run("Update Ok", func(t *testing.T) {
		expected := domain.Section{
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

		query := regexp.QuoteMeta("UPDATE sections SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, maximum_capacity=?, warehouse_id=?, product_type_id=? WHERE id=?;")

		mock.ExpectPrepare(query)
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		repository := NewRepository(db)

		err = repository.Update(context.TODO(), expected)

		assert.NoError(t, err)
	})
	t.Run("Fail Prepare", func(t *testing.T) {
		expected := domain.Section{
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

		query := regexp.QuoteMeta("UPDATE sections SET section_number=?, current_temperature=?, minimum_temperature=?, current_capacity=?, minimum_capacity=?, maximum_capacity=?, warehouse_id=?, product_type_id=? WHERE id=?;")
		mock.ExpectPrepare(query).WillReturnError(errors.New(""))
		repository := NewRepository(db)
		err = repository.Update(context.TODO(), expected)
		assert.Error(t, err)
	})
	t.Run("Fail Exec", func(t *testing.T) {

		expected := domain.Section{
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
		query := regexp.QuoteMeta("UPDATE employees SET first_name=?, last_name=?, warehouse_id=?  WHERE id=?")
		mock.ExpectPrepare(query)
		mock.ExpectExec(query).WillReturnError(errors.New(""))
		repository := NewRepository(db)

		err = repository.Update(context.TODO(), expected)

		assert.Error(t, err)
	})
}

func TestErrors(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Delete Execute Conflict", func(t *testing.T) {

		query := regexp.QuoteMeta("DELETE FROM sections WHERE id=?;")
		mock.ExpectPrepare(query)
		mock.ExpectExec(query).WillReturnError(errors.New(""))
		repository := NewRepository(db)

		err = repository.Delete(context.TODO(), 1)

		assert.Error(t, err)

	})
	t.Run("Delete Prepare Conflict", func(t *testing.T) {

		query := regexp.QuoteMeta("DELETE FROM sections WHERE id=?")
		mock.ExpectPrepare(query).WillReturnError(errors.New(""))
		repository := NewRepository(db)

		err = repository.Delete(context.TODO(), 1)

		assert.Error(t, err)
	})
	t.Run("Delete Row Affected Conflict", func(t *testing.T) {

		query := regexp.QuoteMeta("DELETE FROM sections WHERE id=?;")
		mock.ExpectPrepare(query)
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 0))
		repository := NewRepository(db)

		err = repository.Delete(context.TODO(), 1)

		assert.Error(t, err)
	})
}
