package user_mysql_repository

import (
	"database/sql"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	default_mysql_repository "github.com/mariojuniortrab/hauling-api/internal/infra/repository/mysql/default"
)

type getUserByEmailRepository struct {
	UserRepositoryMysql
}

func NewGetUserByEmailRepository(db *sql.DB) *getUserByEmailRepository {
	repository := &getUserByEmailRepository{}
	repository.DB = db
	return repository
}

func (r *getUserByEmailRepository) GetByEmail(email string, id string) (*user_entity.User, error) {
	whereMap := map[string]interface{}{
		"email": email,
		"id":    id,
	}
	fieldsToGet := []string{"id", "email", "name", "password"}

	mappedResult, err := default_mysql_repository.GetByField(r, fieldsToGet, whereMap)
	if err != nil {
		return nil, err
	}

	user, err := user_entity.NewUserFromMap(mappedResult)
	if err != nil {
		return nil, err
	}

	return user, nil
}
