package repository

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

func isForeignKeyError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == 1452
	}
	return false
}
