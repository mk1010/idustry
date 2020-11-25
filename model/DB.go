package model

import (
	"errors"
	"fmt"
	"industry_identification_center/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

const (
	dbTypeRead  = "read"
	dbTypeWrite = "write"
)

type DB struct {
	key string
	db  *gorm.DB
}

func (d *DB) Init() bool {
	option, ok := config.ConfInstance.DBConfigs[d.key]
	if !ok {
		return false
	}
	//必须先开write
	ok1 := d.dbInstance(option, dbTypeWrite)
	ok2 := d.dbInstance(option, dbTypeRead)
	return ok1 && ok2
}

func (d *DB) dbInstance(option config.DBConfig, dbType string) bool {
	if dbType == dbTypeRead {
		if d.db != nil {
			connStrs := make([]string, 0, len(option.ReadDB))
			for _, readDB := range option.ReadDB {
				connStr, err := d.gormConnect(readDB, &option)
				if err != nil {
					//log
					return false
				}
				connStrs = append(connStrs, connStr)
			}

			_, err := d.openRead(connStrs)
			if err != nil {
				//log
				return false
			}
		} else {
			//加载顺序错误，应该先加载写库
			return false
		}
	} else if dbType == dbTypeWrite {
		connStr, err := d.gormConnect(option.WriteDB, &option)
		if err != nil {
			//log
			return false
		}
		d.db, err = d.openWrite(connStr)
		if err != nil {
			//log
			return false
		}
	}
	return true
}

func (d *DB) gormConnect(dbInfo config.DBConnectInfo, option *config.DBConfig) (string, error) {
	if dbInfo.AuthKey == "" && dbInfo.DefaultHostPort == "" {
		return "", errors.New("数据库配置为空")
	}
	var connStr string
	if dbInfo.AuthKey != "" {
		connStr = fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", dbInfo.UserName, dbInfo.AuthKey, dbInfo.Consul, option.Database, option.Settings)
		return connStr, nil
	}
	connStr = fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", dbInfo.UserName, dbInfo.Password, dbInfo.DefaultHostPort, option.Database, option.Settings)
	return connStr, nil
}

func (d *DB) openWrite(connStr string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(connStr), nil)
}

func (d *DB) openRead(connStrs []string) (*gorm.DB, error) {
	dial := make([]gorm.Dialector, 0, len(connStrs))
	for _, str := range connStrs {
		dial = append(dial, mysql.Open(str))
	}

	d.db.Use(dbresolver.Register(dbresolver.Config{
		Replicas: dial,
	}))
	return d.db, nil
}
