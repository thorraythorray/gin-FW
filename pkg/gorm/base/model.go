package base

import (
	"time"
)

type BaseModel struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:create_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:update_at"`
}
