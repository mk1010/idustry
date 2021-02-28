package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/mk1010/idustry/modules/industry_identification_center/model"
)

func TestModule(t *testing.T) {
	if err := initConfig(); err != nil {
		panic(err)
	}
	if err := initModel(); err != nil {
		panic(fmt.Sprintf("init model error:%v", err))
	}
	exams := make([]model.Exam, 0, 600)
	subjects := []string{"语文", "英语", "数学"}
	for i := 202; i <= 401; i++ {
		for _, subject := range subjects {
			exams = append(exams, model.Exam{
				StudentID: uint32(i),
				Subject:   subject,
				Grade:     uint8(rand.Uint32() % 101),
			})
		}
	}
	if model.DeviceInfoDB.WriteDB().Table("exam").Create(&exams).Error != nil {
		panic("mk")
	}
}
