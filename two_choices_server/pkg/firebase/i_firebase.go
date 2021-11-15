package firebase

import (
	"context"
	"two_choices/src/domain"

	"firebase.google.com/go/auth"
)

// InfoKey - firebase情報キー
const InfoKey string = "firebase_info"

// IFirebase - IFファイアーベース
type IFirebase interface {
	GetUserByEmail(ctx context.Context, email string) (*auth.UserRecord, error)
	RegisterCustomClaim(ctx context.Context, uid string, customClaim map[string]interface{}) error
	VerifyUser(ctx context.Context, authType domain.ServiceAuthType, checkRevoked bool) (context.Context, error)
}
