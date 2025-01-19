package entities

import "time"

type User struct {
	ID        uint64     `gorm:"primaryKey;autoIncrement"`
	Email     string     `gorm:"type:varchar(255);uniqueIndex;not null"`
	Password  string     `gorm:"type:varchar(255);not null"`
	FirstName string     `gorm:"type:varchar(255);not null"`
	LastName  string     `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time  `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt *time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP"`
}
