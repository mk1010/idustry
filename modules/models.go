package modules

import (
	IDCModel "github.com/mk1010/idustry/modules/industry_identification_center/model"
	"github.com/mk1010/idustry/modules/redisclient"
	"github.com/mk1010/idustry/modules/rmq"
)

// var dbs = []*DB{{key: "industry_identification_center"}}

func Init() error {
	if err := redisclient.InitRedisClient(); err != nil {
		return err
	}
	if err := initDBClient(); err != nil {
		return err
	}
	if err := rmq.Init(); err != nil {
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
