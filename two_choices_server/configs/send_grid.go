package configs

import (
	"two_choices/pkg/sendgrid"
	"two_choices/pkg/utils"
)

// LoadSendGridConfig - RDB設定読み込み
func LoadSendGridConfig() (*sendgrid.Config, error) {
	var err error
	sendGridConfig := &sendgrid.Config{}
	if sendGridConfig.APIKey, err = utils.GetStringFromEnv("SEND_GRID_API_KEY"); err != nil {
		return nil, err
	}
	return sendGridConfig, nil
}
