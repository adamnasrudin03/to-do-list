package entity

import "time"

type Todo struct {
	ID              uint64    `gorm:"primaryKey" json:"id"`
	ActivityGroupID uint64    `gorm:"not null" json:"activity_group_id"`
	Title           string    `gorm:"not null" json:"title"`
	IsActive        bool      `gorm:"not null;default:false" json:"is_active" `
	Priority        string    `gorm:"not null;default:'very-high'" json:"priority" `
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
