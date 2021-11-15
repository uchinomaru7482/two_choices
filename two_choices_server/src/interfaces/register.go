package interfaces

import (
	"context"
	"two_choices/configs"
	"two_choices/pkg/firebase"
	"two_choices/pkg/rdb"
	"two_choices/src/infrastructure/persistence"
	interfaces "two_choices/src/interfaces/handler"
	"two_choices/src/usecase"

	pb "two_choices/generated"

	"google.golang.org/grpc"
)

// RegisterServices - サービスの登録
func RegisterServices(ctx context.Context, s *grpc.Server, appConfig *configs.AppConfig, rdbExector *rdb.Exector, firebaseApp firebase.IFirebase) error {
	userPersistence := persistence.NewUserPersistence()
	userUseCase := usecase.NewUserUseCase(userPersistence)
	userHandler := interfaces.NewUserHandler(userUseCase, appConfig, rdbExector)
	pb.RegisterUserCertProfileServiceServer(s, userHandler)

	userAuthenticationUseCase := usecase.NewUserAuthenticationUseCase(userPersistence)
	userAuthenticationHandler := interfaces.NewUserAuthenticationHandler(userAuthenticationUseCase, appConfig, rdbExector, firebaseApp)
	pb.RegisterUserAuthenticationServiceServer(s, userAuthenticationHandler)

	return nil
}
