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
// ErrBadRequestEmptyArray 配列のデータが空の時
var ErrBadRequestEmptyArray = errors.New("request is empty array")
// ErrBadRequestNoPermission 閲覧権限がない時
var ErrBadRequestNoPermission = errors.New("request is no permission")

// MySQLErrorDuplicate ユニークインデックスのエラー
var MySQLErrorDuplicate = &mysql.MySQLError{Number: 1062, Message: "duplicate entry"}
// MySQLErrorAddOrUpdateForeignKey 外部キー製薬のエラー
var MySQLErrorAddOrUpdateForeignKey = &mysql.MySQLError{Number: 1452, Message: "cannot add or update a child row: a foreign key constraint fails"}

func GetMySQLErrorNumber(err error) uint16 {
	var myErr *mysql.MySQLError
	if ok := errors.As(err, &myErr); ok {
		return myErr.Number
	}
	return 0
}
