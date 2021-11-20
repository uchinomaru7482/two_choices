package persistence

import (
	"two_choices/pkg/rdb"
	"two_choices/src/domain"
	"two_choices/src/domain/repository"

	"github.com/pkg/errors"
)

type questionPersistence struct{}

func NewQuestionPersistence() repository.QuestionRepository {
	return &questionPersistence{}
}

// Save - 保存
func (qp questionPersistence) Save(scope *rdb.SessionScope, question *domain.Question) error {
	return errors.WithStack(scope.Conn.Model(question).Save(question).Error)
}

// Count - レコード数
func (qp questionPersistence) Count(scope *rdb.SessionScope) (*uint64, error) {
	var count uint64
	if err := scope.Conn.Table("questions").Count(&count).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &count, nil
}

// GetRandom - 取得/ID
func (qp questionPersistence) GetByID(scope *rdb.SessionScope, id uint64) (*domain.Question, error) {
	domQuestion := &domain.Question{}
	if err := scope.Conn.Where("id=?", id).First(domQuestion).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return domQuestion, nil
}
