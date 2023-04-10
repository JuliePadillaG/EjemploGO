package inboundorder

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_Create_Ok(t *testing.T) {
	//Arrange
	bo := domain.Inbound_order{
		ID:               1,
		Order_date:       "2021-04-04",
		Order_number:     "order#1",
		Employee_id:      4,
		Product_batch_id: 1,
		Warehouse_id:     1,
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare(regexp.QuoteMeta(SAVE))
	mock.ExpectExec(regexp.QuoteMeta(SAVE)).
		WithArgs(bo.Order_date, bo.Order_number, bo.Employee_id, bo.Product_batch_id, bo.Warehouse_id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repository := NewRepository(db)

	//Act
	id, err := repository.Save(context.TODO(), bo)

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, bo.ID, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_Create_Fail(t *testing.T) {
	//Arrange
	bo := domain.Inbound_order{
		ID:               1,
		Order_date:       "2021-04-04",
		Order_number:     "order#1",
		Employee_id:      4,
		Product_batch_id: 1,
		Warehouse_id:     1,
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare(regexp.QuoteMeta(SAVE))
	mock.ExpectExec(regexp.QuoteMeta(SAVE)).
		WithArgs(bo.Order_date, bo.Order_number, bo.Employee_id, bo.Product_batch_id, bo.Warehouse_id).
		WillReturnError(errors.New("INSERT ERROR"))

	repository := NewRepository(db)

	//Act
	id, err := repository.Save(context.TODO(), bo)

	t.Log(err)
	//Assert
	assert.Error(t, err)
	assert.Equal(t, 0, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_GetAll_Ok(t *testing.T) {
	//Arrange
	data := []domain.Inbound_order{
		{
			ID:               1,
			Order_date:       "2021-04-04",
			Order_number:     "order#1",
			Employee_id:      4,
			Product_batch_id: 1,
			Warehouse_id:     1,
		},
		{
			ID:               2,
			Order_date:       "2021-04-06",
			Order_number:     "order#2",
			Employee_id:      4,
			Product_batch_id: 1,
			Warehouse_id:     1,
		},
		{
			ID:               3,
			Order_date:       "2021-04-07",
			Order_number:     "order#3",
			Employee_id:      4,
			Product_batch_id: 1,
			Warehouse_id:     1,
		},
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	row := sqlmock.NewRows([]string{"id", "order_date", "order_number", "employee_id", "product_batch_id", "warehouse_id"})
	for _, bo := range data {
		row.AddRow(bo.ID, bo.Order_date, bo.Order_number, bo.Employee_id, bo.Product_batch_id, bo.Warehouse_id)
	}
	mock.ExpectQuery(regexp.QuoteMeta(GET_ALL)).WillReturnRows(row)
	repository := NewRepository(db)

	//Act
	result, err := repository.GetAll(context.TODO())

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, data, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_GetAll_Fail(t *testing.T) {
	//Arrange
	data := []domain.Inbound_order{
		{
			ID:               1,
			Order_date:       "2021-04-04",
			Order_number:     "order#1",
			Employee_id:      4,
			Product_batch_id: 1,
			Warehouse_id:     1,
		},
		{
			ID:               2,
			Order_date:       "2021-04-06",
			Order_number:     "order#2",
			Employee_id:      4,
			Product_batch_id: 1,
			Warehouse_id:     1,
		},
		{
			ID:               3,
			Order_date:       "2021-04-07",
			Order_number:     "order#3",
			Employee_id:      4,
			Product_batch_id: 1,
			Warehouse_id:     1,
		},
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	row := sqlmock.NewRows([]string{"id", "order_date", "order_number", "employee_id", "product_batch_id", "warehouse_id"})
	for _, bo := range data {
		row.AddRow(bo.ID, bo.Order_date, bo.Order_number, bo.Employee_id, bo.Product_batch_id, bo.Warehouse_id)
	}
	mock.ExpectQuery(regexp.QuoteMeta(GET_ALL)).WillReturnError(errors.New("GET ALL ERROR"))
	repository := NewRepository(db)

	//Act
	result, err := repository.GetAll(context.TODO())

	//Assert
	assert.NotNil(t, err)
	assert.Empty(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_ExistEmployee_Ok(t *testing.T) {
	//Arrange
	id := 1
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id"})
	row.AddRow(id)
	mock.ExpectQuery(regexp.QuoteMeta(EXIST_EMPLOYEE)).WithArgs(id).WillReturnRows(row)
	repository := NewRepository(db)

	//Act
	result := repository.ExistsEmployee(context.TODO(), id)
	//Assert
	assert.True(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_ExistEmployee_Fail(t *testing.T) {
	//Arrange
	id := 3
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id"})
	mock.ExpectQuery(regexp.QuoteMeta(EXIST_EMPLOYEE)).WithArgs(id).WillReturnRows(row)
	repository := NewRepository(db)

	//Act
	result := repository.ExistsEmployee(context.TODO(), id)
	//Assert
	assert.False(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_ExistsInboundOrder_Ok(t *testing.T) {
	//Arrange
	ordernumber := "order#1"
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"order_number"})
	row.AddRow(ordernumber)
	mock.ExpectQuery(regexp.QuoteMeta(EXIST_INBOUND)).WithArgs(ordernumber).WillReturnRows(row)
	repository := NewRepository(db)

	//Act
	result := repository.ExistsInboundOrder(context.TODO(), ordernumber)
	//Assert
	assert.True(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_ExistsInboundOrder_Fail(t *testing.T) {
	//Arrange
	ordernumber := "order#1"
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"order_number"})
	mock.ExpectQuery(regexp.QuoteMeta(EXIST_INBOUND)).WithArgs(ordernumber).WillReturnRows(row)
	repository := NewRepository(db)

	//Act
	result := repository.ExistsInboundOrder(context.TODO(), ordernumber)
	//Assert
	assert.False(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}
