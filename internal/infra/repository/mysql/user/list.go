package user_mysql_repository

import (
	"database/sql"
	"fmt"
	"strings"

	user_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/user"
	util_entity "github.com/mariojuniortrab/hauling-api/internal/domain/entity/util"
)

type listUserRepository struct {
	UserRepositoryMysql
}

func NewListUserRepository(db *sql.DB) *listUserRepository {
	repository := &listUserRepository{}
	repository.DB = db
	return repository
}

func (r *UserRepositoryMysql) List(input *user_entity.ListUserParams) ([]*user_entity.User, error) {
	var result []*user_entity.User

	limit := input.Limit
	offset := (input.Page - 1) * limit

	query := fmt.Sprintf("SELECT ID, name, email, birth, active FROM %s ", TableName)

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

func (r *UserRepositoryMysql) getWhereForList(input *user_entity.ListUserParams) string {
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

func (r *UserRepositoryMysql) getOrderByForList(input *user_entity.ListUserParams) string {
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
