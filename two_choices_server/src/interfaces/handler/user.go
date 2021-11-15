package interfaces

import (
	"context"
	"two_choices/configs"
	"two_choices/pkg/rdb"
	"two_choices/src/domain"
	"two_choices/src/usecase"

	pb "two_choices/generated"
)

type UserHandler interface {
	GetProfile(ctx context.Context, in *pb.UserCertProfile_GetProfileRequest) (*pb.UserCertProfile_GetProfileResponse, error)
}

type userHandler struct {
	userUseCase usecase.UserUseCase
	appConfig   *configs.AppConfig
	RDB         *rdb.Exector
}

func NewUserHandler(uu usecase.UserUseCase, appConfig *configs.AppConfig, rdbExector *rdb.Exector) UserHandler {
	return &userHandler{
		userUseCase: uu,
		appConfig:   appConfig,
		RDB:         rdbExector,
	}
}

func (uh userHandler) GetProfile(ctx context.Context, in *pb.UserCertProfile_GetProfileRequest) (*pb.UserCertProfile_GetProfileResponse, error) {
	var user *domain.User
	var err error
	if err := uh.RDB.Process(rdb.RDBKindSlave, func(scope *rdb.SessionScope) error {
		if user, err = uh.userUseCase.GetByUserID(ctx, scope, in.Id); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.UserCertProfile_GetProfileResponse{
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
