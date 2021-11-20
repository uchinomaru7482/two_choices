package interfaces

import (
	"context"
	"two_choices/configs"
	"two_choices/pkg/rdb"
	"two_choices/src/domain"
	"two_choices/src/usecase"

	pb "two_choices/generated"
)

type UserQuestionHandler interface {
	GetRandom(ctx context.Context, in *pb.Empty) (*pb.UserQuestion_GetRandomResponse, error)
}

type userQuestionHandler struct {
	questionUseCase usecase.QuestionUseCase
	appConfig       *configs.AppConfig
	RDB             *rdb.Exector
}

func NewUserQuestionHandler(qu usecase.QuestionUseCase, appConfig *configs.AppConfig, rdbExector *rdb.Exector) UserQuestionHandler {
	return &userQuestionHandler{
		questionUseCase: qu,
		appConfig:       appConfig,
		RDB:             rdbExector,
	}
}

func (uqh userQuestionHandler) GetRandom(ctx context.Context, in *pb.Empty) (*pb.UserQuestion_GetRandomResponse, error) {
	var domQuestion *domain.Question
	var err error
	if err := uqh.RDB.Process(rdb.RDBKindSlave, func(scope *rdb.SessionScope) error {
		if domQuestion, err = uqh.questionUseCase.GetRandom(ctx, scope); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.UserQuestion_GetRandomResponse{
		Title:        domQuestion.Title,
		FirstAnswer:  domQuestion.FirstAnswer,
		SecondAnswer: domQuestion.SecondAnswer,
		FirstCount:   domQuestion.FirstCount,
		SecondCount:  domQuestion.SecondCount,
	}, nil
}
