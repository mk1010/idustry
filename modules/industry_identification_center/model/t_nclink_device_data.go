package model

import (
	"time"
)

// NclinkDeviceData  ='采样信息元数据表'
type NclinkDeviceData struct {
	ID          uint64    `gorm:"column:id" json:"id" `
	DataID      string    `gorm:"column:data_id" json:"data_id" `
	AdaptorID   string    `gorm:"column:adaptor_id" json:"adaptor_id" `
	DeviceID    string    `gorm:"column:device_id" json:"device_id" `
	ComponentID string    `gorm:"column:component_id" json:"component_id" `
	DataItemID  string    `gorm:"column:data_item_id" json:"data_item_id" `
	Payload     string    `gorm:"column:payload" json:"payload" `
	AdaptorTime int64     `gorm:"column:adaptor_time" json:"adaptor_time" `
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time" `
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time" `
	DeleteTime  time.Time `gorm:"column:delete_time" json:"delete_time" `
}
