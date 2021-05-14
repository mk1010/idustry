package model

import (
	"time"
)

// SamplingInfoMeta  ='采样信息元数据表'
type SamplingInfoMeta struct {
	ID               uint32    `gorm:"column:id" json:"id" `
	SamplingInfoID   string    `gorm:"column:sampling_info_id" json:"sampling_info_id" `
	SamplingInfoType string    `gorm:"column:sampling_info_type" json:"sampling_info_type" `
	SamplingPeriod   uint32    `gorm:"column:sampling_period" json:"sampling_period" `
	UploadPeriod     uint32    `gorm:"column:upload_period" json:"upload_period" `
	CreateTime       time.Time `gorm:"column:create_time" json:"create_time" `
	UpdateTime       time.Time `gorm:"column:update_time" json:"update_time" `
	DeleteTime       time.Time `gorm:"column:delete_time" json:"delete_time" `
}
