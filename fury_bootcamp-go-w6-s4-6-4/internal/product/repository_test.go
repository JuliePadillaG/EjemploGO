package product

import (
	"context"
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/stretchr/testify/assert"
)

var (
	c context.Context
	Err = errors.New("Error")
)

func TestRepositorySave(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	product_test := domain.Product{
		ID: 1,
		Description: "producto congelado",
		ExpirationRate:2,
		FreezingRate:3,
		Height:20.1,
		Length:30.2,
		Netweight:15.2,
		ProductCode:"j3l4k5",
		RecomFreezTemp:20.0,
		Width:30.6,
		ProductTypeID:2,
		SellerID:1,
	}

	t.Run("Repository Save Ok", func(t *testing.T) {

		mock.ExpectPrepare(regexp.QuoteMeta(SAVE_PRODUCT))
		mock.ExpectExec(regexp.QuoteMeta(SAVE_PRODUCT)).WillReturnResult(sqlmock.NewResult(1, 1))

		columns := []string{"id", "description", "expiration_rate", "freezing_rate", "height", "length", "netweight", "product_code", "recommended_freezing_temperature", "width", "product_type_id", "seller_id"}
		rows := sqlmock.NewRows(columns)
		rows.AddRow(product_test.ID, product_test.Description, product_test.ExpirationRate, product_test.FreezingRate, product_test.Height, product_test.Length, product_test.Netweight, product_test.ProductCode, product_test.RecomFreezTemp, product_test.Width, product_test.ProductTypeID, product_test.SellerID)
		mock.ExpectQuery(regexp.QuoteMeta(GET_PRODUCT_BY_ID)).WithArgs(1).WillReturnRows(rows)

		repository := NewRepository(db)
		ctx := context.TODO()

		newID, err := repository.Save(ctx, product_test)
		assert.NoError(t, err)

		productResult, err := repository.Get(ctx, int(newID))
		assert.NoError(t, err)

		assert.NotNil(t, productResult)
		assert.Equal(t, product_test.ID, productResult.ID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Repository Save Fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta(SAVE_PRODUCT))
		mock.ExpectExec(regexp.QuoteMeta(SAVE_PRODUCT)).WillReturnError(Err)

		repository := NewRepository(db)
		ctx := context.TODO()

		id, err := repository.Save(ctx, product_test)

		assert.EqualError(t, err, Err.Error())
		assert.Empty(t, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetAllOK(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id", "description", "expiration_rate", "freezing_rate", "height", "length", "netweight", "product_code", "recommended_freezing_temperature", "width", "product_type_id", "seller_id"}
	rows := sqlmock.NewRows(columns)
	products := []domain.Product{{
		ID: 1,
		Description: "producto congelado",
		ExpirationRate:2,
		FreezingRate:3,
		Height:20.1,
		Length:30.2,
		Netweight:15.2,
		ProductCode:"j3l4k5",
		RecomFreezTemp:20.0,
		Width:30.6,
		ProductTypeID:2,
		SellerID:1,
		},
	}

	for _, product := range products {
		rows.AddRow(product.ID, product.Description, product.ExpirationRate, product.FreezingRate, product.Height, product.Length, product.Netweight, product.ProductCode, product.RecomFreezTemp, product.Width, product.ProductTypeID, product.SellerID)
	}
	
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM products")).WillReturnRows(rows)

	repository := NewRepository(db)
	resultProducts, err := repository.GetAll(context.TODO())

	assert.NoError(t, err)
	assert.Equal(t, products, resultProducts)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAllConflict(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM products")).WillReturnError(sql.ErrConnDone)

	repository := NewRepository(db)
	result, err := repository.GetAll(c)

	assert.Equal(t, sql.ErrConnDone, err)
	assert.Empty(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetOneOk(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "description", "expiration_rate", "freezing_rate", "height", "length", "netweight", "product_code", "recommended_freezing_temperature", "width", "product_type_id", "seller_id"})
	product := domain.Product{
		ID: 1,
		Description: "producto congelado",
		ExpirationRate:2,
		FreezingRate:3,
		Height:20.1,
		Length:30.2,
		Netweight:15.2,
		ProductCode:"j3l4k5",
		RecomFreezTemp:20.0,
		Width:30.6,
		ProductTypeID:2,
		SellerID:1,
	}
	rows.AddRow(product.ID, product.Description, product.ExpirationRate, product.FreezingRate, product.Height, product.Length, product.Netweight, product.ProductCode, product.RecomFreezTemp, product.Width, product.ProductTypeID, product.SellerID)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM products WHERE id=?")).WithArgs(product.ID).WillReturnRows(rows)

	repository := NewRepository(db)
	result, err := repository.Get(c, product.ID)
	assert.Equal(t, product, result)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteOK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	id := 1

	mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM products WHERE id=?"))
	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM products WHERE id=?")).WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))

	repository := NewRepository(db)
	err = repository.Delete(c, id)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM products WHERE id=?")).WillReturnError(sql.ErrNoRows)
	_, err = repository.Get(c, id)
	assert.ErrorContains(t, sql.ErrNoRows, err.Error())
}

func Test_DeleteFail(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	product_test := domain.Product{
		ID: 1,
		Description: "producto congelado",
		ExpirationRate:2,
		FreezingRate:3,
		Height:20.1,
		Length:30.2,
		Netweight:15.2,
		ProductCode:"j3l4k5",
		RecomFreezTemp:20.0,
		Width:30.6,
		ProductTypeID:2,
		SellerID:1,
	}

	mock.ExpectPrepare(regexp.QuoteMeta(DELETE_PRODUCT)).ExpectExec().WithArgs(product_test.ID).WillReturnError(Err)

	repo := NewRepository(db)

	// act
	err = repo.Delete(context.TODO(), int(product_test.ID))

	// assert
	assert.EqualError(t, err, Err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_UpdateOK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	product_test := domain.Product{
		ID: 1,
		Description: "producto congelado",
		ExpirationRate:2,
		FreezingRate:3,
		Height:20.1,
		Length:30.2,
		Netweight:15.2,
		ProductCode:"j3l4k5",
		RecomFreezTemp:20.0,
		Width:30.6,
		ProductTypeID:2,
		SellerID:1,
	}

	product := product_test

	mock.ExpectPrepare(regexp.QuoteMeta(UPDATE_PRODUCT)).
		ExpectExec().WithArgs(product.Description, product.ExpirationRate, product.FreezingRate, product.Height, product.Length, product.Netweight, product.ProductCode, product.RecomFreezTemp, product.Width, product.ProductTypeID, product.SellerID, product.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewRepository(db)
	err = repo.Update(context.TODO(), product)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryGetWithTimeout(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	productId := 1
	columns := []string{"id", "description", "expiration_rate", "freezing_rate", "height", "length", "netweight", "product_code", "recommended_freezing_temperature", "width", "product_type_id", "seller_id"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(productId, "producto congelado", 2, 3, 20.1, 30.2, 15.2, "j3l4k5", 20.0, 30.6, 2, 1)
	mock.ExpectQuery("select id, description, expiration_rate, freezing_rate, height, length, netweight, product_code, recommended_freezing_temperature, width, product_type_id, seller_id").WillDelayFor(10 * time.Second).WillReturnRows(rows)
	repository := NewRepository(db)
	c, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = repository.Get(c, productId)

	assert.Error(t, err)
}

func Test_Repository_Save_Prepare_Fail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mock.ExpectPrepare("INSERT INTO products").WillReturnError(errors.New("insert prepare error"))
	repository := NewRepository(db)
	product := domain.Product{
		ID: 1,
		Description: "producto congelado",
		ExpirationRate:2,
		FreezingRate:3,
		Height:20.1,
		Length:30.2,
		Netweight:15.2,
		ProductCode:"j3l4k5",
		RecomFreezTemp:20.0,
		Width:30.6,
		ProductTypeID:2,
		SellerID:1,
	}
	id, err := repository.Save(context.Background(), product)
	assert.Error(t, err)
	assert.Equal(t, id, 0)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_Repository_Save_Exec_Fail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mock.ExpectPrepare("INSERT INTO products")
	mock.ExpectExec("INSERT INTO products").WillReturnError(errors.New("insert exec error"))
	repository := NewRepository(db)
	product := domain.Product{
		ID: 1,
		Description: "producto congelado",
		ExpirationRate:2,
		FreezingRate:3,
		Height:20.1,
		Length:30.2,
		Netweight:15.2,
		ProductCode:"j3l4k5",
		RecomFreezTemp:20.0,
		Width:30.6,
		ProductTypeID:2,
		SellerID:1,
	}
	id, err := repository.Save(context.Background(), product)
	assert.Error(t, err)
	assert.Equal(t, id, 0)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_DeleteFailRowsAffected(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	product_test := domain.Product{
		ID: 1,
		Description: "producto congelado",
		ExpirationRate:2,
		FreezingRate:3,
		Height:20.1,
		Length:30.2,
		Netweight:15.2,
		ProductCode:"j3l4k5",
		RecomFreezTemp:20.0,
		Width:30.6,
		ProductTypeID:2,
		SellerID:1,
	}

	mock.ExpectPrepare(regexp.QuoteMeta(DELETE_PRODUCT)).ExpectExec().WithArgs(product_test.ID).WillReturnResult(sqlmock.NewResult(1, 2))

	repo := NewRepository(db)

	// act
	err = repo.Delete(context.TODO(), int(product_test.ID))

	// assert
	assert.Nil(t, err, nil)
}

func Test_UpdateFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	product_test := domain.Product{
		ID: 1,
		Description: "producto congelado",
		ExpirationRate:2,
		FreezingRate:3,
		Height:20.1,
		Length:30.2,
		Netweight:15.2,
		ProductCode:"j3l4k5",
		RecomFreezTemp:20.0,
		Width:30.6,
		ProductTypeID:2,
		SellerID:1,
	}
	product := product_test
	mock.ExpectPrepare(regexp.QuoteMeta(UPDATE_PRODUCT)).ExpectExec().WithArgs(product.Description, product.ExpirationRate, product.FreezingRate, product.Height, product.Length, product.Netweight, product.ProductCode, product.RecomFreezTemp, product.Width, product.ProductTypeID, product.SellerID, product.ID).WillReturnError(Err)
		
	repo := NewRepository(db)
	err = repo.Update(context.TODO(), product_test)

	assert.EqualError(t, err, Err.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}