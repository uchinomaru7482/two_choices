package firebase

import (
	"context"
	"log"
	"two_choices/src/domain"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/api/option"
)

// Config - Firebaseの設定
type Config struct {
	CredentialPath string
}

// FirebaseApp - firebase操作の主軸
type FirebaseApp struct {
	firebaseApp *firebase.App
}

// NewFirebaseApp - Firebaseの初期化
func NewFirebaseApp(config *Config) *FirebaseApp {
	opt := option.WithCredentialsFile(config.CredentialPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal("NewFirebaseApp: firebase.NewApp err")
	}
	return &FirebaseApp{firebaseApp: app}
}

// AddCustomClaim - カスタムクレーム追加
func (s *FirebaseApp) RegisterCustomClaim(ctx context.Context, uid string, customClaim map[string]interface{}) error {
	firebaseAuth, err := s.firebaseApp.Auth(ctx)
	if err != nil {
		return err
	}
	if err := firebaseAuth.SetCustomUserClaims(ctx, uid, customClaim); err != nil {
		return err
	}
	return nil
}

// GetUserByEmail - Emailでユーザ情報読込
func (s *FirebaseApp) GetUserByEmail(ctx context.Context, email string) (*auth.UserRecord, error) {
	firebaseAuth, err := s.firebaseApp.Auth(ctx)
	if err != nil {
		return nil, err
	}

	// firebaseからユーザ情報を読み出す
	user, err := firebaseAuth.GetUserByEmail(ctx, email)
	if err != nil {
		if auth.IsUserNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

// VerifyUser - 認証
func (s *FirebaseApp) VerifyUser(ctx context.Context, authType domain.ServiceAuthType, checkRevoked bool) (context.Context, error) {
	firebaseAuth, err := s.firebaseApp.Auth(ctx)
	if err != nil {
		return nil, err
	}

	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	var dToken *auth.Token
	if checkRevoked {
		// IDトークンの取り消しを検出する
		dToken, err = firebaseAuth.VerifyIDTokenAndCheckRevoked(ctx, token)
	} else {
		// IDトークンのペイロードの有効期限:exp、発行時:iat、対象:aud、発行元:iss、件名:sub、認証時間:auth_time
		// その後、IDトークンに署名されていることを確認し、UIDが使えるようになる
		dToken, err = firebaseAuth.VerifyIDToken(ctx, token)
	}
	if err != nil {
		return nil, err
	}

	// TODO カスタムクレームに登録済フラグが付いているか確認
	// 新規登録フロー中の確認はいらないため一旦省略
	/* if !verifyCustomClaim(ctx, authType, dToken) {
		return nil, err
	} */

	authedCtx := context.WithValue(ctx, domain.UIDKey, dToken.UID)
	authedCtx = context.WithValue(authedCtx, InfoKey, dToken.Firebase)
	return authedCtx, nil
}
