package repository

import (
	"be-go-umkm/apps/domain"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepository interface {
	FindAll(ctx context.Context, db *gorm.DB) ([]domain.Account, error)
	FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Account, error)
	Create(ctx context.Context, db *gorm.DB, account domain.Account) (domain.Account, error)
	Update(ctx context.Context, db *gorm.DB, account domain.Account) (domain.Account, error)
	Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error
}
