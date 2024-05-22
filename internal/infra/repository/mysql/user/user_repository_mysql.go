package user_mysql_repository

import (
	"database/sql"
	"fmt"
	"strings"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
)

const tableName = "users"

type userRepositoryMysql struct {
	DB *sql.DB
}

func NewRepositoryMysql(db *sql.DB) *userRepositoryMysql {
	return &userRepositoryMysql{DB: db}
}

func (r *userRepositoryMysql) Create(user *user_entity.User) error {
	query := fmt.Sprintf("INSERT INTO %s (id, name, email, password, active, birth) VALUES (?,?,?,?,?,?)", tableName)

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

	query := fmt.Sprintf("SELECT ID, name, email, birth, active FROM %s ", tableName)

	query += r.getWhereForList(input)
	query += r.getOrderByForList(input)
	query += fmt.Sprintf(" LIMIT %d ", limit)
	query += fmt.Sprintf(" OFFSET %d ", offset)

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user user_entity.User
		var birth string

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &birth, &user.Active)
		if err != nil {
			return nil, err
		}

		parsedDate, err := util_entity.GetDateFromString(birth)
		if err != nil {
			return nil, err
		}

		user.Birth = parsedDate

		result = append(result, &user)
	}

	return result, nil
}

func (r *userRepositoryMysql) GetById(id string) (*user_entity.User, error) {
	var user user_entity.User

	query := fmt.Sprintf("SELECT id, name, email, active, birth FROM %s WHERE id = ?", tableName)

	var birth string

	err := r.DB.QueryRow(query, id).
		Scan(&user.ID, &user.Name, &user.Email, &user.Active, birth)

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

func (r *userRepositoryMysql) GetByEmail(email string, id string) (*user_entity.User, error) {
	var user user_entity.User
	var row *sql.Row

	query := fmt.Sprintf("SELECT id, email, name, password, active, birth FROM %s", tableName)

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

func (r *userRepositoryMysql) Login(email, password string) (*user_entity.User, error) {
	var user user_entity.User

	query := fmt.Sprintf("SELECT id, email, name, password FROM %s WHERE email = ? AND password = ?", tableName)
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
	orderType := " ASC "

	if input.OrderType != "" {
		orderType = " " + input.OrderType + " "
	}

	if input.OrderBy != "" {
		result = fmt.Sprintf(" ORDER BY %s %s ", input.OrderBy, orderType)
	}

	return result
}

func (r *userRepositoryMysql) Remove(id string) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", tableName)

	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryMysql) Update(id string, editedUser *user_entity.User) (*user_entity.User, error) {

	query := fmt.Sprintf("UPDATE %s SET ##fields## WHERE id = ?", tableName)

	var fields []string

	fields = append(fields, "name = ?")
	fields = append(fields, "password = ?")
	fields = append(fields, "birth = ?")
	fields = append(fields, "active = ?")

	query = strings.Replace(query, "##fields##", strings.Join(fields, ","), 1)

	_, err := r.DB.Exec(query, id, editedUser.Name, editedUser.Password, editedUser.Birth, editedUser.Active)
	if err != nil {
		return nil, err
	}

	return r.GetById(id)
}
