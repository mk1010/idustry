package model

import (
	"industry_identification_center/config"

	"github.com/jinzhu/gorm"
)

const(
	dbTypeRead="read"
	dbTypeWrite="write"
)

type DB struct{
	key string
	write *gorm.DB
	read []*gorm.DB
}

func (d *DB) Init()bool{
	option,ok:=config.ConfInstance.DBConfigs[d.key]
	if !ok{
		return false
	}
	ok1:=d.

	return ok1&&ok2
}

func(d *DB)dbInstance(option config.DBConfig,dbType string)bool{

}