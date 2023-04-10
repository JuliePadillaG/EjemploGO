package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Repository encapsulates the storage of a Product.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Exists(ctx context.Context, productCode string) bool
	Save(ctx context.Context, p domain.Product) (int, error)
	Update(ctx context.Context, p domain.Product) error
	Delete(ctx context.Context, id int) error
	GetProductRecords(ctx context.Context, id string) (product_records_report []domain.ProductRecordsReport, err error)
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
	GET_ALL_PRODUCTS = "SELECT * FROM products;"

	GET_PRODUCT_BY_ID = "SELECT * FROM products WHERE id=?;"

	EXISTS_PRODUCT = "SELECT product_code FROM products WHERE product_code=?;"

	SAVE_PRODUCT = "INSERT INTO products(description,expiration_rate,freezing_rate,height,length,netweight,product_code,recommended_freezing_temperature,width,product_type_id,seller_id) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	
	UPDATE_PRODUCT = "UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, length=?, netweight=?, product_code=?, recommended_freezing_temperature=?, width=?, product_type_id=?, seller_id=?  WHERE id=?"
	
	DELETE_PRODUCT = "DELETE FROM products WHERE id=?"

	GET_PRODUCT_RECORDS_BY_PRODUCT_WITHOUT_ID = "SELECT p.id, product_code, COUNT(product_records.product_id) AS report_products_count FROM products p LEFT JOIN product_records on p.id = product_records.products_id GROUP BY id"

	GET_PRODUCT_RECORDS_BY_PRODUCT_WITH_ID = "SELECT p.id, product_code, COUNT(product_records.product_id) AS report_products_count FROM products p LEFT JOIN product_records on p.id = product_records.products_id WHERE p.id = ? GROUP BY id"
)

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	rows, err := r.db.Query(GET_ALL_PRODUCTS)
	if err != nil {
		return nil, err
	}

	var products []domain.Product

	for rows.Next() {
		p := domain.Product{}
		_ = rows.Scan(&p.ID, &p.Description, &p.ExpirationRate, &p.FreezingRate, &p.Height, &p.Length, &p.Netweight, &p.ProductCode, &p.RecomFreezTemp, &p.Width, &p.ProductTypeID, &p.SellerID)
		products = append(products, p)
	}

	return products, nil
}

func (r *repository) Get(ctx context.Context, id int) (domain.Product, error) {
	row := r.db.QueryRow(GET_PRODUCT_BY_ID, id)
	p := domain.Product{}
	err := row.Scan(&p.ID, &p.Description, &p.ExpirationRate, &p.FreezingRate, &p.Height, &p.Length, &p.Netweight, &p.ProductCode, &p.RecomFreezTemp, &p.Width, &p.ProductTypeID, &p.SellerID)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}

func (r *repository) Exists(ctx context.Context, productID string) bool {
	row := r.db.QueryRow(EXISTS_PRODUCT, productID)
	err := row.Scan(&productID)
	return err == nil
}

func (r *repository) Save(ctx context.Context, p domain.Product) (int, error) {
	stmt, err := r.db.Prepare(SAVE_PRODUCT)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(p.Description, p.ExpirationRate, p.FreezingRate, p.Height, p.Length, p.Netweight, p.ProductCode, p.RecomFreezTemp, p.Width, p.ProductTypeID, p.SellerID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, p domain.Product) error {
	stmt, err := r.db.Prepare(UPDATE_PRODUCT)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(p.Description, p.ExpirationRate, p.FreezingRate, p.Height, p.Length, p.Netweight, p.ProductCode, p.RecomFreezTemp, p.Width, p.ProductTypeID, p.SellerID, p.ID)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("error: no affected rows")
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DELETE_PRODUCT)
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

func (r *repository) GetProductRecords(ctx context.Context, id string) (product_records_report []domain.ProductRecordsReport, err error) {
	var rows *sql.Rows
	if id == "" {
		rows, err = r.db.Query(GET_PRODUCT_RECORDS_BY_PRODUCT_WITHOUT_ID)
	} else {
		if !r.Exists(ctx, id) {
			return nil, errors.New("id does not exist")
		}
		rows, err = r.db.Query(GET_PRODUCT_RECORDS_BY_PRODUCT_WITH_ID, id)
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var report domain.ProductRecordsReport
		err = rows.Scan(&report.ProductID, &report.ProductDescription, &report.ProductRecordsCount)
		if err != nil {
			return nil, err
		}
		product_records_report = append(product_records_report, report)
	}

	return product_records_report, nil
}


