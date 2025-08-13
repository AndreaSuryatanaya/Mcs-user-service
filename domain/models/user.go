package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	UUID        uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"varchar(100);not null"`
	Pass        string    `gorm:"varchar(255);not null"`
	PhoneNumber string    `gorm:"varchar(15);not null"`
	Email       string    `gorm:"varchar(100);not null"`
	RoleID      uint      `gorm:"type:uint;not null"`
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	Role        Role `gorm:"foreignKey:role_id;reference:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
