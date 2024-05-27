package user_mysql_repository

import (
	"database/sql"
	"fmt"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
)

type loginUserRepository struct {
	UserRepositoryMysql
}

func NewLoginRepository(db *sql.DB) *loginUserRepository {
	repository := &loginUserRepository{}
	repository.DB = db
	return repository
}

func (r *UserRepositoryMysql) GetByEmail(email string, id string) (*user_entity.User, error) {
	var user user_entity.User
	var row *sql.Row

	query := fmt.Sprintf("SELECT id, email, name, password, active, birth FROM %s", TableName)

	if id != "" {
		query += " WHERE email = ? AND id <> ?"
		row = r.DB.QueryRow(query, email, id)
	} else {
		query += " WHERE email = ?"
		row = r.DB.QueryRow(query, email)
	}

	var birth string

	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Password, &user.Active, &birth)

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

func (r *UserRepositoryMysql) Login(email, password string) (*user_entity.User, error) {
	var user user_entity.User

	query := fmt.Sprintf("SELECT id, email, name, password FROM %s WHERE email = ? AND password = ?", TableName)
	row := r.DB.QueryRow(query, email, password)
	err := row.Scan(&user.ID, &user.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
