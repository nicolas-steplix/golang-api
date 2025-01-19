package entities

type ActivityLogType struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Description string `gorm:"type:varchar(255);not null"`
}
