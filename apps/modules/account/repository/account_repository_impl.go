package repository

import (
	"be-go-umkm/apps/domain"
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountRepositoryImpl struct{}

func NewAccountRepository() AccountRepository {
	return &AccountRepositoryImpl{}
}

func (r *AccountRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]domain.Account, error) {
	var account []domain.Account
	err := db.WithContext(ctx).Find(&account).Error
	return account, err
}

func (r *AccountRepositoryImpl) FindByID(ctx context.Context, db *gorm.DB, id uuid.UUID) (domain.Account, error) {
	var account domain.Account
	err := db.WithContext(ctx).First(&account, "id = ?", id).Error
	return account, err
}

func (r *AccountRepositoryImpl) Create(ctx context.Context, db *gorm.DB, account domain.Account) (domain.Account, error) {
	account.ID = uuid.New()
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	err := db.WithContext(ctx).Create(&account).Error
	return account, err
}

func (r *AccountRepositoryImpl) Update(ctx context.Context, db *gorm.DB, account domain.Account) (domain.Account, error) {
	account.UpdatedAt = time.Now()
	err := db.WithContext(ctx).Model(&domain.Account{}).Where("id = ?", account.ID).Updates(map[string]interface{}{
		"name":    account.Name,
		"type":    account.Type,
		"balance": account.Balance,
	}).Error
	return account, err
}

func (r *AccountRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, id uuid.UUID) error {
	err := db.WithContext(ctx).Delete(&domain.Account{}, id).Error
	return err
}
