package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// GetStringFromEmv - 環境変数取得/文字列
func GetStringFromEnv(key string) (string, error) {
	env := os.Getenv(key)
	if len(env) <= 0 {
		return "", errors.New("GetStringFromEnv: env not found")
	}
	return env, nil
}

// GetIntFromEnv - 環境変数取得/整数
func GetIntFromEnv(key string) (int, error) {
	strEnv, err := GetStringFromEnv(key)
	if err != nil {
		return 0, err
	}
	intEnv, err := strconv.Atoi(strEnv)
	if err != nil {
		return 0, err
	}
	return intEnv, nil
}

// GetSecondFromEnv - 環境変数取得/秒数
func GetSecondFromEnv(key string) (time.Duration, error) {
	intEnv, err := GetIntFromEnv(key)
	if err != nil {
		return 0, err
	}
	secondEnv := time.Duration(intEnv) * time.Second
	return secondEnv, nil
}
