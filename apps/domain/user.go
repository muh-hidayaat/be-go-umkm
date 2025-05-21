package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID              uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id"`
	Name            string     `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Email           string     `gorm:"column:email;type:varchar(255);unique;not null" json:"email"`
	EmailVerifiedAt *time.Time `gorm:"column:email_verified_at;type:timestamp; null" json:"email_verified_at"`
	Password        string     `gorm:"column:password;type:varchar(255);not null" json:"-"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
