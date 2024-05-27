package user_mysql_repository

import (
	"database/sql"
	"fmt"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
)

type signupRepository struct {
	DetailUserRepository
	UserRepositoryMysql
}

func NewSignupRepository(db *sql.DB) *signupRepository {
	repository := &signupRepository{}
	repository.DB = db
	return repository
}

func (r *UserRepositoryMysql) Create(user *user_entity.User) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name, email, password, active, birth) VALUES (?,?,?,?,?,?)", TableName)

	_, err := r.DB.Exec(query,
		user.ID, user.Name, user.Email, user.Password, user.Active, user.Birth)

	if err != nil {
		return err
	}

	return nil
}
