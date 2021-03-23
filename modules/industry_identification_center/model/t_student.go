//# Version:1.0.6
package model

type Student struct {
	ID          uint32 `gorm:"column:id" json:"id" `
	Name        string `gorm:"column:name" json:"name" `
	Age         uint8  `gorm:"column:age" json:"age" `
	Region      string `gorm:"column:region" json:"region" `
	PhoneNumber string `gorm:"column:phone_number" json:"phone_number,omitempty" `
}
