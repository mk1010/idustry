package model

import (
	"time"
)

// AdaptorMeta  ='适配器元数据表'
type AdaptorMeta struct {
	ID          uint32    `gorm:"column:id" json:"id" `
	AdaptorID   string    `gorm:"column:adaptor_id" json:"adaptor_id" `
	Name        string    `gorm:"column:name" json:"name" `
	AdaptorType string    `gorm:"column:adaptor_type" json:"adaptor_type" `
	Description string    `gorm:"column:description" json:"description" `
	DeviceID    string    `gorm:"column:device_id" json:"device_id" `
	Config      string    `gorm:"column:config" json:"config" `
	CreateTime  time.Time `gorm:"column:create_time" json:"create_time" `
	UpdateTime  time.Time `gorm:"column:update_time" json:"update_time" `
	DeleteTime  time.Time `gorm:"column:delete_time" json:"delete_time" `
}
