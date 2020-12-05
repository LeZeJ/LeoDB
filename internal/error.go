package internal

import "errors"

//自定义错误类型
var (
	ErrorNotFound = errors.New("Not Found ")
	ErrDeletion = errors.New("Type Deletion")

)
