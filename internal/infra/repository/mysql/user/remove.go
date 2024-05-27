package user_mysql_repository

import (
	"database/sql"
	"fmt"
)

type removeUserRepository struct {
	UserRepositoryMysql
	DetailUserRepository
	DB *sql.DB
}

func NewRemoveUserRepository(db *sql.DB) *removeUserRepository {
	repository := &removeUserRepository{}
	repository.DB = db
	return repository
}

func (r *removeUserRepository) Remove(id string) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", TableName)

	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
