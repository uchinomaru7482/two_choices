package repository

import (
	"two_choices/pkg/rdb"
	"two_choices/src/domain"
)

type QuestionRepository interface {
	Save(scope *rdb.SessionScope, question *domain.Question) error
	Count(scope *rdb.SessionScope) (*uint64, error)
	GetByID(scope *rdb.SessionScope, id uint64) (*domain.Question, error)
}
