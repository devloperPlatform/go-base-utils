package db

import (
	"fmt"
	"github.com/devloperPlatform/go-base-utils/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
	"time"
)

var dbInfo *dbInfos

func init() {
	dbUserName := env.GetEnvOrDefault("DB_USER_NAME", "root")
	dbPassword := env.GetEnvOrDefault("DB_PASSWORD", "123456")
	dbHost := env.GetEnvOrDefault("DB_HOST", "127.0.0.1")
	dbPort := env.GetEnvOrDefault("DB_PORT", "3306")
	dbName := env.GetEnvOrDefault("DB_NAME", "test")
	dbInfo = &dbInfos{
		host:     dbHost,
		port:     dbPort,
		dbName:   dbName,
		username: dbUserName,
		password: dbPassword,
		debug:    false,
	}
}

func DB() *dbInfos {
	return dbInfo
}

type dbInfos struct {
	sync.Mutex
	host     string
	port     string
	username string
	password string
	dbName   string
	db       *gorm.DB
	debug    bool
}

func (d *dbInfos) SetHost(host string) *dbInfos {
	d.host = host
	return d
}

func (d *dbInfos) SetPort(port string) *dbInfos {
	d.port = port
	return d
}

func (d *dbInfos) SetUserName(userName string) *dbInfos {
	d.username = userName
	return d
}

func (d *dbInfos) SetPassword(password string) *dbInfos {
	d.password = password
	return d
}

func (d *dbInfos) SetDbName(dbName string) *dbInfos {
	d.dbName = dbName
	return d
}

func (d *dbInfos) StartDebug() *dbInfos {
	d.debug = true
	return d
}

func (d *dbInfos) GetDb() *gorm.DB {
	d.Lock()
	defer d.Unlock()
	if d.db == nil {
		d.db = d.initMySqlDb(d.host, d.port, d.username, d.password, d.dbName)
	}
	if d.debug {
		return d.db.Debug()
	}
	return d.db
}

func (d *dbInfos) initMySqlDb(host, port, username, password, dbname string) *gorm.DB {
	mysqlOpen, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=30s", username, password, host, port, dbname))
	if err != nil {
		panic(err)
	}
	mysqlOpen.DB().SetConnMaxLifetime(59 * time.Second)
	return mysqlOpen
}

func (d *dbInfos) MySqlWithTx(fn func(tx *gorm.DB)) {
	tx := d.GetDb().Begin()
	defer func() {
		err := recover()
		if err != nil {
			tx.Rollback()
			panic(err)
		}
		tx.Commit()
	}()

	fn(tx)
}
