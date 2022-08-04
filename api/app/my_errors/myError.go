package myerrors

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// PasswordWorngError パスワードが間違っているときのエラー
var PasswordWorngError = errors.New("Password is wronging")

var MySQLErrorDuplicate = &mysql.MySQLError{Number: 1062, Message: "duplicate entry"}

func GetMySQLErrorNumber(err error) uint16 {
	var myErr *mysql.MySQLError
	if ok := errors.As(err, &myErr); ok {
		return myErr.Number
	}
	return 0
}
