package carry

import (
	"context"
	"errors"
	"testing"

	"regexp"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_Repository_Store_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mock.ExpectPrepare("INSERT INTO carries")
	mock.ExpectExec("INSERT INTO carries").WillReturnResult(sqlmock.NewResult(1, 1))
	repository := NewRepository(db)
	carry := domain.Carry{
		CID:          "1",
		Company_name: "test",
		Address:      "test",
		Telephone:    "test",
		Locality_id:  1,
	}
	id, err := repository.Save(context.Background(), carry)
	assert.NoError(t, err)
	assert.Equal(t, id, 1)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func Test_Repository_Store_Prepare_Error_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mock.ExpectPrepare("INSERT INTO carries").WillReturnError(errors.New("insert prepare error"))
	repository := NewRepository(db)
	carry := domain.Carry{
		CID:          "1",
		Company_name: "test",
		Address:      "test",
		Telephone:    "test",
		Locality_id:  1,
	}
	id, err := repository.Save(context.Background(), carry)
	assert.Error(t, err)
	assert.Equal(t, id, 0)
	assert.NoError(t, mock.ExpectationsWereMet())

}

func Test_Repository_Store_Exec_Error_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	mock.ExpectPrepare("INSERT INTO carries")
	mock.ExpectExec("INSERT INTO carries").WillReturnError(errors.New("insert exec error"))
	repository := NewRepository(db)
	carry := domain.Carry{
		CID:          "1",
		Company_name: "test",
		Address:      "test",
		Telephone:    "test",
		Locality_id:  1,
	}
	id, err := repository.Save(context.Background(), carry)
	assert.Error(t, err)
	assert.Equal(t, id, 0)
	assert.NoError(t, mock.ExpectationsWereMet())

}
func Test_Repository_Exists_Carries_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	columns := []string{"cid"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow("1")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT cid FROM carries WHERE cid=?;")).WithArgs("1").WillReturnRows(rows)
	repository := NewRepository(db)
	ctx := context.TODO()

	exists := repository.Exists(ctx, "1")
	assert.True(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_Repository_Exists_Locality_Mock(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(1)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM locality WHERE id=?;")).WithArgs(1).WillReturnRows(rows)
	repository := NewRepository(db)
	ctx := context.TODO()

	exists := repository.ExistsLocality(ctx, 1)
	assert.True(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())
}
