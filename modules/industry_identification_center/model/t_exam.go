package model

import (
)
type Exam struct {
	ID          uint32  `gorm:"column:id" json:"id" `
	StudentID        uint32  `gorm:"column:student_id" json:"student_id" `
	Grade         uint8   `gorm:"column:grade" json:"grade" `
	Subject      string  `gorm:"column:subject" json:"subject" `
}
