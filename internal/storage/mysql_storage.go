// Copyright 2022. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package storage implement interface and functions
// for storage
package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"url/internal/config"
	"url/pkg/logger"
)

// MySqlStorage struct for Mysql and MariaDB storage connection
type MySqlStorage struct {
	db     *sql.DB
	Cfg    *config.Cfg
	Logger *zap.Logger
}

// NewMysqlStorage is a constructor for mysql storage
func NewMysqlStorage() *MySqlStorage {
	var err error
	var m = &MySqlStorage{}

	m.Cfg = config.GetConfig()
	m.Logger = logger.GetLogger()
	m.db, err = sql.Open(m.Cfg.Data.Database.Driver, m.Cfg.Data.Database.Source)
	m.Logger.Debug("NewMysqlStorage", zap.Any("m.db", m.db))
	if err != nil {
		m.Logger.Panic("Error connect to database ", zap.Error(err))
	}
	m.Logger.Info("Connected to database", zap.String("database", m.Cfg.Data.Database.Source))
	return m
}

// Close graceful shutdown connection
func (m *MySqlStorage) Close() {
	m.db.Close()
	m.Logger.Info("Closed connection to database")
}

func (m *MySqlStorage) GetFullUrlByButty(buttyUrl string) (string, error) {
	var fullUrl string

	err := m.db.QueryRow("SELECT url FROM urls WHERE btUrl=?", buttyUrl).Scan(&fullUrl)
	if err != nil {
		m.Logger.Info("GetFullUrlByButty error", zap.String("error", err.Error()))
		return "", err
	}
	return fullUrl, nil
}

func (m *MySqlStorage) NewButtyUrl(url string) (string, error) {
	// TODO Сделать проверку от дублирования конечных ссылок

	var buttyUrl string
	var err error

	var tx *sql.Tx
	tx, err = m.db.Begin()
	if err != nil {
		m.Logger.Error("NewButtyUrl begin TX error", zap.String("error", err.Error()))
		return "", err
	}

	var result sql.Result
	result, err = tx.Exec("INSERT INTO urls VALUES (0, '', '', NOW())")

	if err != nil {
		m.Logger.Info("NewButtyUrl insert error", zap.String("error", err.Error()))
		_ = tx.Rollback()
		return "", err
	}
	var id int64
	id, err = result.LastInsertId()
	if err != nil {
		m.Logger.Info("NewButtyUrl last_id error", zap.String("error", err.Error()))
		_ = tx.Rollback()
		return "", err
	}
	buttyUrl = idToUrl(id)

	_, err = tx.Exec("UPDATE urls SET btUrl=?, url=? WHERE id=?", buttyUrl, url, id)
	if err != nil {
		m.Logger.Info("NewButtyUrl update error", zap.String("error", err.Error()))
		_ = tx.Rollback()
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		m.Logger.Error("NewButtyUrl Commit TX error", zap.String("error", err.Error()))
		return "", err
	}

	return buttyUrl, nil
}

func (m *MySqlStorage) ClearOldLinks(days int) error {
	_, err := m.db.Exec("DELETE FROM urls WHERE DATEDIFF(NOW(), createDate)>?", m.Cfg.Service.LinksLivesDays)
	if err != nil {
		m.Logger.Info("ClearOldLinks error", zap.Error(err))
	}
	return err
}
