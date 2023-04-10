package product_records

import (
	"context"
	"database/sql"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, pr domain.ProductRecords) (int, error)
	ExistsProductRecord(ctx context.Context, id int) bool
	UniqueProduct(ctx context.Context, productID int) bool
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
	SAVE_PRODUCT_RECORD = "INSERT INTO product_records (last_update_date, purchase_price, sale_price, products_id) VALUES (?, ?, ?, ?);"

	EXIST_PRODUCT_RECORD = "SELECT pr.id FROM product_records pr WHERE pr.id=?"

	UNIQUE_PRODUCT = "SELECT p.id FROM products p WHERE p.id=?"
)

func (r *repository) ExistsProductRecord(ctx context.Context, id int) bool {
	row := r.db.QueryRow(EXIST_PRODUCT_RECORD, id)
	err := row.Scan(&id)
	return err == nil
}

func (r *repository) UniqueProduct(ctx context.Context, productID int) bool {
	row := r.db.QueryRow(UNIQUE_PRODUCT, productID)
	err := row.Scan(&productID)
	return err == nil
}

func (r *repository) Save(ctx context.Context, pr domain.ProductRecords) (int, error) {
	stm, err := r.db.Prepare(SAVE_PRODUCT_RECORD)

	if err != nil {
		return 0, err
	}

	result, err := stm.Exec(pr.LastUpdateDate, pr.PurchasePrice, pr.SalePrice, pr.ProductID)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}
