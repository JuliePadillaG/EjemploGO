package locality

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/mercadolibre/fury_bootcamp-go-w6-s4-6-4/internal/domain"
)

const (
	GET_SELLERS_BY_ID = "SELECT l.id, l.locality_name, COUNT(l.id) FROM seller s INNER JOIN locality l ON s.locality_id = l.id WHERE l.id=? GROUP BY id;"
	GET_SELLERS       = "SELECT l.id, l.locality_name, COUNT(l.id) FROM seller s INNER JOIN locality l ON s.locality_id = l.id GROUP BY id;"
	EXIST_LOCALITY    = "SELECT id FROM locality WHERE id=?;"
	GET_LOCALITY      = "SELECT id, locality_name, province_name, country_name FROM locality WHERE id =?;"
	CREATE_LOCALITY   = "INSERT INTO locality (id, locality_name, province_name, country_name) VALUES (?, ?, ?, ?)"
)

type Repository interface {
	Exists(ctx context.Context, id int) bool
	Create(ctx context.Context, l domain.Locality) (int, error)
	GetAllSellersByLocality(ctx context.Context, localityID string) ([]domain.ResponseLocality, error)
	GetCarriesReport(ctx context.Context, id string) ([]domain.CarriesReport, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Exists(ctx context.Context, id int) bool {
	row := r.db.QueryRow(EXIST_LOCALITY, id)
	err := row.Scan(&id)
	return err == nil
}

func (r *repository) Create(ctx context.Context, l domain.Locality) (int, error) {
	stmt, err := r.db.Prepare(CREATE_LOCALITY)
	if err != nil {
		return 0, err
	}

	if r.Exists(ctx, l.ID) {
		return 0, errors.New("id already exists")
	}

	res, err := stmt.Exec(l.ID, l.LocalityName, l.ProvinceName, l.CountryName)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) GetAllSellersByLocality(ctx context.Context, localityID string) ([]domain.ResponseLocality, error) {
	var rows *sql.Rows
	var err error

	if localityID == "" {
		rows, err = r.db.Query(GET_SELLERS)
	} else {
		localityID, _ := strconv.Atoi(localityID)
		if r.Exists(ctx, localityID) {
			rows, err = r.db.Query(GET_SELLERS_BY_ID, localityID)
		} else {
			return []domain.ResponseLocality{}, errors.New("locality_id not found")
		}
	}

	if err != nil {
		return []domain.ResponseLocality{}, err
	}

	var localities []domain.ResponseLocality

	for rows.Next() {
		var l domain.ResponseLocality
		err := rows.Scan(&l.ID, &l.LocalityName, &l.SellersCount)
		if err != nil {
			return []domain.ResponseLocality{}, err
		}
		localities = append(localities, l)
	}

	return localities, nil
}

func (r *repository) GetCarriesReport(ctx context.Context, id string) (carriesReports []domain.CarriesReport, err error) {
	var rows *sql.Rows
	if id == "" {
		query := "SELECT locality.id, locality.locality_name, COUNT(*) AS carries_count FROM carries RIGHT JOIN locality on carries.locality_id = locality.id GROUP BY id;"
		rows, err = r.db.Query(query)
	} else {
		intId, _ := strconv.Atoi(id)
		if !r.Exists(ctx, intId) {
			return nil, errors.New("id does not exist")
		}
		query := "SELECT locality.id, locality.locality_name, COUNT(*) AS carries_count FROM carries right join locality on carries.locality_id = locality.id WHERE locality.id = ? GROUP BY id;"
		rows, err = r.db.Query(query, intId)
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var report domain.CarriesReport
		err = rows.Scan(&report.LocalityID, &report.LocalityName, &report.CarriesCount)
		if err != nil {
			return nil, err
		}
		carriesReports = append(carriesReports, report)
	}

	return carriesReports, nil
}
