package user_repository

import (
	"database/sql"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/entity/user"
)

type userRepositoryMysql struct {
	DB *sql.DB
}

func NewRepositoryMysql(db *sql.DB) *userRepositoryMysql {
	return &userRepositoryMysql{DB: db}
}

func (r *userRepositoryMysql) Create(user *user_entity.User) error {
	_, err := r.DB.Exec("INSERT INTO users (id, name, username, password, active) VALUES (?,?)",
		user.ID, user.Name, user.Username, user.Password, user.Active)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryMysql) ListAll() ([]*user_entity.User, error) {
	var result []*user_entity.User

	rows, err := r.DB.Query("SELECT id, name, username, active FROM users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user user_entity.User

		err = rows.Scan(&user.ID, &user.Name, &user.Username, &user.Active)
		if err != nil {
			return nil, err
		}

		result = append(result, &user)
	}

	return result, nil
}

func (r *userRepositoryMysql) GetById(id string) (*user_entity.User, error) {
	var user user_entity.User

	err := r.DB.QueryRow("SELECT id, name, username, active FROM brands WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Username, &user.Active)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &user, nil

}

func (r *userRepositoryMysql) GetByUsername(username string) (*user_entity.User, error) {
	var user user_entity.User

	row := r.DB.QueryRow("SELECT id, username, name, password FROM users WHERE username = ?", username)
	err := row.Scan(&user.ID, &user.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
