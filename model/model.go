package model

import (
	"errors"
	"fmt"
	"industry_identification_center/model/redisclient"
)

var (
	DefaultDB = &DB{key: "industry_identification_center"}
	dbs       = []*DB{DefaultDB}
)

func Init() error {
	if err := redisclient.InitRedisClient(); err != nil {
		return err
	}
	if err := initDBClient(); err != nil {
		return err
	}

	return nil
}

func initDBClient() error {
	for _, db := range dbs {
		ok := db.Init()
		if !ok {
			return errors.New(fmt.Sprintf("初始化数据库失败:db_key=%v", db.key))
		}
	}
	return nil
}
