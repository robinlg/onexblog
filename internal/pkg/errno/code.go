// Copyright 2025 Robin Liu <robinliu27@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/robinlg/onexblog. The professional
// version of this repository is https://github.com/robinlg/onexblog.

package errno

import (
	"net/http"

	"github.com/robinlg/onexlib/pkg/errorsx"
)

var (
	// OK 代表请求成功.
	OK = &errorsx.ErrorX{Code: http.StatusOK, Message: ""}

	// ErrInternal 表示所有未知的服务器端错误.
	ErrInternal = errorsx.ErrInternal

	// ErrInvalidArgument 表示参数验证失败.
	ErrInvalidArgument = errorsx.ErrInvalidArgument

	// ErrUnauthenticated 表示认证失败.
	ErrUnauthenticated = errorsx.ErrUnauthenticated

	// ErrPermissionDenied 表示请求没有权限.
	ErrPermissionDenied = errorsx.ErrPermissionDenied

	// ErrPageNotFound 表示页面未找到.
	ErrPageNotFound = &errorsx.ErrorX{Code: http.StatusNotFound, Reason: "NotFound.PageNotFound", Message: "Page not found."}

	// ErrSignToken 表示签发 JWT Token 时出错.
	ErrSignToken = &errorsx.ErrorX{Code: http.StatusUnauthorized, Reason: "Unauthenticated.SignToken", Message: "Error occurred while signing the JSON web token."}

	// ErrTokenInvalid 表示 JWT Token 格式无效.
	ErrTokenInvalid = &errorsx.ErrorX{Code: http.StatusUnauthorized, Reason: "Unauthenticated.TokenInvalid", Message: "Token was invalid."}

	// ErrDBRead 表示数据库读取失败.
	ErrDBRead = &errorsx.ErrorX{Code: http.StatusInternalServerError, Reason: "InternalError.DBRead", Message: "Database read failure."}

	// ErrDBWrite 表示数据库写入失败.
	ErrDBWrite = &errorsx.ErrorX{Code: http.StatusInternalServerError, Reason: "InternalError.DBWrite", Message: "Database write failure."}
)
