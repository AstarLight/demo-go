package main

import (

	"github.com/marmotedu/errors"
  )

const (
	ErrSuccess = 200

	// 未知的服务器错误
	ErrUnknown = 1000

	//请求参数结构体绑定错误
	ErrBind = 1001
	//验证错误
	ErrValidation = 1002
	// token非法
	ErrTokenInvalid = 1003
	// 资源不存在
	ErrPageNotFound = 1004
	//json编码失败
	ErrEncodingJSON = 1005
	// json解码失败
	ErrDecodingJSON = 1006
	// 数据库错误
	ErrDatabase = 1007
	// 用户已存在
	ErrUserAlreadyExist = 1008
	// 密码不正确
	ErrPasswordIncorrect = 1009
	// 签名非法
	ErrSignatureInvalid = 1010
	// 没有权限
	ErrPermissionDenied = 1011
	// 请求过多
	ErrTooManyRequest = 1012
)

type defaultCoder struct {
	// C refers to the integer code of the ErrCode.
	C int

	// HTTP status that should be used for the associated error code.
	HTTP int

	// External (user) facing error text.
	Ext string

	// Ref specify the reference document.
	Ref string
}

// Code returns the integer code of the coder.
func (coder defaultCoder) Code() int {
	return coder.C

}

// String implements stringer. String returns the external error message,
// if any.
func (coder defaultCoder) String() string {
	return coder.Ext
}

// HTTPStatus returns the associated HTTP status code, if any. Otherwise,
// returns 200.
func (coder defaultCoder) HTTPStatus() int {
	if coder.HTTP == 0 {
		return 500
	}

	return coder.HTTP
}

// Reference returns the reference document.
func (coder defaultCoder) Reference() string {
	return coder.Ref
}


// 注册，将业务错误码映射为http status code
func init() {
	errors.Register(defaultCoder{ErrSuccess, 200, "OK", ""})
	errors.Register(defaultCoder{ErrUserAlreadyExist, 400, "User already exist", ""})
	errors.Register(defaultCoder{ErrBind, 400, "Error occurred while binding the request body to the struct", ""})
	errors.Register(defaultCoder{ErrPasswordIncorrect, 401, "Password was incorrect", ""})
	errors.Register(defaultCoder{ErrSignatureInvalid, 401, "Signature is invalid", ""})
	errors.Register(defaultCoder{ErrTokenInvalid, 401, "Token invalid", ""})
	errors.Register(defaultCoder{ErrTokenInvalid, 403, "Permission denied", ""})
	errors.Register(defaultCoder{ErrPageNotFound, 404, "Page not found", ""})
	errors.Register(defaultCoder{ErrTooManyRequest, 429, "Too Many Requests", ""})
	errors.Register(defaultCoder{ErrEncodingJSON, 500, "JSON data could not be encoded", ""})
	errors.Register(defaultCoder{ErrDecodingJSON, 500, "JSON data could not be decoded", ""})
	errors.Register(defaultCoder{ErrDatabase, 500, "Database error", ""})
	errors.Register(defaultCoder{ErrUnknown, 500, "Internal server error", ""})

}