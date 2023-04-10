package locality

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/stretchr/testify/assert"
)

var (
	ErrForzado = errors.New("Error forzado")
)

func TestExistLocalityOK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	locality_test := domain.Locality{
		ID:           1759,
		LocalityName: "Gonzalez Catan",
		ProvinceName: "Buenos Aires",
		CountryName:  "Argentina",
	}

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	rows.AddRow(locality_test.ID)
	mock.ExpectQuery(regexp.QuoteMeta(EXIST_LOCALITY)).WithArgs(1).WillReturnRows(rows)

	repo := NewRepository(db)

	resp := repo.Exists(context.TODO(), 1)

	assert.True(t, resp)
}

func TestExistLocalityFail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	mock.ExpectQuery(regexp.QuoteMeta(EXIST_LOCALITY)).WithArgs(2).WillReturnRows(rows)

	repo := NewRepository(db)

	resp := repo.Exists(context.TODO(), 2)

	assert.False(t, resp)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestCreateLocality(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	locality_test := domain.Locality{
		ID:           1759,
		LocalityName: "Gonzalez Catan",
		ProvinceName: "Buenos Aires",
		CountryName:  "Argentina",
	}

	t.Run("create ok", func(t *testing.T) {

		mock.ExpectPrepare(regexp.QuoteMeta(CREATE_LOCALITY))
		mock.ExpectExec(regexp.QuoteMeta(CREATE_LOCALITY)).WillReturnResult(sqlmock.NewResult(1, 1))

		repository := NewRepository(db)
		ctx := context.TODO()

		newID, err := repository.Create(ctx, locality_test)
		assert.NoError(t, err)

		assert.NotNil(t, newID)
		assert.Equal(t, 1, newID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("create fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta(CREATE_LOCALITY))
		mock.ExpectExec(regexp.QuoteMeta(CREATE_LOCALITY)).WillReturnError(ErrForzado)

		repository := NewRepository(db)
		ctx := context.TODO()

		id, err := repository.Create(ctx, locality_test)

		assert.EqualError(t, err, ErrForzado.Error())
		assert.Empty(t, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("create fail id exists", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		columns := []string{"id", "locality_name", "province_name", "country_name"}
		rows := sqlmock.NewRows(columns)
		localities := []domain.Locality{locality_test}

		for _, locality := range localities {
			rows.AddRow(locality.ID, locality.LocalityName, locality.ProvinceName, locality.CountryName)
		}

		mock.ExpectPrepare(regexp.QuoteMeta(CREATE_LOCALITY))
		mock.ExpectExec(regexp.QuoteMeta(CREATE_LOCALITY)).WithArgs(locality_test.ID, locality_test.LocalityName, locality_test.ProvinceName, locality_test.CountryName).WillReturnError(ErrExists)

		repository := NewRepository(db)
		ctx := context.TODO()

		id, err := repository.Create(ctx, locality_test)

		assert.EqualError(t, err, ErrExists.Error())
		assert.Empty(t, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetSellers(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("get ok", func(t *testing.T) {

		columns := []string{"locality_id", "locality_name", "sellers_count"}
		rows := sqlmock.NewRows(columns)
		localities := []domain.ResponseLocality{{ID: 1759, LocalityName: "Gonzalez Catan", SellersCount: 1}}

		for _, l := range localities {
			rows.AddRow(l.ID, l.LocalityName, l.SellersCount)
		}

		mock.ExpectQuery(regexp.QuoteMeta(GET_SELLERS)).WillReturnRows(rows)

		repo := NewRepository(db)
		result, err := repo.GetAllSellersByLocality(context.TODO(), "")

		assert.NoError(t, err)
		assert.Equal(t, localities, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("get fail with id", func(t *testing.T) {

		columns := []string{"locality_id", "locality_name", "sellers_count"}
		rows := sqlmock.NewRows(columns)
		localities := []domain.ResponseLocality{{ID: 1759, LocalityName: "Gonzalez Catan", SellersCount: 1}}

		for _, l := range localities {
			rows.AddRow(l.ID, l.LocalityName, l.SellersCount)
		}

		mock.ExpectQuery(regexp.QuoteMeta(GET_SELLERS_BY_ID)).WithArgs("1759")

		repo := NewRepository(db)
		result, err := repo.GetAllSellersByLocality(context.TODO(), "1759")

		assert.EqualError(t, err, "locality_id not found")
		assert.Empty(t, result)
	})

}

func TestGetCarries(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"locality_id", "locality_name", "carries_count"}
	rows := sqlmock.NewRows(columns)
	carries := []domain.CarriesReport{{LocalityID: 1759, LocalityName: "Gonzalez Catan", CarriesCount: 1}}

	for _, l := range carries {
		rows.AddRow(l.LocalityID, l.LocalityName, l.CarriesCount)
	}

	t.Run("get ok", func(t *testing.T) {

		mock.ExpectQuery(regexp.QuoteMeta("SELECT locality.id, locality.locality_name, COUNT(*) AS carries_count FROM carries RIGHT JOIN locality on carries.locality_id = locality.id GROUP BY id;")).WillReturnRows(rows)

		repo := NewRepository(db)
		result, err := repo.GetCarriesReport(context.TODO(), "")

		assert.NoError(t, err)
		assert.Equal(t, carries, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("get fail with id", func(t *testing.T) {

		mock.ExpectQuery(regexp.QuoteMeta("SELECT locality.id, locality.locality_name, COUNT(*) AS carries_count FROM carries right join locality on carries.locality_id = locality.id WHERE locality.id = ? GROUP BY id;")).WithArgs(1759)

		repo := NewRepository(db)
		result, err := repo.GetCarriesReport(context.TODO(), "1759")

		assert.EqualError(t, err, "id does not exist")
		assert.Empty(t, result)
	})

}
