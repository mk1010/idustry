package model

import (
	IDCModel "industry_identification_center/modules/industry_identification_center/model"
	"industry_identification_center/modules/redisclient"
)

var (
	dbs = []*DB{&DB{key: "industry_identification_center"}}
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
	err := IDCModel.Init()
	if err != nil {
		return err
	}
	return nil
}
