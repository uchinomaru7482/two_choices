package usecase

import (
	"context"
	"two_choices/pkg/rdb"
	"two_choices/pkg/utils"
	"two_choices/src/domain"
	"two_choices/src/domain/repository"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

const getRandomRetryCount uint32 = 5

type QuestionUseCase interface {
	GetRandom(ctx context.Context, scope *rdb.SessionScope) (*domain.Question, error)
}

type questionUseCase struct {
	questionRepository repository.QuestionRepository
}

func NewQuestionUseCase(qr repository.QuestionRepository) QuestionUseCase {
	return &questionUseCase{
		questionRepository: qr,
	}
}

func (qu questionUseCase) GetRandom(ctx context.Context, scope *rdb.SessionScope) (*domain.Question, error) {
	// レコード総数取得
	count, err := qu.questionRepository.Count(scope)
	if err != nil {
		return nil, err
	}
	// ランダムなIDの質問を取得
	// 欠番を考慮したリトライ処理
	var domQuestion *domain.Question
	for i := uint32(1); i <= getRandomRetryCount; i++ {
		randID, err := utils.GenerateRandomInt(int64(*count))
		if err != nil {
			return nil, err
		}
		domQuestion, err = qu.questionRepository.GetByID(scope, uint64(randID))
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				continue
			}
			return nil, err
		}
		return domQuestion, nil
	}
	return nil, errors.New("get question faild")
}
