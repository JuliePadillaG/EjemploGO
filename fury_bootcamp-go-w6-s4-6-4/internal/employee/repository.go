package employee

import (
	"context"
	"database/sql"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Repository encapsulates the storage of a employee.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Employee, error)
	Get(ctx context.Context, id int) (domain.Employee, error)
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, e domain.Employee) (int, error)
	Update(ctx context.Context, e domain.Employee) error
	Delete(ctx context.Context, id int) error

	ReportInboundOrders(ctx context.Context) ([]domain.ReportInBO, error)
	ReportInboundOrdersByID(ctx context.Context, id int) ([]domain.ReportInBO, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Employee, error) {
	query := "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var employees []domain.Employee

	for rows.Next() {
		e := domain.Employee{}
		_ = rows.Scan(&e.ID, &e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID)
		employees = append(employees, e)
	}

	return employees, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Employee, error) {
	query := "SELECT id, card_number_id, first_name, last_name, warehouse_id FROM employees WHERE id=?;"
	row := r.db.QueryRow(query, id)
	e := domain.Employee{}
	err := row.Scan(&e.ID, &e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID)
	if err != nil {
		return domain.Employee{}, err
	}

	return e, nil
}

func (r *repository) Exists(ctx context.Context, cardNumberID string) bool {
	query := "SELECT card_number_id FROM employees WHERE card_number_id=?;"
	row := r.db.QueryRow(query, cardNumberID)
	err := row.Scan(&cardNumberID)
	return err == nil
}

func (r *repository) Save(ctx context.Context, e domain.Employee) (int, error) {
	query := "INSERT INTO employees(card_number_id,first_name,last_name,warehouse_id) VALUES (?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&e.CardNumberID, &e.FirstName, &e.LastName, &e.WarehouseID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, e domain.Employee) error {
	query := "UPDATE employees SET first_name=?, last_name=?, warehouse_id=?  WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&e.FirstName, &e.LastName, &e.WarehouseID, &e.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM employees WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}

func (r *repository) ReportInboundOrders(ctx context.Context) ([]domain.ReportInBO, error) {
	//var rows *sql.Rows
	query := "SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id , count(inbo.employee_id) AS inbound_orders_count " +
		"FROM employees e LEFT JOIN inbound_orders inbo ON  e.id=inbo.employee_id " +
		"GROUP BY e.id;"
	rows, err := r.db.Query(query)
	if err != nil {
		return []domain.ReportInBO{}, err
	}
	var report []domain.ReportInBO
	for rows.Next() {
		var r domain.ReportInBO
		err := rows.Scan(&r.ID, &r.CardNumberID, &r.FirstName, &r.LastName, &r.WarehouseID, &r.Inbound_orders_count)
		if err != nil {
			return []domain.ReportInBO{}, err
		}
		report = append(report, r)
	}
	return report, nil
}

func (r *repository) ReportInboundOrdersByID(ctx context.Context, id int) ([]domain.ReportInBO, error) {
	query := "SELECT e.id, e.card_number_id, e.first_name, e.last_name, e.warehouse_id , count(inbo.employee_id) AS inbound_orders_count " +
		"FROM employees e LEFT JOIN inbound_orders inbo ON  e.id= inbo.employee_id " +
		"WHERE e.id=? GROUP BY e.id;"
	rows, err := r.db.Query(query, id)
	if err != nil {
		return []domain.ReportInBO{}, err
	}
	var report []domain.ReportInBO
	for rows.Next() {
		var r domain.ReportInBO
		err := rows.Scan(&r.ID, &r.CardNumberID, &r.FirstName, &r.LastName, &r.WarehouseID, &r.Inbound_orders_count)
		if err != nil {
			return []domain.ReportInBO{}, err
		}
		report = append(report, r)
	}
	return report, nil
}
