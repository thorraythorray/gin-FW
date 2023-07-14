package gorm

import (
	"time"
)

type BaseGorm struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:create_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;column:update_at"`
}
