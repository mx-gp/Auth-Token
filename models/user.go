package models

import "time"

// User DB Model
type User struct {
	ID        int       `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"id"`
	Token     string    `gorm:"type:varchar(100);not null" json:"token"`
	IPAddress string    `gorm:"type:varchar(100);not null" json:"ipAddress"`
	CreatedAt time.Time `gorm:"type:timestamp autoCreateTime:true" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp autoUpdateTime:true" json:"updatedAt"`
}
