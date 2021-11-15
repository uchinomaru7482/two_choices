package interceptors

import (
	"context"
	"log"
	"two_choices/pkg/firebase"
	"two_choices/src/domain"

	"google.golang.org/grpc"
)

// AuthInterceptor - 認可
func AuthInterceptor(auths firebase.IFirebase) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// method := path.Base(info.FullMethod)

		// 認証が必要なサービスのみ認証を行う
		if authService, ok := info.Server.(domain.IAuthService); ok {
			ctxRet, err := auths.VerifyUser(ctx, authService.GetServiceAuthType(), false)
			if err != nil {
				log.Fatalf("ERROR: %+v", err)
			}
			ctx = ctxRet
		}

		// リクエストされた関数を実行
		res, err := handler(ctx, req)
		if err != nil {
			// %+v 構造体を出力する時に、フィールド名も出力される
			log.Fatalf("ERROR: %+v", err)
		}
		return res, nil
	}
}
