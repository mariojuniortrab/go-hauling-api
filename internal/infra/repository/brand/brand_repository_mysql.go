package brand_repository

import (
	"database/sql"

	brand_entity "github.com/mariojuniortrab/hauling-api/internal/entity/brand"
)

type BrandRepositoryMysql struct {
	DB *sql.DB
}

func NewRepositoryMysql(db *sql.DB) *BrandRepositoryMysql {
	return &BrandRepositoryMysql{DB: db}
}

func (r *BrandRepositoryMysql) Create(brand *brand_entity.Brand) error {
	_, err := r.DB.Exec("insert into brands (id, name) values (?,?)", brand.ID, brand.Name)

	if err != nil {
		return err
	}

	return nil
}

func (r *BrandRepositoryMysql) ListAll() ([]*brand_entity.Brand, error) {
	rows, err := r.DB.Query("select id, name from brands")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var brands []*brand_entity.Brand

	for rows.Next() {
		var brand brand_entity.Brand

		err = rows.Scan(&brand.ID, &brand.Name)
		if err != nil {
			return nil, err
		}

		brands = append(brands, &brand)
	}

	return brands, nil
}

func (r *BrandRepositoryMysql) GetById(id string) (*brand_entity.Brand, error) {
	var brand brand_entity.Brand

	err := r.DB.QueryRow("select id, name from brands where id = ?", id).Scan(&brand.ID, &brand.Name)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &brand, nil

}

func (r *BrandRepositoryMysql) GetByName(name string) (*brand_entity.Brand, error) {
	var brand brand_entity.Brand

	row := r.DB.QueryRow("select id, name from brands where name = ?", name)
	err := row.Scan(&brand.ID, &brand.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &brand, nil

}

func (r *BrandRepositoryMysql) GetByNameForEdition(name string, id string) (*brand_entity.Brand, error) {
	var brand brand_entity.Brand

	err := r.DB.QueryRow("select id, name from brands where name = ? and id <> ?", name, id).Scan(&brand.ID, &brand.Name)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &brand, nil

}
