package employee

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllEmployee_Ok(t *testing.T) {
	//Arrange
	data := []domain.Employee{
		{
			ID:           1,
			CardNumberID: "402323",
			FirstName:    "Jhon",
			LastName:     "Doe",
			WarehouseID:  1,
		},
		{
			ID:           2,
			CardNumberID: "402322",
			FirstName:    "Jhon",
			LastName:     "Doe",
			WarehouseID:  3,
		},
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"})

	for _, e := range data {
		rows.AddRow(e.ID, e.CardNumberID, e.FirstName, e.LastName, e.WarehouseID)
	}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees")).WillReturnRows(rows)
	repository := NewRepository(db)

	//Act
	employees, err := repository.GetAll(context.TODO())

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, data, employees)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_GetAllEmployee_Fail(t *testing.T) {
	//Arrange
	data := []domain.Employee{
		{
			ID:           1,
			CardNumberID: "402323",
			FirstName:    "Jhon",
			LastName:     "Doe",
			WarehouseID:  1,
		},
		{
			ID:           2,
			CardNumberID: "402322",
			FirstName:    "Jhon",
			LastName:     "Doe",
			WarehouseID:  3,
		},
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"})

	for _, e := range data {
		rows.AddRow(e.ID, e.CardNumberID, e.FirstName, e.LastName, e.WarehouseID)
	}
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees")).WillReturnError(errors.New("Get All Error"))
	repository := NewRepository(db)

	//Act
	employees, err := repository.GetAll(context.TODO())

	//Assert
	assert.NotNil(t, err)
	assert.Empty(t, employees)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_GetEmployee_Ok(t *testing.T) {
	//Arrange
	employee := domain.Employee{
		ID:           1,
		CardNumberID: "402323",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  1,
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"})
	rows.AddRow(employee.ID, employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id=?;")).WithArgs(employee.ID).WillReturnRows(rows)
	repository := NewRepository(db)

	//Act
	result, err := repository.Get(context.TODO(), employee.ID)

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, employee, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_GetEmployee_Fail(t *testing.T) {
	//Arrange
	employee := domain.Employee{
		ID:           1,
		CardNumberID: "402323",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  1,
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id"})
	rows.AddRow(employee.ID, employee.CardNumberID, employee.FirstName, employee.LastName, employee.WarehouseID)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id=?;")).WithArgs(2).WillReturnRows(rows)
	repository := NewRepository(db)

	//Act
	result, err := repository.Get(context.TODO(), employee.ID)

	//Assert
	assert.Error(t, err)
	assert.Empty(t, result)
}

func Test_SaveEmployee_Ok(t *testing.T) {
	//Arrange
	employee_test := domain.Employee{
		ID:           1,
		CardNumberID: "402323",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  1,
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)"))

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)")).
		WithArgs(employee_test.CardNumberID, employee_test.FirstName, employee_test.LastName, employee_test.WarehouseID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	repository := NewRepository(db)

	//Act
	id, err := repository.Save(context.TODO(), employee_test)

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, employee_test.ID, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_Save_EmployeeFail(t *testing.T) {
	//Arrange
	employee_test := domain.Employee{}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare(regexp.QuoteMeta("INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)"))

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)")).
		WithArgs(employee_test.CardNumberID, employee_test.FirstName, employee_test.LastName, employee_test.WarehouseID).
		WillReturnError(errors.New("Save Error"))
	repository := NewRepository(db)

	//Act
	id, err := repository.Save(context.TODO(), employee_test)

	//Assert
	assert.Error(t, err)
	assert.Empty(t, id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_ExistEmployee_Ok(t *testing.T) {
	//Arrange
	cardNumber := "402323"
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"card_number_id"})
	row.AddRow(cardNumber)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT card_number_id FROM employees WHERE card_number_id=?;")).
		WithArgs(cardNumber).WillReturnRows(row)

	repository := NewRepository(db)
	//Act
	result := repository.Exists(context.TODO(), cardNumber)

	//Assert
	assert.True(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_ExistEmployee_Fail(t *testing.T) {
	//Arrange
	cardNumber := "402323"
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"card_number_id"})

	mock.ExpectQuery(regexp.QuoteMeta("SELECT card_number_id FROM employees WHERE card_number_id=?;")).
		WithArgs(cardNumber).WillReturnRows(row)

	repository := NewRepository(db)
	//Act
	result := repository.Exists(context.TODO(), cardNumber)

	//Assert
	assert.False(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_UpdateEmployee_Ok(t *testing.T) {
	employee := domain.Employee{
		ID:           1,
		CardNumberID: "402323",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  1,
	}

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta("UPDATE employees SET first_name=?, last_name=?, warehouse_id=?  WHERE id=?"))

	mock.ExpectExec(regexp.QuoteMeta("UPDATE employees SET first_name=?, last_name=?, warehouse_id=?  WHERE id=?")).
		WithArgs(employee.FirstName, employee.LastName, employee.WarehouseID, employee.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	repository := NewRepository(db)

	//Act
	err = repository.Update(context.TODO(), employee)

	//Assert
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_UpdateEmployee_Fail(t *testing.T) {
	employee := domain.Employee{
		ID:           1,
		CardNumberID: "402323",
		FirstName:    "Jhon",
		LastName:     "Doe",
		WarehouseID:  1,
	}

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta("UPDATE employees SET first_name=?, last_name=?, warehouse_id=?  WHERE id=?"))

	mock.ExpectExec(regexp.QuoteMeta("UPDATE employees SET first_name=?, last_name=?, warehouse_id=?  WHERE id=?")).
		WithArgs(employee.FirstName, employee.LastName, employee.WarehouseID, employee.ID).
		WillReturnError(errors.New("Update Error"))
	repository := NewRepository(db)

	//Act
	err = repository.Update(context.TODO(), employee)

	//Assert
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_DeleteEmployee_Ok(t *testing.T) {
	//Arrange
	id := 1
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM employees WHERE id=?"))

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM employees WHERE id=?")).
		WithArgs(id).WillReturnResult(sqlmock.NewResult(1, 1))
	repository := NewRepository(db)

	//Act
	err = repository.Delete(context.TODO(), id)

	//Assert
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())

	/*t.Run("no employee exists", func(t *testing.T) {

	})*/
}

func Test_DeleteEmployee_Fail(t *testing.T) {
	id := 2
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	mock.ExpectPrepare(regexp.QuoteMeta("DELETE FROM employees WHERE id=?"))

	mock.ExpectExec(regexp.QuoteMeta("DELETE FROM employees WHERE id=?")).
		WithArgs(id).WillReturnError(errors.New("Delete Error"))

	repository := NewRepository(db)

	//Act
	err = repository.Delete(context.TODO(), id)

	//Assert
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_ReportInboundOrders_Ok(t *testing.T) {
	//Arrange
	reports := []domain.ReportInBO{
		{
			ID:                   1,
			CardNumberID:         "402323",
			FirstName:            "Jhon",
			LastName:             "Doe",
			WarehouseID:          1,
			Inbound_orders_count: 2,
		},
		{
			ID:                   2,
			CardNumberID:         "402324",
			FirstName:            "Jhon",
			LastName:             "Doe",
			WarehouseID:          1,
			Inbound_orders_count: 0,
		},
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"})

	for _, bo := range reports {
		rows.AddRow(bo.ID, bo.CardNumberID, bo.FirstName, bo.LastName, bo.WarehouseID, bo.Inbound_orders_count)
	}
	stmt := (regexp.QuoteMeta("SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id , count(inbo.employee_id) AS inbound_orders_count " +
		"FROM employees e LEFT JOIN inbound_orders inbo ON  e.id=inbo.employee_id " +
		"GROUP BY e.id;"))

	mock.ExpectQuery(stmt).WillReturnRows(rows)
	repository := NewRepository(db)

	//Act
	result, err := repository.ReportInboundOrders(context.TODO())

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, reports, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_ReportInboundOrders_Fail(t *testing.T) {
	//Arrange
	reports := []domain.ReportInBO{
		{
			ID:                   1,
			CardNumberID:         "402323",
			FirstName:            "Jhon",
			LastName:             "Doe",
			WarehouseID:          1,
			Inbound_orders_count: 2,
		},
		{
			ID:                   2,
			CardNumberID:         "402324",
			FirstName:            "Jhon",
			LastName:             "Doe",
			WarehouseID:          1,
			Inbound_orders_count: 0,
		},
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"})

	for _, bo := range reports {
		rows.AddRow(bo.ID, bo.CardNumberID, bo.FirstName, bo.LastName, bo.WarehouseID, bo.Inbound_orders_count)
	}
	stmt := regexp.QuoteMeta("SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id , count(inbo.employee_id) AS inbound_orders_count " +
		"FROM employees e LEFT JOIN inbound_orders inbo ON  e.id=inbo.employee_id " +
		"GROUP BY e.id;")

	mock.ExpectQuery(stmt).WillReturnError(errors.New("ReportInboundOrders Error"))
	repository := NewRepository(db)

	//Act
	result, err := repository.ReportInboundOrders(context.TODO())

	//Assert
	assert.NotNil(t, err)
	assert.Empty(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_ReportInboundOrdersByID_Ok(t *testing.T) {
	//Arrange
	report := []domain.ReportInBO{
		{
			ID:                   1,
			CardNumberID:         "402323",
			FirstName:            "Jhon",
			LastName:             "Doe",
			WarehouseID:          1,
			Inbound_orders_count: 2,
		},
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"})
	rows.AddRow(report[0].ID, report[0].CardNumberID, report[0].FirstName, report[0].LastName, report[0].WarehouseID, report[0].Inbound_orders_count)
	stmt := regexp.QuoteMeta("SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id , count(inbo.employee_id) AS inbound_orders_count " +
		"FROM employees e LEFT JOIN inbound_orders inbo ON  e.id= inbo.employee_id " +
		"WHERE e.id=? GROUP BY e.id;")

	mock.ExpectQuery(stmt).WithArgs(report[0].ID).WillReturnRows(rows)
	repository := NewRepository(db)

	//Act
	result, err := repository.ReportInboundOrdersByID(context.TODO(), report[0].ID)

	//Assert
	assert.NoError(t, err)
	assert.Equal(t, report, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func Test_ReportInboundOrdersByID_Fail(t *testing.T) {
	//Arrange
	report := []domain.ReportInBO{
		{
			ID:                   1,
			CardNumberID:         "402323",
			FirstName:            "Jhon",
			LastName:             "Doe",
			WarehouseID:          1,
			Inbound_orders_count: 2,
		},
	}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	rows := sqlmock.NewRows([]string{"id", "card_number_id", "first_name", "last_name", "warehouse_id", "inbound_orders_count"})
	rows.AddRow(report[0].ID, report[0].CardNumberID, report[0].FirstName, report[0].LastName, report[0].WarehouseID, report[0].Inbound_orders_count)
	stmt := regexp.QuoteMeta("SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id , count(inbo.employee_id) AS inbound_orders_count " +
		"FROM employees e LEFT JOIN inbound_orders inbo ON  e.id= inbo.employee_id " +
		"WHERE e.id=? GROUP BY e.id;")

	mock.ExpectQuery(stmt).WithArgs(2).WillReturnRows(rows)
	repository := NewRepository(db)

	//Act
	result, err := repository.ReportInboundOrdersByID(context.TODO(), report[0].ID)

	//Assert
	assert.Error(t, err)
	assert.Empty(t, result)
}
