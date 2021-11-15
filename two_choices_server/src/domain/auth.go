package domain

import (
	"context"

	"github.com/pkg/errors"
)

// ServiceAuthType - 認証種別
type ServiceAuthType uint32

const (
	// ServiceAuthTypeUser - ユーザ
	ServiceAuthTypeUser ServiceAuthType = iota
	// ServiceAuthTypeNewUser - 新規登録中のユーザ
	ServiceAuthTypeNewUser
)

// UIDKey - ユーザ判定キー
const UIDKey string = "uid"

// CustomClaimRegistUserCompleteKey - カスタムクレーム：ユーザ登録完了フラグ
const CustomClaimRegistUserCompleteKey string = "regist_complete"

// CustomClaimUserIDKey - カスタムクレーム：ユーザID
const CustomClaimUserIDKey string = "user_id"

// MailState - メール送信用state
type MailState struct {
	State string
	Name  string
	Email string
}

// IAuthService - 要認証サービス
type IAuthService interface {
	GetServiceAuthType() ServiceAuthType
}

// VerifyEmailState - State認証
func (mailState MailState) VerifyEmailState(ctx context.Context, jwtEmail, state string) error {
	if mailState.Email != jwtEmail || mailState.State != state {
		return errors.New("varify email err")
	}
	return nil
}
