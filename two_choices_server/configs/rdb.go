package configs

import (
	"two_choices/pkg/rdb"
	"two_choices/pkg/utils"
)

// LoadRDBConfig - RDB設定読み込み
func LoadRDBConfig() (*rdb.Config, error) {
	var err error
	rdbConfig := &rdb.Config{}
	if rdbConfig.MasterHost, err = utils.GetStringFromEnv("MASTER_DB_HOST"); err != nil {
		return nil, err
	}
	if rdbConfig.SlaveHost, err = utils.GetStringFromEnv("SLAVE_DB_HOST"); err != nil {
		return nil, err
	}
	if rdbConfig.Name, err = utils.GetStringFromEnv("DB_PORT"); err != nil {
		return nil, err
	}
	if rdbConfig.Port, err = utils.GetIntFromEnv("DB_NAME"); err != nil {
		return nil, err
	}
	if rdbConfig.User, err = utils.GetStringFromEnv("DB_USER"); err != nil {
		return nil, err
	}
	if rdbConfig.Password, err = utils.GetStringFromEnv("DB_PASSWORD"); err != nil {
		return nil, err
	}
	if rdbConfig.MaxOpenConns, err = utils.GetIntFromEnv("DB_MAX_CONNS"); err != nil {
		return nil, err
	}
	if rdbConfig.MaxIdleConns, err = utils.GetIntFromEnv("DB_IDLE_CONNS"); err != nil {
		return nil, err
	}
	if rdbConfig.ConnMaxLifetime, err = utils.GetSecondFromEnv("DB_CONN_MAX_LIFE_TIME"); err != nil {
		return nil, err
	}
	return rdbConfig, nil
}
