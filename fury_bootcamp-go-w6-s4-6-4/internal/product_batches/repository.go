package productbatches

import (
	"context"
	"database/sql"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)


const(
	//SAVE_MOVIE = "INSERT INTO movies (title, rating, awards, length, genre_id) VALUES (?, ?, ?, ?, ?);"

	CREATE_PRODUCT_BATCH = `INSERT INTO product_batches(batch_number, current_quantity, current_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, minimum_temperature, sections_id, products_id) VALUES(?,?,?,?,?,?,?,?,?,?);`

	GET_PRODUCT_BATCH = `SELECT * FROM product_batches WHERE id = ?;`

	READ_PRODUCT_BATCH = `SELECT p.sections_id, s.section_number, SUM(p.current_quantity) cq FROM melisprint.product_batches p INNER JOIN melisprint.sections s ON p.sections_id = s.id WHERE s.id=? GROUP BY p.sections_id;`

	EXISTS_SECTION_ID = `SELECT sections_id FROM melisprint.product_batches WHERE sections_id=?;`

	EXISTS_PRODUCT_ID = `SELECT products_id FROM melisprint.product_batches WHERE products_id=?;`

	EXISTS = `SELECT batch_number FROM melisprint.product_batches WHERE batch_number =?;`
)


type Repository interface {
	CreatePB(ctx context.Context, pb domain.Product_batches) (int, error)
	ReadPB(ctx context.Context, id int) (domain.ReportProduct, error)
	ExistenceSectionId(ctx context.Context, section_id int) bool
	ExistenceProductId(ctx context.Context, product_id int) bool
	GetPB(ctx context.Context, id int) (domain.Product_batches, error)
	ExistsProductBatches(ctx context.Context, batch_number int) bool
}

type repository struct{
	db *sql.DB
}


func NewRepository(db *sql.DB) Repository{
	return &repository{
		db: db,
	}
}


func (r *repository) CreatePB(ctx context.Context, pb domain.Product_batches) (int, error){
	query := CREATE_PRODUCT_BATCH
	stmt, err := r.db.Prepare(query)
	if err != nil{
		return 0, err
	}

	res, err := stmt.Exec(&pb.BatchNumber, &pb.CurrentQuantity, &pb.CurrentTemperature, &pb.DueDate, &pb.InitialQuantity, &pb.ManufacturingDate, &pb.ManufacturingHour, &pb.MinimumTemperature, &pb.SectionId, &pb.ProductId)
	if err != nil{
		
		return 0, err
		
	}
	
	id, err := res.LastInsertId()
	if err != nil{
		return 0, err
	}
	

	return int(id), nil
}

func (r *repository) ReadPB(ctx context.Context, id int) (domain.ReportProduct, error) {
	query := READ_PRODUCT_BATCH
	row := r.db.QueryRow(query,id)
	//s := domain.Section{}
	//p := domain.Product_batches{}
	data := domain.ReportProduct{}
	
	err := row.Scan(&data.SectionId, &data.SectionNumber, &data.CurrentQuantity)
	if err != nil {
		return domain.ReportProduct{}, err
	}
	
	
	return data, nil	
}

func (r *repository) GetPB(ctx context.Context, id int) (domain.Product_batches, error){
	query := GET_PRODUCT_BATCH
	row := r.db.QueryRow(query, id)
	pb := domain.Product_batches{}
	err := row.Scan(&pb.ID, &pb.BatchNumber, &pb.CurrentQuantity, &pb.CurrentTemperature, &pb.DueDate, &pb.InitialQuantity, &pb.ManufacturingDate, &pb.ManufacturingHour, &pb.MinimumTemperature, &pb.SectionId, &pb.ProductId)
	if err != nil{
		return pb, err
	}
	return pb, nil
}

func (r *repository) ExistenceSectionId(ctx context.Context, section_id int) bool { 
	query := EXISTS_SECTION_ID
	row := r.db.QueryRow(query, section_id)
	err := row.Scan(&section_id)


	return err == nil
}

func (r *repository) ExistenceProductId(ctx context.Context, product_id int) bool { 
	query := EXISTS_PRODUCT_ID
	row := r.db.QueryRow(query, product_id)
	err := row.Scan(&product_id)
	return err == nil
}

func (r *repository) ExistsProductBatches(ctx context.Context, batch_number int) bool {
	query := EXISTS
	row := r.db.QueryRow(query, batch_number)
	err := row.Scan(&batch_number)
	return err == nil
}