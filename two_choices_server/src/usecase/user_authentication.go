package usecase

import (
	"context"
	"encoding/json"
	"two_choices/pkg/cache"
	"two_choices/pkg/firebase"
	"two_choices/pkg/rdb"
	"two_choices/pkg/sendgrid"
	"two_choices/pkg/utils"
	"two_choices/src/domain"
	"two_choices/src/domain/repository"

	"firebase.google.com/go/auth"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type UserAuthenticationUseCase interface {
	SendVerificationMail(ctx context.Context, scope *rdb.SessionScope, uid, name, email string, sendGridConfig *sendgrid.Config, firebaseApp firebase.IFirebase) (*SendVerificationMailResult, error)
	UserRegistration(ctx context.Context, scope *rdb.SessionScope, uid, state string, sendGridConfig *sendgrid.Config, firebaseApp firebase.IFirebase, serviceAuthType domain.ServiceAuthType) (*UserRegistrationResult, error)
}

type userAuthenticationUseCase struct {
	userRepository repository.UserRepository
}

func NewUserAuthenticationUseCase(ur repository.UserRepository) UserAuthenticationUseCase {
	return &userAuthenticationUseCase{
		userRepository: ur,
	}
}

// stateDigit - stateの桁数
const stateDigit uint32 = 32

// verifyEmailTTL - メール認証TTL:30分
const verifyEmailTTL uint32 = 60 * 30

// SendVerificationMailResult - ユーザ登録結果
type SendVerificationMailResult int

const (
	// SendVerificationMailResultSuccess - 成功
	SendVerificationMailResultSuccess SendVerificationMailResult = iota
	// SendVerificationMailResultAlreadyExistsAddress - 登録済のアドレス
	SendVerificationMailResultAlreadyExistsAddress
)

// UserRegistrationResult - ユーザ登録結果
type UserRegistrationResult int

const (
	// UserRegistrationResultSuccess - 成功
	UserRegistrationResultSuccess UserRegistrationResult = iota
	// UserRegistrationResultTimeout - 時間切れ
	UserRegistrationResultTimeout
	// UserRegistrationResultIncorrectState - 不正なstate
	UserRegistrationResultIncorrectState
)

// verifyEmailCacheKey - メール認証用キャッシュキー
func verifyEmailCacheKey(firebaseUID string) string {
	return "verify_email_uid_" + firebaseUID
}

func (uau userAuthenticationUseCase) SendVerificationMail(ctx context.Context, scope *rdb.SessionScope, uid, name, email string, sendGridConfig *sendgrid.Config, firebaseApp firebase.IFirebase) (*SendVerificationMailResult, error) {
	sendVerificationMail := SendVerificationMailResultSuccess
	// メールアドレスバリデーション
	if len(email) >= 254 {
		return nil, errors.New("SendVerificationMail: email valid err")
	}
	// TODO バリデーションの中身はdomainに実装する

	// ユーザの登録済チェック
	_, err := uau.userRepository.GetByEmail(scope, email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.WithStack(err)
	} else if err == nil {
		return nil, errors.New("SendVerificationMail: already exists email err")
	}

	// Firebaseへのユーザ存在チェック
	firebaseUser, err := firebaseApp.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Firebaseにユーザが登録されていない
	if firebaseUser == nil {
		return nil, errors.New("SendVerificationMail: firebase user not found")
	}
	// Firebase認証済
	registerComplete, ok := firebaseUser.CustomClaims[domain.CustomClaimRegistUserCompleteKey]
	if ok && registerComplete == true {
		return nil, errors.New("SendVerificationMail: custom claim register complete true")
	}

	// state生成
	state, err := utils.GenerteRandStrLowAlphaAndNumber(stateDigit)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	hash := domain.MailState{
		State: state,
		Name:  name,
		Email: email,
	}

	// stateキャッシュ保存(有効期限30分)
	if err := cache.Redis.SetBytes(verifyEmailCacheKey(uid), hash, int(verifyEmailTTL)); err != nil {
		return nil, errors.WithStack(err)
	}

	// メールを送信する
	mail := &domain.Mail{
		FromName:         domain.FromName,
		FromEmail:        domain.FromEmail,
		ToName:           name,
		ToEmail:          email,
		Subject:          "会員登録",
		PlainTextContent: "localhost:3000/s/auth/registration_complete/?state=" + state,
		HTMLContent:      "",
	}
	if err := mail.Send(ctx, sendGridConfig); err != nil {
		return nil, errors.WithStack(err)
	}
	return &sendVerificationMail, nil
}

func (uau userAuthenticationUseCase) UserRegistration(ctx context.Context, scope *rdb.SessionScope, uid, state string, sendGridConfig *sendgrid.Config, firebaseApp firebase.IFirebase, serviceAuthType domain.ServiceAuthType) (*UserRegistrationResult, error) {
	userRegistrationResult := UserRegistrationResultSuccess
	// Email取得
	firebaseInfoRaw := ctx.Value(firebase.InfoKey)
	if firebaseInfoRaw == nil {
		return nil, errors.New("UserRegistration: firebase info raw not found")
	}
	firebaseInfo := firebaseInfoRaw.(auth.FirebaseInfo)
	val, ok := firebaseInfo.Identities["email"]
	if !ok {
		return nil, errors.New("UserRegistration: firebase info not found")
	}
	JWTEmail := val.([]interface{})[0].(string)

	// JWTを厳密に判定
	ctx, err := firebaseApp.VerifyUser(ctx, serviceAuthType, true)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// state一致チェック
	jHash, err := cache.Redis.GetBytes(verifyEmailCacheKey(uid))
	if err != nil || jHash == nil {
		userRegistrationResult = UserRegistrationResultIncorrectState
		return &userRegistrationResult, nil
	}
	hash := &domain.MailState{
		State: "",
		Name:  "",
		Email: "",
	}
	// json形式を構造体に変換
	err = json.Unmarshal(jHash, hash)
	if err = hash.VerifyEmailState(ctx, JWTEmail, state); err != nil {
		return nil, errors.WithStack(err)
	}

	// DBに情報を登録
	user := &domain.User{
		Name:  hash.Name,
		Email: hash.Email,
		Uid:   uid,
	}
	if err := uau.userRepository.Save(scope, user); err != nil {
		return nil, errors.WithStack(err)
	}

	// 登録完了メールを送信
	mail := &domain.Mail{
		FromName:         domain.FromName,
		FromEmail:        domain.FromEmail,
		ToName:           hash.Name,
		ToEmail:          hash.Email,
		Subject:          "会員登録",
		PlainTextContent: "ご登録いただきありがとうございます。お楽しみ下さい。",
		HTMLContent:      "",
	}
	if err := mail.Send(ctx, sendGridConfig); err != nil {
		return nil, errors.WithStack(err)
	}

	// JWTにカスタムクレームを追加
	customClaim := map[string]interface{}{
		domain.CustomClaimRegistUserCompleteKey: true,
		domain.CustomClaimUserIDKey:             user.ID,
	}
	if err = firebaseApp.RegisterCustomClaim(ctx, uid, customClaim); err != nil {
		return nil, errors.WithStack(err)
	}

	// Redisのデータ削除
	if err := cache.Redis.Delete(verifyEmailCacheKey(uid)); err != nil {
		return nil, errors.WithStack(err)
	}
	return &userRegistrationResult, nil
}
