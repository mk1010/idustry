package model

import (
	IDCModel "github.com/mk1010/idustry/modules/industry_identification_center/model"
	"github.com/mk1010/idustry/modules/redisclient"
)

var dbs = []*DB{&DB{key: "industry_identification_center"}}

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
