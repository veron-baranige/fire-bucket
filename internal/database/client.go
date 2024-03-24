package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/veron-baranige/fire-bucket/internal/config"
)

var (
	Client *Queries
	conn   *sql.DB
)

func SetupClient() error {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Get(config.DbUser),
		config.Get(config.DbPassword),
		config.Get(config.DbHost),
		config.Get(config.DbPort),
		config.Get(config.DbName),
	)

	dbConn, err := sql.Open(config.Get(config.DbDriver), connStr)
	if err != nil {
		return err
	}

	Client = New(dbConn)
	conn = dbConn

	return nil
}

func GetDbConn() *sql.DB {
	return conn
}