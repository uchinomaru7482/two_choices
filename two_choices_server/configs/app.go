package configs

import (
	"two_choices/pkg/cache"
	"two_choices/pkg/firebase"
	"two_choices/pkg/rdb"
	"two_choices/pkg/sendgrid"
)

// AppConfig - 設定
type AppConfig struct {
	RDBConfig      *rdb.Config
	CacheConfig    *cache.Config
	SendGridConfig *sendgrid.Config
	FirebaseConfig *firebase.Config
}

// LoadAppConfig - 設定読み込み
func (s *AppConfig) LoadAppConfig() error {
	var err error

	// RDB設定
	if s.RDBConfig, err = LoadRDBConfig(); err != nil {
		return err
	}
	// Cache設定
	if s.CacheConfig, err = LoadCacheConfig(); err != nil {
		return err
	}
	// SendGrid設定
	if s.SendGridConfig, err = LoadSendGridConfig(); err != nil {
		return err
	}
	// Firebase設定
	if s.FirebaseConfig, err = LoadFirebaseConfig(); err != nil {
		return err
	}

	return nil
}
