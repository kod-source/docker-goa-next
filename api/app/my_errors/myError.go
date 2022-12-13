package myerrors

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

// ErrPasswordWorng パスワードが間違っているときのエラー
var ErrPasswordWorng = errors.New("password is wronging")
var ErrEmptyString = errors.New("title is empty string")
var ErrBadRequestSting = errors.New("request is empty string")
// ErrBadRequestInt 数字の0が来たときにエラー
var ErrBadRequestInt = errors.New("request is empty int")

var MySQLErrorDuplicate = &mysql.MySQLError{Number: 1062, Message: "duplicate entry"}

func GetMySQLErrorNumber(err error) uint16 {
	var myErr *mysql.MySQLError
	if ok := errors.As(err, &myErr); ok {
		return myErr.Number
	}
	return 0
}
