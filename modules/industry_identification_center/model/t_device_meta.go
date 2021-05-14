package model

import (
	"time"
)

// DeviceMeta  ='设备元数据表'
type DeviceMeta struct {
	ID          uint32    `gorm:"column:id" json:"id" `
	DeviceID    string    `gorm:"column:device_id" json:"device_id" `
	Name        string    `gorm:"column:name" json:"name" `
	DeviceType  string    `gorm:"column:device_type" json:"device_type" `
	Description string    `gorm:"column:description" json:"description" `
	DeviceGroup string    `gorm:"column:device_group" json:"device_group" `
	ComponentID string    `gorm:"column:component_id" json:"component_id" `
	Config      string    `gorm:"column:config" json:"config" `
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time" `
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time" `
	DeleteTime  time.Time `gorm:"column:delete_time" json:"delete_time" `
}
