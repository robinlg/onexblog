// Copyright 2025 Robin Liu <robinliu27@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/robinlg/onexblog. The professional
// version of this repository is https://github.com/robinlg/onexblog.

package grpc

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/robinlg/onexblog/internal/pkg/contextx"
	"github.com/robinlg/onexblog/internal/pkg/known"
	"github.com/robinlg/onexblog/internal/pkg/log"
)

// AuthnBypasswInterceptor 是一个 gRPC 拦截器，模拟所有请求都通过认证。
func AuthnBypasswInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		// 从请求头中获取用户ID
		userID := "user-000001" // 默认用户ID
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			// 获取header中指定的用户ID，假设header名为"x-user-id"
			if values := md.Get(known.XUserID); len(values) > 0 {
				userID = values[0]
			}
		}

		log.Debugw("Simulated authentication successful", "userID", userID)

		// 将默认的用户信息存入上下文
		//nolint: staticcheck
		ctx = context.WithValue(ctx, known.XUserID, userID)

		// 为 log 和 contextx 提供用户上下文支持
		ctx = contextx.WithUserID(ctx, userID)

		// 继续处理请求
		return handler(ctx, req)
	}
}
