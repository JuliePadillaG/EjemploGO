package seller

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

func TestExist(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := domain.Seller{
		ID:          1,
		CID:         19,
		CompanyName: "LG",
		Address:     "Avenida 11122",
		Telephone:   "0303456",
	}

	t.Run("exist", func(t *testing.T) {
		columns := []string{"cid"}
		rows := sqlmock.NewRows(columns)

		rows.AddRow(s.CID)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT cid FROM seller WHERE cid=?;")).WithArgs(1).WillReturnRows(rows)

		repo := NewRepository(db)

		resp := repo.Exists(context.TODO(), 1)

		assert.True(t, resp)
	})

	t.Run("exist locality", func(t *testing.T) {
		columns := []string{"id"}
		rows := sqlmock.NewRows(columns)

		rows.AddRow(s.LocalityID)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM locality WHERE id=?;")).WithArgs(s.LocalityID).WillReturnRows(rows)

		repo := NewRepository(db)

		resp := repo.LocalityExists(context.TODO(), s.LocalityID)

		assert.True(t, resp)
	})
}

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := domain.Seller{
		ID:          1,
		CID:         19,
		CompanyName: "LG",
		Address:     "Avenida 11122",
		Telephone:   "0303456",
	}

	t.Run("create ok", func(t *testing.T) {

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO seller (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO seller (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)")).WillReturnResult(sqlmock.NewResult(1, 1))

		repository := NewRepository(db)
		ctx := context.TODO()

		newID, err := repository.Save(ctx, s)
		assert.NoError(t, err)

		assert.NotNil(t, newID)
		assert.Equal(t, 1, newID)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("create fail", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO seller (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO seller (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)")).WillReturnError(ErrForzado)

		repository := NewRepository(db)
		ctx := context.TODO()

		id, err := repository.Save(ctx, s)

		assert.EqualError(t, err, ErrForzado.Error())
		assert.Empty(t, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("create fail id exists", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)

		defer db.Close()

		columns := []string{"id", "cid", "company_name", "address", "telephone", "locality_id"}
		rows := sqlmock.NewRows(columns)
		sellers := []domain.Seller{s}

		for _, element := range sellers {
			rows.AddRow(element.ID, element.CID, element.CompanyName, element.Address, element.Telephone, element.LocalityID)
		}

		mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO seller (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"))
		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO seller (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)")).WithArgs(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID).WillReturnError(ErrCidExists)

		repository := NewRepository(db)
		ctx := context.TODO()

		id, err := repository.Save(ctx, s)

		assert.EqualError(t, err, ErrCidExists.Error())
		assert.Empty(t, id)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("get ok", func(t *testing.T) {

		columns := []string{"id", "cid", "company_name", "address", "telephone", "locality_id"}
		rows := sqlmock.NewRows(columns)
		s := domain.Seller{
			ID:          1,
			CID:         19,
			CompanyName: "LG",
			Address:     "Avenida 11122",
			Telephone:   "0303456",
		}

		rows.AddRow(s.ID, s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM seller WHERE id=?;")).WithArgs(1).WillReturnRows(rows)

		repo := NewRepository(db)
		result, err := repo.Get(context.TODO(), 1)

		assert.NoError(t, err)
		assert.Equal(t, s, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("get all ok", func(t *testing.T) {

		columns := []string{"id", "cid", "company_name", "address", "telephone", "locality_id"}
		rows := sqlmock.NewRows(columns)
		sellers := []domain.Seller{
			{
				ID:          1,
				CID:         19,
				CompanyName: "LG",
				Address:     "Avenida 11122",
				Telephone:   "0303456",
			},
		}

		for _, s := range sellers {
			rows.AddRow(s.ID, s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID)
		}

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM seller")).WillReturnRows(rows)

		repo := NewRepository(db)
		result, err := repo.GetAll(context.TODO())

		assert.NoError(t, err)
		assert.Equal(t, sellers, result)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := domain.Seller{
		ID:          1,
		CID:         19,
		CompanyName: "LG",
		Address:     "Avenida 11122",
		Telephone:   "0303456",
		LocalityID:  1759,
	}

	t.Run("delete ok", func(t *testing.T) {

		columns := []string{"id", "cid", "company_name", "address", "telephone", "locality_id"}
		rows := sqlmock.NewRows(columns)

		rows.AddRow(s.ID, s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID)

		mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM seller WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("DELETE FROM seller WHERE id=?")).WithArgs(s.ID).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewRepository(db)
		err := repo.Delete(context.TODO(), s.ID)

		assert.NoError(t, err)
	})
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	s := domain.Seller{
		ID:          1,
		CID:         19,
		CompanyName: "LG",
		Address:     "Avenida 11122",
		Telephone:   "0303456",
		LocalityID:  1759,
	}

	t.Run("update ok", func(t *testing.T) {

		columns := []string{"id", "cid", "company_name", "address", "telephone", "locality_id"}
		rows := sqlmock.NewRows(columns)

		rows.AddRow(s.ID, s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID)

		mock.ExpectPrepare(regexp.QuoteMeta("UPDATE seller SET cid=?, company_name=?, address=?, telephone=?, locality_id=? WHERE id=?"))
		mock.ExpectExec(regexp.QuoteMeta("UPDATE seller SET cid=?, company_name=?, address=?, telephone=?, locality_id=? WHERE id=?")).WithArgs(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID, s.ID).WillReturnResult(sqlmock.NewResult(1, 1))

		repo := NewRepository(db)
		err := repo.Update(context.TODO(), s)

		assert.NoError(t, err)
	})
}
