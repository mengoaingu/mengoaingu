package dbutils

import "github.com/go-sql-driver/mysql"

func IsMySQLDuplicatedError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr != nil && mysqlErr.Number == 1062 {
			// Do nothing
			return true
		}
	}
	return false
}

func IsMySQLOutOfRangeError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr != nil && mysqlErr.Number == 1264 {
			// Do nothing
			return true
		}
	}
	return false
}

func IsMySQLCheckConstraintError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr != nil && mysqlErr.Number == 3819 {
			// Do nothing
			return true
		}
	}
	return false
}

func IsMySQLDataTooLongError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr != nil && mysqlErr.Number == 1406 {
			// Do nothing
			return true
		}
	}
	return false
}
