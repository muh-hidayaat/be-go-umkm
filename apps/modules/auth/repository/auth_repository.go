package repository

import (
	"be-go-umkm/apps/domain"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.User, error)
	FindByEmail(ctx context.Context, db *gorm.DB, email string) (domain.User, error)
	Create(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error)
	Update(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error)
	SaveVerificationCode(ctx context.Context, db *gorm.DB, token, email string) error
}
