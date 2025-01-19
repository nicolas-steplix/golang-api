package entities

import "time"

type Branch struct {
	ID          uint       `gorm:"primaryKey;autoIncrement"`
	Code        string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	Description string     `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
