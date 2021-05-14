package model

import (
	"time"
)

// ComponentMeta  ='组件元数据表'
type ComponentMeta struct {
	ID            uint32    `gorm:"column:id" json:"id" `
	ComponentID   string    `gorm:"column:component_id" json:"component_id" `
	Name          string    `gorm:"column:name" json:"name" `
	ComponentType string    `gorm:"column:component_type" json:"component_type" `
	Description   string    `gorm:"column:description" json:"description" `
	Config        string    `gorm:"column:config" json:"config" `
	DataItemID    string    `gorm:"column:data_item_id" json:"data_item_id" `
	SampleInfoID  string    `gorm:"column:sample_info_id" json:"sample_info_id" `
	CreateTime    time.Time `gorm:"column:create_time" json:"create_time" `
	UpdateTime    time.Time `gorm:"column:update_time" json:"update_time" `
	DeleteTime    time.Time `gorm:"column:delete_time" json:"delete_time" `
}
