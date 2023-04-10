package purchase_orders

import (
	"context"
	"database/sql"
	"log"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

// Repository encapsulates the storage of the purchase orders.
type Repository interface {
	Get(ctx context.Context, id int) ([]domain.ReportPurchaseOrders, error)
	ExistsBuyersID(ctx context.Context, buyerID int) bool
	ExistsProductRecordsID(ctx context.Context, productRecordID int) bool
	Save(ctx context.Context, p domain.PurchaseOrders) (int, error)
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
        SAVE_BUYER = `
                INSERT INTO purchase_orders(
                        order_number, order_date, tracking_code, buyers_id, product_records_id, order_status_id)
                VALUES (?,?,?,?,?,?);`
        EXISTS_PRODUCT_RECORD_ID =  `SELECT id FROM purchase_orders WHERE id=?;`
        EXISTS_BUYER_ID =  `SELECT id FROM buyers WHERE id=?;`
        GET_REPORT_PURCHASEORDERS_BY_BUYERID = `
                SELECT b.id, b.card_number_id, b.first_name, b.last_name, COUNT(*) AS purchase_orders_count
                FROM purchase_orders p
                LEFT JOIN buyers b ON p.buyers_id = b.id
                WHERE b.id = ?
                GROUP BY b.id;`
        GET_REPORT_PURCHASEORDERS = `
                SELECT b.id, b.card_number_id, b.first_name, b.last_name, COUNT(*) AS purchase_orders_count
                FROM purchase_orders p
                LEFT JOIN buyers b ON p.buyers_id = b.id
                GROUP BY b.id;`
)


func (r *repository) ExistsBuyersID(ctx context.Context, buyerID int) bool {
	query := EXISTS_BUYER_ID
	row := r.db.QueryRow(query, buyerID)
	err := row.Scan(&buyerID)
	return err == nil
}


func (r *repository) ExistsProductRecordsID(ctx context.Context, productRecordID int) bool {
	query := EXISTS_PRODUCT_RECORD_ID
	row := r.db.QueryRow(query, productRecordID)
	err := row.Scan(&productRecordID)
	return err == nil
}

func (r *repository) Get(ctx context.Context, id int) ([]domain.ReportPurchaseOrders, error) {

        var (
                query string
                err error
                rows *sql.Rows
        )

        if id != 0 {
                query = GET_REPORT_PURCHASEORDERS_BY_BUYERID
	        rows, err = r.db.Query(query, id)
        } else {
                query = GET_REPORT_PURCHASEORDERS
	        rows, err = r.db.Query(query)
        }

	if err != nil {
		return nil, err
	}

        var report []domain.ReportPurchaseOrders

	for rows.Next() {
		p := domain.ReportPurchaseOrders{}
		_ = rows.Scan(&p.ID, &p.CardNumberID, &p.FirstName, &p.LastName, &p.PurchaseOrdersCount)
		report = append(report, p)
	}

	return report, nil
}


func (r *repository) Save(ctx context.Context, p domain.PurchaseOrders) (int, error) {

        query := SAVE_BUYER
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&p.OrderNumber, &p.OrderDate, &p.TrackingCode, &p.BuyerID, &p.ProductRecordID, &p.OrderStatusID)
	if err != nil {
		return 0, err
	}

        log.Print(res)

        id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

