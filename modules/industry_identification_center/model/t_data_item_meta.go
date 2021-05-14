package model

import (
	"time"
)

// DataItemMeta  ='数据项元数据表'
type DataItemMeta struct {
	ID           uint32    `gorm:"column:id" json:"id" `
	DataItemID   string    `gorm:"column:data_item_id" json:"data_item_id" `
	Name         string    `gorm:"column:name" json:"name" `
	DataItemType string    `gorm:"column:data_item_type" json:"data_item_type" `
	Description  string    `gorm:"column:description" json:"description" `
	Items        string    `gorm:"column:items" json:"items" `
	DataUnit     string    `gorm:"column:data_unit" json:"data_unit" `
	CreateTime   time.Time `gorm:"column:create_time" json:"create_time" `
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time" `
	DeleteTime   time.Time `gorm:"column:delete_time" json:"delete_time" `
}
