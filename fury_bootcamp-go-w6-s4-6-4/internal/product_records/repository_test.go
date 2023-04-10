package product_records

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/stretchr/testify/assert"
)

var ctx context.Context

func TestRepositoryStoreOK(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare("INSERT INTO product")
	mock.ExpectExec("INSERT INTO product").WillReturnResult(sqlmock.NewResult(1, 1))
	productId := 1

	repository := NewRepository(db)
	product_record := domain.ProductRecords{
		ID:             productId,
		LastUpdateDate: "2022-11-17",
		PurchasePrice:  23.5,
		SalePrice:      67.5,
		ProductID:      1,
	}

	p, err := repository.Save(ctx, product_record)

	assert.NoError(t, err)
	assert.NotZero(t, p)
	assert.Equal(t, product_record.ID, p)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryStoreFail(t *testing.T) {

	pr := domain.ProductRecords{
		ID:             1,
		LastUpdateDate: "2021-11-17",
		PurchasePrice:  89.8,
		SalePrice:      123.9,
		ProductID:      1,
	}

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(SAVE_PRODUCT_RECORD))
	mock.ExpectExec(regexp.QuoteMeta(SAVE_PRODUCT_RECORD)).
		WithArgs(pr.LastUpdateDate, pr.PurchasePrice, pr.SalePrice, pr.ProductID).
		WillReturnError(errors.New("INSERT ERROR"))

	repository := NewRepository(db)

	id, err := repository.Save(context.TODO(), pr)

	t.Log(err)

	assert.Error(t, err)
	assert.Equal(t, 0, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestExistProductRecordOK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	product_record_test := domain.ProductRecords{
		ID:             1,
		LastUpdateDate: "2022-12-04",
		PurchasePrice:  20.9,
		SalePrice:      90.8,
		ProductID:      1,
	}

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	rows.AddRow(product_record_test.ID)
	mock.ExpectQuery(regexp.QuoteMeta(EXIST_PRODUCT_RECORD)).WithArgs(1).WillReturnRows(rows)

	repo := NewRepository(db)

	resp := repo.ExistsProductRecord(context.TODO(), 1)

	assert.True(t, resp)
}

func TestExistProductRecordFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	mock.ExpectQuery(regexp.QuoteMeta(EXIST_PRODUCT_RECORD)).WithArgs(2).WillReturnRows(rows)

	repo := NewRepository(db)

	resp := repo.ExistsProductRecord(context.TODO(), 2)

	assert.False(t, resp)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUniqueProductOK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	product_test := domain.Product{
		ID:             1,
		Description:    "producto congelado",
		ExpirationRate: 2,
		FreezingRate:   3,
		Height:         20.1,
		Length:         30.2,
		Netweight:      15.2,
		ProductCode:    "j3l4k5",
		RecomFreezTemp: 20.0,
		Width:          30.6,
		ProductTypeID:  2,
		SellerID:       1,
	}

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	rows.AddRow(product_test.ID)
	mock.ExpectQuery(regexp.QuoteMeta(UNIQUE_PRODUCT)).WithArgs(1).WillReturnRows(rows)

	repo := NewRepository(db)

	resp := repo.UniqueProduct(context.TODO(), 1)

	assert.True(t, resp)
}
