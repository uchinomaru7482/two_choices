package configs

import (
	"two_choices/pkg/firebase"
	"two_choices/pkg/utils"
)

// LoadFirebaseConfig - Firebase設定読み込み
func LoadFirebaseConfig() (*firebase.Config, error) {
	var err error
	firebaseConfig := &firebase.Config{}
	if firebaseConfig.CredentialPath, err = utils.GetStringFromEnv("FIREBASE_CREDENTIAL_PATH"); err != nil {
		return nil, err
	}
	return firebaseConfig, nil
}
