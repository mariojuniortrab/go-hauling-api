package user_mysql_repository

import (
	"database/sql"
	"fmt"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
)

type DetailUserRepository struct {
	UserRepositoryMysql
}

func NewDetailUserRepository(db *sql.DB) *DetailUserRepository {
	repository := &DetailUserRepository{}
	repository.DB = db
	return repository
}

func (r *DetailUserRepository) GetById(id string) (*user_entity.User, error) {
	var user user_entity.User

	query := fmt.Sprintf("SELECT id, name, email, active, birth FROM %s WHERE id = ?", TableName)

	var birth string

	err := r.DB.QueryRow(query, id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Active, &birth)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	parsedDate, err := util_entity.GetDateFromString(birth)
	if err != nil {
		return nil, err
	}

	user.Birth = parsedDate

	return &user, nil

}
