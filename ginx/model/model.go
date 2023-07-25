package model

import (
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:create_at" json:"create_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:update_at" json:"update_at"`
}
