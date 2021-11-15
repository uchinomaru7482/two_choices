package usecase

import (
	"context"
	"two_choices/pkg/rdb"
	"two_choices/src/domain"
	"two_choices/src/domain/repository"
)

type UserUseCase interface {
	GetByUserID(ctx context.Context, scope *rdb.SessionScope, userID uint64) (*domain.User, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) GetByUserID(ctx context.Context, scope *rdb.SessionScope, userID uint64) (*domain.User, error) {
	user, err := uu.userRepository.GetByUserID(scope, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
