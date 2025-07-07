package model

import (
	"gorm.io/gorm"
)

// 定义数据模型
type Task struct {
	gorm.Model
	Uid       uint   `gorm:"not null"`
	Title     string `json:"title"`
	Status    int    `gorm:"default:0"`
	Content   string `gorm:"type:longtext"`
	StartTime int64
	EndTime   int64
}
