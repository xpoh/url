package storage

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"url/internal/config"
	"url/pkg/logger"
)

type MySqlStorage struct {
	db     *sql.DB
	Cfg    *config.Cfg
	Logger *zap.Logger
}

func NewMysqlStorage() *MySqlStorage {
	var err error
	var m = MySqlStorage{}

	m.Cfg = config.GetConfig()
	m.Logger = logger.GetLogger()

	m.db, err = sql.Open(m.Cfg.Data.Database.Driver, m.Cfg.Data.Database.Source)
	if err != nil {
		m.Logger.Panic("Error connect to database ", zap.Error(err))
	}
	m.Logger.Info("Connected to database", zap.String("database", m.Cfg.Data.Database.Source))
	return &m
}

func (m *MySqlStorage) Close() {
	m.db.Close()
	m.Logger.Info("Closed connection to database")
}

func (m *MySqlStorage) GetFullUrlByButty(buttyUrl string) (string, error) {
	var fullUrl string

	err := m.db.QueryRow("SELECT url FROM urls WHERE id=?", buttyUrl).Scan(&fullUrl)
	if err != nil {
		m.Logger.Info("GetFullUrlByButty error", zap.Error(err))
		return "", err
	}
	return fullUrl, nil
}

func (m *MySqlStorage) NewButtyUrl(url string) (string, error) {
	// TODO Сделать проверку от дублирования конечных ссылок

	var buttyUrl string

	tx, err := m.db.Begin()
	if err != nil {
		m.Logger.Error("NewButtyUrl begin TX error", zap.Error(err))
		return "", err
	}
	result, err1 := m.db.Exec("INSERT INTO urls VALUES (0, '', NOW())")

	if err1 != nil {
		err1 = tx.Rollback()
		m.Logger.Info("NewButtyUrl insert error", zap.Error(err))
		return "", err1
	}
	id, err2 := result.LastInsertId()
	if err2 != nil {
		err2 = tx.Rollback()
		m.Logger.Info("NewButtyUrl last_id error", zap.Error(err))
		return "", err2
	}
	buttyUrl = idToUrl(id)

	_, err = m.db.Exec("UPDATE urls SET btUrl='?', url='?' WHERE id=?", buttyUrl, url, id)
	if err != nil {
		err = tx.Rollback()
		m.Logger.Info("NewButtyUrl update error", zap.Error(err))
		return "", err
	}

	err = tx.Commit()
	if err != nil {
		m.Logger.Error("NewButtyUrl Commit TX error", zap.Error(err))
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
