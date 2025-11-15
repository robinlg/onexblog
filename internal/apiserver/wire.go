// Copyright 2025 Robin Liu <robinliu27@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/robinlg/onexblog. The professional
// version of this repository is https://github.com/robinlg/onexblog.

//go:build wireinject
// +build wireinject

package apiserver

import (
	"github.com/google/wire"
	auth "github.com/robinlg/onexlib/pkg/authz"

	"github.com/robinlg/onexblog/internal/apiserver/biz"
	"github.com/robinlg/onexblog/internal/apiserver/pkg/validation"
	"github.com/robinlg/onexblog/internal/apiserver/store"
	ginmw "github.com/robinlg/onexblog/internal/pkg/middleware/gin"
	grpcmw "github.com/robinlg/onexblog/internal/pkg/middleware/grpc"

	"github.com/robinlg/onexblog/internal/pkg/server"
)

func InitializeWebServer(*Config) (server.Server, error) {
	wire.Build(
		wire.NewSet(NewWebServer, wire.FieldsOf(new(*Config), "ServerMode")),
		wire.Struct(new(ServerConfig), "*"), // * 表示注入全部字段
		wire.NewSet(store.ProviderSet, biz.ProviderSet),
		ProvideDB, // 提供数据库实例
		validation.ProviderSet,
		wire.NewSet(
			wire.Struct(new(UserRetriever), "*"),
			wire.Bind(new(ginmw.UserRetriever), new(*UserRetriever)),
			wire.Bind(new(grpcmw.UserRetriever), new(*UserRetriever)),
		),
		auth.ProviderSet,
	)

	return nil, nil
}
