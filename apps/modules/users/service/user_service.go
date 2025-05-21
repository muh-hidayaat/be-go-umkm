package service

import (
	"be-go-umkm/apps/domain"
	"be-go-umkm/apps/modules/users/repository"
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.User, error)
	// FindByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
	ChangePassword(ctx context.Context, id uuid.UUID, oldPassword, newPassword, confirmPassword string) error
}

type UserServiceImpl struct {
	repo repository.UserRepository
	db   *gorm.DB
}

func NewUserService(repo repository.UserRepository, db *gorm.DB) UserService {
	return &UserServiceImpl{repo: repo, db: db}
}

// HashPassword helper
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *UserServiceImpl) FindAll(ctx context.Context) ([]domain.User, error) {
	return s.repo.FindAll(ctx, s.db)
}

func (s *UserServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return s.repo.FindByID(ctx, s.db, id)
}

// func (s *UserServiceImpl) FindByEmail(ctx context.Context, email string) (domain.User, error) {
// 	return s.repo.FindByEmail(ctx, s.db, email)
// }

func (s *UserServiceImpl) Create(ctx context.Context, user domain.User) (domain.User, error) {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = hashedPassword
	return s.repo.Create(ctx, s.db, user)
}

func (s *UserServiceImpl) Update(ctx context.Context, user domain.User) (domain.User, error) {
	// Hash password hanya jika diisi baru
	if user.Password != "" {
		hashedPassword, err := hashPassword(user.Password)
		if err != nil {
			return user, err
		}
		user.Password = hashedPassword
	} else {
		existingUser, err := s.repo.FindByID(ctx, s.db, user.ID)
		if err != nil {
			return user, err
		}
		user.Password = existingUser.Password // keep old password if not provided
	}
	return s.repo.Update(ctx, s.db, user)
}

func comparePassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}

func (s *UserServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, s.db, id)
}

func (s *UserServiceImpl) ChangePassword(ctx context.Context, id uuid.UUID, oldPassword, newPassword, confirmPassword string) error {
	// Ensure newPassword matches confirmPassword
	if newPassword != confirmPassword {
		return fmt.Errorf("new password and confirm password do not match")
	}

	// Fetch the user by ID
	user, err := s.repo.FindByID(ctx, s.db, id)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	// Compare the old password
	if err := comparePassword(user.Password, oldPassword); err != nil {
		return fmt.Errorf("incorrect old password: %w", err)
	}

	// Hash the new password
	hashedPassword, err := hashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("failed to hash new password: %w", err)
	}

	// Update the user's password
	user.Password = hashedPassword
	_, err = s.repo.Update(ctx, s.db, user)
	if err != nil {
		return fmt.Errorf("failed to update password: %w", err)
	}

	return nil
}
