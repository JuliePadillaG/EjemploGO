package buyer

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

var (
	ErrForzadoBuyer = errors.New("Error forzado")
	ErrForzadoScanBuyer = errors.New("sql: Scan error on column index 0, name \"id\": converting NULL to int is unsupported")
)

type createDataBaseRepoSuite struct {
	suite.Suite
	context      context.Context
	sqlMock      sqlmock.Sqlmock
	dbRepository Repository
}

func (s *createDataBaseRepoSuite) SetupTest() {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic(fmt.Errorf("error initializing sql mock %v", err))
	}

        // defer db.Close()

        s.context = context.TODO()
	s.sqlMock = mock
	s.dbRepository = NewRepository(db)
}

func TestCreateDataBaseSuite(t *testing.T) {
	suite.Run(t, &createDataBaseRepoSuite{})
}

func (s *createDataBaseRepoSuite) Test_SaveBuyerOK() {

        // Arrange
        idExpected := 1
        var testBuyer = domain.Buyer{
                CardNumberID: "232345",
                FirstName:    "foo",
                LastName:     "bar",
        }

        stmt := regexp.QuoteMeta(SAVE_BUYER)
        params := testBuyer
	s.sqlMock.ExpectPrepare(stmt).WillReturnError(nil)
	s.sqlMock.ExpectExec(stmt).
                WithArgs(params.CardNumberID, params.FirstName, params.LastName).
                WillReturnResult(sqlmock.NewResult(1, 1))

	// Act
	id, err := s.dbRepository.Save(s.context, testBuyer)

	// Assert
        s.Nil(err)
	s.Equal(idExpected, id)
	s.NoError(err)
}

func (s *createDataBaseRepoSuite) Test_SaveBuyerFail() {

        // Arrange
        var testBuyer = domain.Buyer{}
        stmt := regexp.QuoteMeta(SAVE_BUYER)
        params := testBuyer
	s.sqlMock.ExpectPrepare(stmt).WillReturnError(nil)
	s.sqlMock.ExpectExec(stmt).
                WithArgs(params.CardNumberID, params.FirstName, params.LastName).
                WillReturnError(ErrForzadoBuyer)

	// Act
	id, err := s.dbRepository.Save(s.context, testBuyer)

	// Assert
	s.EqualError(err, ErrForzadoBuyer.Error())
	s.Empty(id)
	s.NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *createDataBaseRepoSuite) Test_GetBuyerOK() {

        // Arrange
	rows := s.sqlMock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"})
        var testBuyer = domain.Buyer{
                ID:           1,
                CardNumberID: "232345",
                FirstName:    "foo",
                LastName:     "bar",
        }

        stmt := regexp.QuoteMeta(GET_BUYER_BY_ID)
	rows.AddRow(testBuyer.ID, testBuyer.CardNumberID, testBuyer.FirstName, testBuyer.LastName)
	s.sqlMock.ExpectQuery(stmt).WithArgs(testBuyer.ID).WillReturnRows(rows)

        // Act
	result, err := s.dbRepository.Get(s.context, testBuyer.ID)

        // Assert
	s.Equal(testBuyer, result)
	s.NoError(err)
	s.NoError(s.sqlMock.ExpectationsWereMet())
}


func (s *createDataBaseRepoSuite) Test_GetBuyerFail() {

        // Arrange
	rows := s.sqlMock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"})
        var testBuyer = domain.Buyer{
                ID:           1,
                CardNumberID: "232345",
                FirstName:    "foo",
                LastName:     "bar",
        }

        stmt := regexp.QuoteMeta(GET_BUYER_BY_ID)
	rows.AddRow(nil, testBuyer.CardNumberID, testBuyer.FirstName, testBuyer.LastName).
                RowError(2, ErrForzadoScanBuyer)
	s.sqlMock.ExpectQuery(stmt).WithArgs(testBuyer.ID).WillReturnRows(rows)

        // Act
	_, err := s.dbRepository.Get(s.context, testBuyer.ID)

        log.Printf("%+v\n", err)

        // Assert
	s.EqualError(err, ErrForzadoScanBuyer.Error())
}


func (s *createDataBaseRepoSuite) Test_GetAllBuyerOK() {

        // Arrange
	rows := s.sqlMock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"})
        testBuyers := []domain.Buyer{
                {
                        ID:           1,
                        CardNumberID: "232345",
                        FirstName:    "foo",
                        LastName:     "bar",
                },
                {
                        ID:           2,
                        CardNumberID: "232346",
                        FirstName:    "foo1",
                        LastName:     "bar1",
                },
        }

	for _, buyer := range testBuyers {
	        rows.AddRow(buyer.ID, buyer.CardNumberID, buyer.FirstName, buyer.LastName)
	}

        stmt := regexp.QuoteMeta(GET_ALL_BUYERS)
	s.sqlMock.ExpectQuery(stmt).WillReturnRows(rows)

        // Act
	resultBuyers, err := s.dbRepository.GetAll(s.context)

        // Arrange
	s.NoError(err)
	s.Equal(testBuyers, resultBuyers)
	s.NoError(s.sqlMock.ExpectationsWereMet())
}


func (s *createDataBaseRepoSuite) Test_GetAllBuyerFail() {

        // Arrange
	rows := s.sqlMock.NewRows([]string{"id", "card_number_id", "first_name", "last_name"})
        testBuyers := []domain.Buyer{
                {
                        ID:           1,
                        CardNumberID: "232345",
                        FirstName:    "foo",
                        LastName:     "bar",
                },
                {
                        ID:           2,
                        CardNumberID: "232346",
                        FirstName:    "foo1",
                        LastName:     "bar1",
                },
        }

	for _, buyer := range testBuyers {
	        rows.AddRow(buyer.ID, buyer.CardNumberID, buyer.FirstName, buyer.LastName)
	}

        stmt := regexp.QuoteMeta(GET_ALL_BUYERS)
	s.sqlMock.ExpectQuery(stmt).WillReturnError(ErrForzadoBuyer)

        // Act
	resultBuyers, err := s.dbRepository.GetAll(s.context)

        // Arrange
	s.EqualError(err, ErrForzadoBuyer.Error())
	s.Empty(resultBuyers)
	s.NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *createDataBaseRepoSuite) Test_ExistBuyer() {

        // Arrange
        testBuyer := domain.Buyer{
                ID:           1,
                CardNumberID: "232346",
                FirstName:    "foo1",
                LastName:     "bar1",
        }

        params := "232346"
        stmt := regexp.QuoteMeta(EXISTS_BUYER)

        rows := s.sqlMock.NewRows([]string{"card_number_id"})
        rows.AddRow(testBuyer.CardNumberID)
	s.sqlMock.ExpectQuery(stmt).WithArgs(params).WillReturnRows(rows)

        // Act
        exist := s.dbRepository.Exists(s.context, testBuyer.CardNumberID)

        // Arrange
	s.True(exist)
}


func (s *createDataBaseRepoSuite) Test_ExistBuyerFail() {

        // Arrange
        params := "232346"
        stmt := regexp.QuoteMeta(EXISTS_BUYER)

        rows := s.sqlMock.NewRows([]string{"card_number_id"})
	s.sqlMock.ExpectQuery(stmt).WithArgs(params).WillReturnRows(rows)

        // Act
        exist := s.dbRepository.Exists(s.context, params)

        // Arrange
	s.False(exist)
        s.NoError(s.sqlMock.ExpectationsWereMet())
}


func (s *createDataBaseRepoSuite) Test_DeleteBuyerOK() {

        // Arrange
	params := 1
        stmt := regexp.QuoteMeta(DELETE_BUYER)
	s.sqlMock.ExpectPrepare(stmt)
	s.sqlMock.ExpectExec(stmt).WithArgs(params).WillReturnResult(sqlmock.NewResult(1, 1))

        // Act
        err := s.dbRepository.Delete(s.context, params)

        // Assert
	s.NoError(err)
	s.NoError(s.sqlMock.ExpectationsWereMet())

        // -----------------------------------

        // Arrange
        stmt = regexp.QuoteMeta(GET_BUYER_BY_ID)
	s.sqlMock.ExpectQuery(stmt).WillReturnError(sql.ErrNoRows)

        // Act
	_, err = s.dbRepository.Get(s.context, params)

        // Assert
	s.ErrorContains(sql.ErrNoRows, err.Error())
}


func (s *createDataBaseRepoSuite) Test_DeleteBuyerFail() {

        // Arrange
        testBuyer := domain.Buyer{
                ID:           1,
                CardNumberID: "232346",
                FirstName:    "foo1",
                LastName:     "bar1",
        }

        stmt := regexp.QuoteMeta(DELETE_BUYER)
	s.sqlMock.ExpectPrepare(stmt).ExpectExec().WithArgs(testBuyer.ID).WillReturnError(ErrForzadoBuyer)

        // Act
        err := s.dbRepository.Delete(s.context, testBuyer.ID)

        // Assert
	s.EqualError(err, ErrForzadoBuyer.Error())
	s.NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *createDataBaseRepoSuite) Test_UpdateBuyerOK() {

        // Arrange
        testBuyer := domain.Buyer{
                ID:           1,
                CardNumberID: "232346",
                FirstName:    "foo1",
                LastName:     "bar1",
        }

        stmt := regexp.QuoteMeta(UPDATE_BUYER)
	s.sqlMock.ExpectPrepare(stmt).
		ExpectExec().
                WithArgs(testBuyer.FirstName, testBuyer.LastName, testBuyer.ID).
                WillReturnResult(sqlmock.NewResult(0, 1))

        // Act
        err := s.dbRepository.Update(s.context, testBuyer)

        // Assert
	s.NoError(err)
	s.NoError(s.sqlMock.ExpectationsWereMet())
}


func (s *createDataBaseRepoSuite) Test_UpdateBuyerFail() {

        // Arrange
        testBuyer := domain.Buyer{
                ID:           1,
                CardNumberID: "232346",
                FirstName:    "foo1",
                LastName:     "bar1",
        }

        stmt := regexp.QuoteMeta(UPDATE_BUYER)
	s.sqlMock.ExpectPrepare(stmt).
		ExpectExec().
                WithArgs(testBuyer.FirstName, testBuyer.LastName, testBuyer.ID).
                WillReturnError(ErrForzadoBuyer)

        // Act
        err := s.dbRepository.Update(s.context, testBuyer)

        // Assert
	s.Error(err)
	s.NoError(s.sqlMock.ExpectationsWereMet())
}
