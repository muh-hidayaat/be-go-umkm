package domain

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type Account struct {
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	UserID    uuid.UUID `gorm:"type:char(36);not null" json:"user_id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Type      string    `gorm:"type:varchar(255);not null" json:"type"`
	Balance   int       `gorm:"type:int;not null" json:"balance"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (model *Account) BeforeCreate(scope *gorm.DB) error {
	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()
	model.ID = uuid.New()
	return nil
}

func (model *Account) BeforeUpdate(tx *gorm.DB) error {
	model.UpdatedAt = time.Now()
	return nil
}
