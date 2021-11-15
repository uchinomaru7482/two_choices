package utils

import (
	"crypto/rand"

	"github.com/pkg/errors"
)

// GenerteRandStrLowAlphaAndNumber - ランダム文字列生成(小文字アルファベット、数値)
func GenerteRandStrLowAlphaAndNumber(digit uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	return randomString(digit, letters)
}

// randomString - ランダム文字列生成
func randomString(digit uint32, letters string) (string, error) {
	b := make([]byte, digit)
	// bにランダムなバイトが含まれる
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.New("randomString: rand.Read error")
	}

	var randStr string
	for _, v := range b {
		randStr += string(letters[int(v)%len(letters)])
	}
	return randStr, nil
}
