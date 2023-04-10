package carry

import (
	"context"
	"database/sql"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, c domain.Carry) (int, error)
	Exists(ctx context.Context, carryID string) bool
	ExistsLocality(ctx context.Context, localityID int) bool
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, c domain.Carry) (int, error) {
	query := "INSERT INTO carries (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?);"
	stmt, err := r.db.Prepare(query)

	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(c.CID, c.Company_name, c.Address, c.Telephone, c.Locality_id)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Exists(ctx context.Context, carryID string) bool {
	query := "SELECT cid FROM carries WHERE cid=?;"
	row := r.db.QueryRow(query, carryID)
	err := row.Scan(&carryID)
	return err == nil

}

func (r *repository) ExistsLocality(ctx context.Context, localityID int) bool {
	query := "SELECT id FROM locality WHERE id=?;"
	row := r.db.QueryRow(query, localityID)
	err := row.Scan(&localityID)
	return err == nil
}
