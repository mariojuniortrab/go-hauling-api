package user_mysql_repository

import (
	"database/sql"
)

const TableName = "users"

type UserRepositoryMysql struct {
	DB *sql.DB
}
