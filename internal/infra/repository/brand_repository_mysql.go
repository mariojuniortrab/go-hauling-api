package repository

import (
	"database/sql"

	"github.com/mariojuniortrab/hauling-api/internal/entity"
)

type BrandRepositoryMysql struct {
	DB *sql.DB
}

func NewBrandRepositoryMysql(db *sql.DB) *BrandRepositoryMysql {
	return &BrandRepositoryMysql{DB: db}
}

func (r *BrandRepositoryMysql) Create(brand *entity.Brand) error {
	_, err := r.DB.Exec("Insert into Brands (id, name) values (?,?)")

	if err != nil {
		return err
	}

	return nil
}

func (r *BrandRepositoryMysql) FindAll() ([]*entity.Brand, error) {
	rows, err := r.DB.Query("Select id, name from brands")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var brands []*entity.Brand

	for rows.Next() {
		var brand entity.Brand

		err = rows.Scan(&brand.ID, &brand.Name)
		if err != nil {
			return nil, err
		}

		brands = append(brands, &brand)
	}

	return brands, nil
}
