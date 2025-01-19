package entities

import (
	"time"

	"gorm.io/datatypes"
)

type ActivityLog struct {
	ID                uint64         `gorm:"primaryKey;autoIncrement"`
	ActivityLogTypeID uint           `gorm:"not null;index"`
	UserID            uint64         `gorm:"not null;index"`
	BranchID          uint           `gorm:"not null;index"`
	Description       string         `gorm:"type:varchar(512);not null"`
	Metadata          datatypes.JSON `gorm:"type:json;not null"`
	CreatedAt         time.Time      `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt         *time.Time     `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
