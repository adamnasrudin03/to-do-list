package entity

import "time"

type Activity struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Email     string    `gorm:"not null" json:"email" `
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
}
