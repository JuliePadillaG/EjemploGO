package purchase_orders

import (
	"context"
	"fmt"
	"time"

	// "log"
	"errors"
	"regexp"
	"testing"

	// "database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/suite"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

var (
	ErrForzadoPurchaseOrders = errors.New("Error forzado")
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

func (s *createDataBaseRepoSuite) Test_SavePurchaseOrdersOK() {

        // Arrange
        idExpected := 1
        orderDate := time.Now()
        var testPurchaseOrders = domain.PurchaseOrders{
                OrderNumber: "232345",
                OrderDate: &orderDate,
                TrackingCode: "bar",
                BuyerID: 1,
                ProductRecordID: 1,
                OrderStatusID: 1,
        }

        stmt := regexp.QuoteMeta(SAVE_BUYER)
        params := testPurchaseOrders
	s.sqlMock.ExpectPrepare(stmt).WillReturnError(nil)
	s.sqlMock.ExpectExec(stmt).
                WithArgs(params.OrderNumber, params.OrderDate, params.TrackingCode, params.BuyerID, params.ProductRecordID, params.OrderStatusID).
                WillReturnResult(sqlmock.NewResult(1, 1))

	// Act
	id, err := s.dbRepository.Save(s.context, testPurchaseOrders)

	// Assert
        s.Nil(err)
	s.Equal(idExpected, id)
	s.NoError(err)
}


func (s *createDataBaseRepoSuite) Test_SavePurchaseOrdersFail() {

        // Arrange
        var testpurchaseorders = domain.PurchaseOrders{}
        stmt := regexp.QuoteMeta(SAVE_BUYER)
        params := testpurchaseorders
	s.sqlMock.ExpectPrepare(stmt).WillReturnError(nil)
	s.sqlMock.ExpectExec(stmt).
                WithArgs(params.OrderNumber, params.OrderDate, params.TrackingCode, params.BuyerID, params.ProductRecordID, params.OrderStatusID).
                WillReturnError(ErrForzadoPurchaseOrders)

	// Act
	id, err := s.dbRepository.Save(s.context, testpurchaseorders)

	// Assert
	s.EqualError(err, ErrForzadoPurchaseOrders.Error())
	s.Empty(id)
	s.NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *createDataBaseRepoSuite) Test_ExistBuyersID() {

        params := 1
        stmt := regexp.QuoteMeta(EXISTS_BUYER_ID)
        columns := []string{"id"}
        rows := s.sqlMock.NewRows(columns)
        rows.AddRow(1)
	s.sqlMock.ExpectQuery(stmt).WithArgs(params).WillReturnRows(rows)

        // Act
        exist := s.dbRepository.ExistsBuyersID(s.context, 1)

        // Arrange
	s.True(exist)
        s.NoError(s.sqlMock.ExpectationsWereMet())
}


func (s *createDataBaseRepoSuite) Test_ExistProductRecordID() {

        params := 1
        stmt := regexp.QuoteMeta(EXISTS_PRODUCT_RECORD_ID)
        columns := []string{"id"}
        rows := s.sqlMock.NewRows(columns)
        rows.AddRow(1)
	s.sqlMock.ExpectQuery(stmt).WithArgs(params).WillReturnRows(rows)

        // Act
        exist := s.dbRepository.ExistsProductRecordsID(s.context, 1)

        // Arrange
	s.True(exist)
        s.NoError(s.sqlMock.ExpectationsWereMet())
}


func (s *createDataBaseRepoSuite) Test_GetPurchaseOrdersWithQueryStringOK() {

        // Arrange
        reportPurchaseOrders := []domain.ReportPurchaseOrders{
                {
                        ID: 1,
                        CardNumberID: "402323",
		        FirstName: "Peter",
                        LastName: "Peter",
                        PurchaseOrdersCount: 3,
                },
        }

        columns := []string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}
        rows := s.sqlMock.NewRows(columns)

        for _, report := range reportPurchaseOrders {
                rows.AddRow(report.ID, report.CardNumberID, report.FirstName, report.LastName, report.PurchaseOrdersCount)
        }

        stmt := regexp.QuoteMeta(GET_REPORT_PURCHASEORDERS_BY_BUYERID)
        s.sqlMock.ExpectQuery(stmt).
                WithArgs(1).
                WillReturnRows(rows)

	// Act
	reports, err := s.dbRepository.Get(s.context, 1)

	// Assert
        s.Nil(err)
	s.Equal(len(reportPurchaseOrders), len(reports))
        s.NoError(s.sqlMock.ExpectationsWereMet())
}

func (s *createDataBaseRepoSuite) Test_GetPurchaseOrdersWithoutQueryStringOK() {

        // Arrange
        reportPurchaseOrders := []domain.ReportPurchaseOrders{
                {
                        ID: 1,
                        CardNumberID: "402323",
		        FirstName: "Peter",
                        LastName: "Peter",
                        PurchaseOrdersCount: 3,
                },
                {
                        ID: 2,
                        CardNumberID: "402324",
		        FirstName: "Jhon",
                        LastName: "Jhon",
                        PurchaseOrdersCount: 3,
                },
        }

        columns := []string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}
        rows := s.sqlMock.NewRows(columns)

        for _, report := range reportPurchaseOrders {
                rows.AddRow(report.ID, report.CardNumberID, report.FirstName, report.LastName, report.PurchaseOrdersCount)
        }

        stmt := regexp.QuoteMeta(GET_REPORT_PURCHASEORDERS)
        s.sqlMock.ExpectQuery(stmt).WillReturnRows(rows)

	// Act
	reports, err := s.dbRepository.Get(s.context, 0)

	// Assert
        s.Nil(err)
	s.Equal(len(reportPurchaseOrders), len(reports))
        s.NoError(s.sqlMock.ExpectationsWereMet())
}


func (s *createDataBaseRepoSuite) Test_GetPurchaseOrdersWithQueryStringFail() {

        // Arrange
        reportPurchaseOrders := []domain.ReportPurchaseOrders{
                {
                        ID: 1,
                        CardNumberID: "402323",
		        FirstName: "Peter",
                        LastName: "Peter",
                        PurchaseOrdersCount: 3,
                },
        }

        columns := []string{"id", "card_number_id", "first_name", "last_name", "purchase_orders_count"}
        rows := s.sqlMock.NewRows(columns)

        for _, report := range reportPurchaseOrders {
                rows.AddRow(report.ID, report.CardNumberID, report.FirstName, report.LastName, report.PurchaseOrdersCount)
        }

        stmt := regexp.QuoteMeta(GET_REPORT_PURCHASEORDERS_BY_BUYERID)
        s.sqlMock.ExpectQuery(stmt).
                WithArgs(1).
                WillReturnError(ErrForzadoPurchaseOrders)

	// Act
	reports, err := s.dbRepository.Get(s.context, 1)

	// Assert
	s.EqualError(err, ErrForzadoPurchaseOrders.Error())
	s.Empty(reports)
	s.NoError(s.sqlMock.ExpectationsWereMet())
}
