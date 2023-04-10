package inboundorder

import (
	"context"
	"database/sql"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Inbound_order, error)
	Save(ctx context.Context, b_order domain.Inbound_order) (int, error)
	ExistsEmployee(ctx context.Context, id_employee int) bool
	ExistsInboundOrder(ctx context.Context, order_number string) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	GET_ALL        = "SELECT id, order_date, order_number, employee_id, warehouse_id, product_batch_id FROM inbound_orders"
	SAVE           = "INSERT INTO inbound_orders(order_date,order_number,employee_id,product_batch_id,warehouse_id) VALUES (?,?,?,?,?)"
	EXIST_EMPLOYEE = "SELECT id FROM employees WHERE id=?"
	EXIST_INBOUND  = "SELECT order_number FROM inbound_orders WHERE order_number=?"
)

func (r *repository) GetAll(ctx context.Context) ([]domain.Inbound_order, error) {
	rows, err := r.db.Query(GET_ALL)
	if err != nil {
		return nil, err
	}

	var inbound_orders []domain.Inbound_order

	for rows.Next() {
		inborder := domain.Inbound_order{}
		_ = rows.Scan(&inborder.ID, &inborder.Order_date, &inborder.Order_number, &inborder.Employee_id, &inborder.Product_batch_id, &inborder.Warehouse_id)
		inbound_orders = append(inbound_orders, inborder)
	}

	return inbound_orders, nil
}

func (r *repository) ExistsEmployee(ctx context.Context, id_employee int) bool {
	row := r.db.QueryRow(EXIST_EMPLOYEE, id_employee)
	err := row.Scan(&id_employee)
	return err == nil
}

func (r *repository) ExistsInboundOrder(ctx context.Context, order_number string) bool {
	row := r.db.QueryRow(EXIST_INBOUND, order_number)
	err := row.Scan(&order_number)
	return err == nil
}
func (r *repository) Save(ctx context.Context, b_order domain.Inbound_order) (int, error) {
	stmt, err := r.db.Prepare(SAVE)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&b_order.Order_date, &b_order.Order_number, &b_order.Employee_id, &b_order.Product_batch_id, &b_order.Warehouse_id)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
