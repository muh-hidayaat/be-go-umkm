package service

import (
	"be-go-umkm/apps/domain"
	"be-go-umkm/apps/helpers"
	"be-go-umkm/apps/modules/auth/repository"
	"context"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	FindByID(ctx context.Context, id uuid.UUID) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
	GenerateAuthToken(user domain.User) (string, error) // Generate JWT token
	ComparePassword(hashedPassword, plainPassword string) error
}

type UserServiceImpl struct {
	repo repository.UserRepository
	// repoCustomer customer.CustomerRepository
	db *gorm.DB
}

func NewUserService(repo repository.UserRepository, db *gorm.DB) UserService {
	return &UserServiceImpl{
		repo: repo,
		db:   db,
		// repoCustomer: repoCustomer,
	}
}

// HashPassword helper
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *UserServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	return s.repo.FindByID(ctx, s.db, id)
}

func (s *UserServiceImpl) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.repo.FindByEmail(ctx, s.db, email)
}

func (s *UserServiceImpl) Create(ctx context.Context, user domain.User) (domain.User, error) {
	// Generate OTP for user verification
	// otp := helpers.GenerateRandomOTP()

	// Hash the password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = hashedPassword
	// user.VerificationCode = &otp

	// Start a transaction
	tx := s.db.WithContext(ctx).Begin()

	// Ensure to commit or rollback the transaction
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Try to create the user in the database
	result, err := s.repo.Create(ctx, tx, user)
	if err != nil {
		tx.Rollback() // Rollback the transaction in case of an error
		return result, err
	}

	if err := tx.Commit().Error; err != nil {
		return result, err
	}

	return result, nil
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

func (s *UserServiceImpl) GenerateAuthToken(user domain.User) (string, error) {
	return helpers.GenerateJWT(user.ID)
}

func (s *UserServiceImpl) ComparePassword(hashedPassword, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
}
