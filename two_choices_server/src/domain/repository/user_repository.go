package repository

import (
	"two_choices/pkg/rdb"
	"two_choices/src/domain"
)

type UserRepository interface {
	Save(scope *rdb.SessionScope, user *domain.User) error
	GetByUserID(scope *rdb.SessionScope, userID uint64) (*domain.User, error)
	GetByEmail(scope *rdb.SessionScope, email string) (*domain.User, error)
}
