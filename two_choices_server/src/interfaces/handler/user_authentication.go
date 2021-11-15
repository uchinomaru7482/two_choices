package interfaces

import (
	"context"
	"two_choices/configs"
	"two_choices/pkg/firebase"
	"two_choices/pkg/rdb"
	"two_choices/src/domain"
	"two_choices/src/usecase"

	pb "two_choices/generated"
)

type UserAuthenticationHandler interface {
	SendVerificationMail(ctx context.Context, in *pb.UserAuthentication_SendVerificationMailRequest) (*pb.UserAuthentication_SendVerificationMailResponse, error)
	UserRegistration(ctx context.Context, in *pb.UserAuthentication_UserRegistrationRequest) (*pb.UserAuthentication_UserRegistrationResponse, error)
}

// userAuthenticationHandler - ユーザ登録ハンドラー
type userAuthenticationHandler struct {
	userAuthenticationUseCase usecase.UserAuthenticationUseCase
	appConfig                 *configs.AppConfig
	RDB                       *rdb.Exector
	firebaseApp               firebase.IFirebase
}

// NewUserAuthenticationHandler - 生成
func NewUserAuthenticationHandler(uau usecase.UserAuthenticationUseCase, appConfig *configs.AppConfig, rdbExector *rdb.Exector, firebaseApp firebase.IFirebase) UserAuthenticationHandler {
	return &userAuthenticationHandler{
		userAuthenticationUseCase: uau,
		appConfig:                 appConfig,
		RDB:                       rdbExector,
		firebaseApp:               firebaseApp,
	}
}

// GetServiceAuthType - 認証種別(新規登録ユーザ)
func (uh userAuthenticationHandler) GetServiceAuthType() domain.ServiceAuthType {
	return domain.ServiceAuthTypeNewUser
}

// SendVerificationMail - 認証メール送信
func (uh userAuthenticationHandler) SendVerificationMail(ctx context.Context, in *pb.UserAuthentication_SendVerificationMailRequest) (*pb.UserAuthentication_SendVerificationMailResponse, error) {
	uid := ctx.Value(domain.UIDKey).(string)
	var sendVerificationMailResult pb.UserAuthentication_SendVerificationMailResult
	// 確認メールを送信する
	if err := uh.RDB.Process(rdb.RDBKindSlave, func(scope *rdb.SessionScope) error {
		result, err := uh.userAuthenticationUseCase.SendVerificationMail(ctx, scope, uid, in.Name, in.Email, uh.appConfig.SendGridConfig, uh.firebaseApp)
		if err != nil {
			return err
		}
		switch *result {
		case usecase.SendVerificationMailResultSuccess:
			sendVerificationMailResult = pb.UserAuthentication_SEND_MAIL_SUCCESS
		case usecase.SendVerificationMailResultAlreadyExistsAddress:
			sendVerificationMailResult = pb.UserAuthentication_ALREADY_EXISTS_ADDRESS
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.UserAuthentication_SendVerificationMailResponse{
		Result: sendVerificationMailResult,
	}, nil
}

// UserRegistration - ユーザ登録
func (uh userAuthenticationHandler) UserRegistration(ctx context.Context, in *pb.UserAuthentication_UserRegistrationRequest) (*pb.UserAuthentication_UserRegistrationResponse, error) {
	uid := ctx.Value(domain.UIDKey).(string)
	var userRegistrationResult pb.UserAuthentication_UserRegistrationResult
	if err := uh.RDB.Process(rdb.RDBKindMaster, func(scope *rdb.SessionScope) error {
		result, err := uh.userAuthenticationUseCase.UserRegistration(ctx, scope, uid, in.State, uh.appConfig.SendGridConfig, uh.firebaseApp, uh.GetServiceAuthType())
		if err != nil {
			return err
		}
		switch *result {
		case usecase.UserRegistrationResultSuccess:
			userRegistrationResult = pb.UserAuthentication_USER_REGISTRATION_SUCCESS
		case usecase.UserRegistrationResultTimeout:
			userRegistrationResult = pb.UserAuthentication_TIMEOUT
		case usecase.UserRegistrationResultIncorrectState:
			userRegistrationResult = pb.UserAuthentication_INCORRECT_STATE
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &pb.UserAuthentication_UserRegistrationResponse{
		Result: userRegistrationResult,
	}, nil
}
