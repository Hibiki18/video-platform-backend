package models

import "time"

// Video 表结构
type Video struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username    string    `gorm:"not null;index" json:"username"`   // 上传者用户名
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	FilePath    string    `gorm:"type:varchar(255);not null" json:"file_path"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}