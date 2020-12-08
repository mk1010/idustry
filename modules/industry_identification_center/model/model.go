package model

import "fmt"

var (
	DeviceInfoDB = NewDB("industry_identification_center")
)

func Init() error {
	ok := DeviceInfoDB.Init()
	if !ok {
		return fmt.Errorf("初始化数据库失败：db_key=%v。", DeviceInfoDB.key)
	}
	return nil
}
