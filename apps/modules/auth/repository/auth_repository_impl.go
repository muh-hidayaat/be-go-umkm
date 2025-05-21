package repository

import (
	"be-go-umkm/apps/domain"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (r *UserRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.User, error) {
	var user domain.User
	err := db.WithContext(ctx).First(&user, "id = ?", id).Error
	return user, err
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, db *gorm.DB, email string) (domain.User, error) {
	var user domain.User
	err := db.WithContext(ctx).First(&user, "email = ?", email).Error
	return user, err
}

func (r *UserRepositoryImpl) Create(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error) {
	err := db.WithContext(ctx).Create(&user).Error
	return user, err
}

func (r *UserRepositoryImpl) Update(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error) {
	// Ambil data existing (opsional, kalau mau cek apakah user ada atau tidak dulu)
	var existingUser domain.User
	if err := db.WithContext(ctx).First(&existingUser, "id = ?", user.ID).Error; err != nil {
		return user, err // return langsung kalau gak ketemu
	}

	// Map hanya field yang mau di-update
	updateFields := map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
	}

	// Opsional: update password jika ada di request
	if user.Password != "" {
		updateFields["password"] = user.Password
	}

	err := db.WithContext(ctx).
		Model(&user).
		Where("id = ?", user.ID).
		Updates(updateFields).Error

	return user, err
}

func (r *UserRepositoryImpl) SaveVerificationCode(ctx context.Context, db *gorm.DB, token, email string) error {
	// Menyiapkan nilai untuk email_verified_at menjadi waktu saat ini
	now := time.Now()

	result := db.WithContext(ctx).
		Model(&domain.User{}).
		Where("email = ? AND verification_code = ?", email, token).
		Updates(map[string]interface{}{
			"verification_code": nil,
			"email_verified_at": now,
		})

	if result.RowsAffected == 0 {

		return fmt.Errorf("user with email %s and verification code %s not found or already verified", email, token)
	}

	return result.Error
}
