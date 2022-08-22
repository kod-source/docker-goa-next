package myerrors

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// PasswordWorngError パスワードが間違っているときのエラー
var PasswordWorngError = errors.New("Password is wronging")
var EmptyStringError = errors.New("Title is empty string")
var BadRequestStingError = errors.New("Request is empty string")
// BadRequestIntError 数字の0が来たときにエラー
var BadRequestIntError = errors.New("Request is empty int")

var MySQLErrorDuplicate = &mysql.MySQLError{Number: 1062, Message: "duplicate entry"}

func GetMySQLErrorNumber(err error) uint16 {
	var myErr *mysql.MySQLError
	if ok := errors.As(err, &myErr); ok {
		return myErr.Number
	}
	return 0
}
