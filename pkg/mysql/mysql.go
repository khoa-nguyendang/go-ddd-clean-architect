package mysql

import (
	config "app/core/configs"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

// MySQL config
type MysqlConfig struct {
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	MysqlDbname   string
	MysqlDriver   string
}

func NewConfig(cfg *config.Config) (MysqlConfig, error) {
	if cfg == nil {
		return MysqlConfig{}, errors.New("invalid configs")
	}
	c := MysqlConfig{
		MysqlHost:     cfg.Mysql.MysqlHost,
		MysqlPort:     cfg.Mysql.MysqlPort,
		MysqlUser:     cfg.Mysql.MysqlUser,
		MysqlPassword: cfg.Mysql.MysqlPassword,
		MysqlDbname:   cfg.Mysql.MysqlDbname,
		MysqlDriver:   cfg.Mysql.MysqlDriver,
	}
	return c, nil
}

// Return new Mysql db instance
func NewMysqlDB(c MysqlConfig) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
		c.MysqlUser,
		c.MysqlPassword,
		c.MysqlHost,
		c.MysqlPort,
		c.MysqlDbname,
	)
	log.Default().Printf("mysql datasoure: %v", dataSourceName)
	db, err := sqlx.Connect(c.MysqlDriver, dataSourceName)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func RunMigration(c *config.Config) error {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true",
		c.Mysql.MysqlUser,
		c.Mysql.MysqlPassword,
		c.Mysql.MysqlHost,
		c.Mysql.MysqlPort,
		c.Mysql.MysqlDbname,
	)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Default().Printf("RunMigration.error1: %v", err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Default().Printf("RunMigration.error2: %v", err)
	}
	log.Default().Printf("RunMigration: %v", dataSourceName)
	m, err := migrate.NewWithDatabaseInstance("file://migrations", "mysql", driver)
	if err != nil {
		log.Default().Printf("RunMigration.error3: %v", err)
		return err
	}
	m.Up()

	return err
}
