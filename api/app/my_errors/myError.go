package myerrors

import "errors"

// PasswordWorngError パスワードが間違っているときのエラー
var PasswordWorngError = errors.New("Password is wronging")
