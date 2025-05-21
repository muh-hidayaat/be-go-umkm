package service

import (
	"be-go-umkm/apps/domain"
	"be-go-umkm/apps/modules/account/repository"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountService interface {
	FindAll(ctx context.Context) ([]domain.Account, error)
	FindByID(ctx context.Context, id uuid.UUID) (domain.Account, error)
	Create(ctx context.Context, account domain.Account) (domain.Account, error)
	Update(ctx context.Context, account domain.Account) (domain.Account, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type AccountServiceImpl struct {
	repo repository.AccountRepository
	db   *gorm.DB
}

func NewAccountService(repo repository.AccountRepository, db *gorm.DB) AccountService {
	return &AccountServiceImpl{repo: repo, db: db}
}

func (s *AccountServiceImpl) FindAll(ctx context.Context) ([]domain.Account, error) {
	return s.repo.FindAll(ctx, s.db)
}

func (s *AccountServiceImpl) FindByID(ctx context.Context, id uuid.UUID) (domain.Account, error) {
	return s.repo.FindByID(ctx, s.db, id)
}

func (s *AccountServiceImpl) Create(ctx context.Context, account domain.Account) (domain.Account, error) {
	return s.repo.Create(ctx, s.db, account)
}

func (s *AccountServiceImpl) Update(ctx context.Context, account domain.Account) (domain.Account, error) {
	return s.repo.Update(ctx, s.db, account)
}

func (s *AccountServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := s.repo.FindByID(ctx, s.db, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, s.db, id)
}
