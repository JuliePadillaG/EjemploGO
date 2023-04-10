package productbatches

import (
	"context"
	"errors"
	"log"
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

var FakeProductBatches = domain.Product_batches{

	ID:                 1,
	BatchNumber:        123,
	CurrentQuantity:    10,
	CurrentTemperature: 10,
	DueDate:            "2021-10-10",
	InitialQuantity:    10,
	ManufacturingDate:  "2021-10-10",
	ManufacturingHour:  10,
	MinimumTemperature: 10,
	ProductId:          1,
	SectionId:          1,
}

var FakeSection = domain.Section{
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

var FakeReportProduct = domain.ReportProduct{
	SectionId:       1,
	SectionNumber:   1,
	CurrentQuantity: 10,
}

func TestCreateProductBatches(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("Save Ok", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(CREATE_PRODUCT_BATCH))
		mock.ExpectExec(regexp.QuoteMeta(CREATE_PRODUCT_BATCH)).WillReturnResult(sqlmock.NewResult(1, 1))

		columns := []string{"id", "batch_number", "current_quantity", "current_temperature", "due_date", "initial_quantity", "manufacturing_date", "manufacturing_hour", "minimum_temperature", "product_id", "section_id"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(FakeProductBatches.ID, FakeProductBatches.BatchNumber, FakeProductBatches.CurrentQuantity, FakeProductBatches.CurrentTemperature, FakeProductBatches.DueDate, FakeProductBatches.InitialQuantity, FakeProductBatches.ManufacturingDate, FakeProductBatches.ManufacturingHour, FakeProductBatches.MinimumTemperature, FakeProductBatches.ProductId, FakeProductBatches.SectionId)

		mock.ExpectQuery(regexp.QuoteMeta(GET_PRODUCT_BATCH)).WithArgs(1).WillReturnRows(rows)

		repo := NewRepository(db)
		id, err := repo.CreatePB(context.Background(), FakeProductBatches)
		assert.NoError(t, err)

		productBatches, err := repo.GetPB(context.TODO(), id)
		assert.NoError(t, err)

		assert.NotNil(t, productBatches)
		assert.Equal(t, FakeProductBatches.ID, productBatches.ID)
		assert.NoError(t, mock.ExpectationsWereMet())

	})

	t.Run("Fail Exec", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(CREATE_PRODUCT_BATCH))
		mock.ExpectExec(regexp.QuoteMeta(CREATE_PRODUCT_BATCH)).WillReturnError(errorExec)

		repo := NewRepository(db)

		id, err := repo.CreatePB(context.Background(), FakeProductBatches)

		assert.EqualError(t, err, errorExec.Error())
		assert.Equal(t, 0, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Fail Prepare", func(t *testing.T) {
		mock.ExpectPrepare(regexp.QuoteMeta(CREATE_PRODUCT_BATCH)).WillReturnError(errorPrepare)

		repo := NewRepository(db)
		id, err := repo.CreatePB(context.Background(), FakeProductBatches)

		assert.EqualError(t, err, errorPrepare.Error())
		assert.Equal(t, 0, id)
		log.Println(err)
	})

}

func TestReadReportProduct(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("ReadPB Ok", func(t *testing.T) {

		colums := []string{"SectionId", "SectionNumber", "CurrentQuantity"}
		rows := sqlmock.NewRows(colums)

		rows.AddRow(&FakeProductBatches.SectionId, &FakeSection.SectionNumber, &FakeProductBatches.CurrentQuantity)

		mock.ExpectQuery(regexp.QuoteMeta(READ_PRODUCT_BATCH)).WithArgs(1).WillReturnRows(rows)

		repo := NewRepository(db)
		exist, err := repo.ReadPB(context.TODO(), 1)
		log.Println()

		assert.Nil(t, err)
		assert.Equal(t, FakeReportProduct, exist)

	})

}

func TestProductIdExistsOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	colums := []string{"id"}
	rows := sqlmock.NewRows(colums)

	rows.AddRow(FakeProductBatches.ProductId) // El 1 es el argumento que pide en la query de ExistenceProductId (id=?)
	mock.ExpectQuery(regexp.QuoteMeta(EXISTS_PRODUCT_ID)).WithArgs(1).WillReturnRows(rows)

	//instanciamos el repositorio
	repo := NewRepository(db)
	exist := repo.ExistenceProductId(context.TODO(), 1)

	assert.True(t, exist, "El productoId existe")

}

func TestSectionIdExistsOk(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	colums := []string{"id"}
	rows := sqlmock.NewRows(colums)

	rows.AddRow(FakeProductBatches.SectionId) // El 1 es el argumento que pide en la query de ExistenceProductId (id=?)
	mock.ExpectQuery(regexp.QuoteMeta(EXISTS_SECTION_ID)).WithArgs(1).WillReturnRows(rows)

	//instanciamos el repositorio
	repo := NewRepository(db)
	exist := repo.ExistenceSectionId(context.TODO(), 1)

	assert.True(t, exist, "La seccionId existe")

}
