package user_mysql_repository

import (
	"database/sql"
	"fmt"
	"strings"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
)

type updateUserRepository struct {
	UserRepositoryMysql
	DetailUserRepository
}

func NewUpdateUserRepository(db *sql.DB) *updateUserRepository {
	repository := &updateUserRepository{}
	repository.DB = db
	return repository
}

func (r *updateUserRepository) GetForUpdate(id string) (*user_entity.User, error) {
	var user user_entity.User

	query := fmt.Sprintf("SELECT id, name, password, active, birth FROM %s WHERE id = ?", TableName)

	var birth string

	err := r.DB.QueryRow(query, id).
		Scan(&user.ID, &user.Name, &user.Password, &user.Active, &birth)

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

func (r *updateUserRepository) Update(id string, editedUser *user_entity.User) (*user_entity.User, error) {

	query := fmt.Sprintf("UPDATE %s SET ##fields## WHERE id = ?", TableName)

	var fields []string

	fields = append(fields, "name = ?")
	fields = append(fields, "password = ?")
	fields = append(fields, "birth = ?")
	fields = append(fields, "active = ?")

	query = strings.Replace(query, "##fields##", strings.Join(fields, ","), 1)

	_, err := r.DB.Exec(query, editedUser.Name, editedUser.Password, editedUser.Birth, editedUser.Active, id)
	if err != nil {
		return nil, err
	}

	return r.GetById(id)
}
