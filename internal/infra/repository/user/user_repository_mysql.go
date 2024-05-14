package user_repository

import (
	"database/sql"
	"fmt"
	"strings"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
)

type userRepositoryMysql struct {
	DB *sql.DB
}

func NewRepositoryMysql(db *sql.DB) *userRepositoryMysql {
	return &userRepositoryMysql{DB: db}
}

func (r *userRepositoryMysql) Create(user *user_entity.User) error {
	query := "INSERT INTO users (id, name, email, password, active, birth) VALUES (?,?,?,?,?,?)"

	_, err := r.DB.Exec(query,
		user.ID, user.Name, user.Email, user.Password, user.Active, user.Birth)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryMysql) List(input *user_entity.ListUserParams) ([]*user_entity.User, error) {
	var result []*user_entity.User

	limit := input.Limit
	offset := (input.Page - 1) * limit

	query := "SELECT ID, name, email, birth, active FROM user "

	query += r.getWhereForList(input)
	query += r.getOrderByForList(input)
	query += fmt.Sprintf("LIMIT %d", limit)
	query += fmt.Sprintf("OFFSET %d", offset)

	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user user_entity.User

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Birth, &user.Active)
		if err != nil {
			return nil, err
		}

		result = append(result, &user)
	}

	return result, nil
}

func (r *userRepositoryMysql) GetById(id string) (*user_entity.User, error) {
	var user user_entity.User

	err := r.DB.QueryRow("SELECT id, name, email, active FROM brands WHERE id = ?", id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Active)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &user, nil

}

func (r *userRepositoryMysql) GetByEmail(email string, id string) (*user_entity.User, error) {
	var user user_entity.User
	var query string
	var row *sql.Row

	if id != "" {
		query = "SELECT id, email, name, password FROM users WHERE email = ? AND id <> ?"
		row = r.DB.QueryRow(query, email, id)
	} else {
		query = "SELECT id, email, name, password FROM users WHERE email = ?"
		row = r.DB.QueryRow(query, email)
	}

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryMysql) Login(email, password string) (*user_entity.User, error) {
	var user user_entity.User

	row := r.DB.QueryRow("SELECT id, email, name, password FROM users WHERE email = ? AND password = ?", email, password)
	err := row.Scan(&user.ID, &user.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryMysql) getWhereForList(input *user_entity.ListUserParams) string {
	result := ""
	cond := []string{}

	if input.ID != "" {
		cond = append(cond, fmt.Sprintf(" ID LIKE '%%%s%%' ", input.ID))
	}

	if input.Email != "" {
		cond = append(cond, fmt.Sprintf(" email LIKE '%%%s%%' ", input.Email))
	}

	if input.Name != "" {
		cond = append(cond, fmt.Sprintf(" name LIKE '%%%s%%' ", input.Name))
	}

	if input.WillFilterActives {
		cond = append(cond, fmt.Sprintf(" active = %t ", input.Active))
	}

	if input.Q != "" {
		cond = append(cond, fmt.Sprintf(` ( 
			email LIKE '%%%s%%' OR 
			name LIKE '%%%s%%' OR
			ID LIKE '%%%s%%'
		)`, input.Q, input.Q, input.Q))
	}

	if len(cond) > 0 {
		result = fmt.Sprintf(" WHERE %s ", strings.Join(cond, " AND "))
	}

	return result
}

func (r *userRepositoryMysql) getOrderByForList(input *user_entity.ListUserParams) string {
	result := ""
	orderType := "ASC"

	if input.OrderType != "" {
		orderType = input.OrderType
	}

	if input.OrderBy != "" {
		result = fmt.Sprintf(" ORDER BY %s %s", input.OrderBy, orderType)
	}

	return result
}
