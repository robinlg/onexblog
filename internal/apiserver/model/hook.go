// Copyright 2025 Robin Liu <robinliu27@163.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/robinlg/onexblog. The professional
// version of this repository is https://github.com/robinlg/onexblog.

package model

import (
	"time"

	"gorm.io/gorm"

	"github.com/robinlg/onexblog/internal/pkg/rid"
	"github.com/robinlg/onexblog/pkg/auth"
)

// AfterCreate 在创建数据库记录之后生成 postID.
func (m *PostM) AfterCreate(tx *gorm.DB) error {
	m.PostID = rid.PostID.New(uint64(m.ID))

	return tx.Save(m).Error
}

// BeforeCreate 在创建数据库记录之前加密明文密码.
func (m *UserM) BeforeCreate(tx *gorm.DB) error {
	// Encrypt the user password.
	var err error
	m.Password, err = auth.Encrypt(m.Password)
	if err != nil {
		return err
	}
	now := time.Now()
	m.CreatedAt = now
	m.UpdatedAt = now

	return nil
}

// AfterCreate 在创建数据库记录之后生成 userID.
func (m *UserM) AfterCreate(tx *gorm.DB) error {
	m.UserID = rid.UserID.New(uint64(m.ID))

	return tx.Save(m).Error
}

// BeforeCreate 更新创建时间.
func (m *PostM) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	m.CreatedAt = now
	m.UpdatedAt = now

	return nil
}
