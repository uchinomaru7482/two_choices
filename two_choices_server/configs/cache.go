package configs

import (
	"two_choices/pkg/cache"
	"two_choices/pkg/utils"
)

// LoadCacheConfig - Cache設定読み込み
func LoadCacheConfig() (*cache.Config, error) {
	var err error
	cacheConfig := &cache.Config{}
	if cacheConfig.Host, err = utils.GetStringFromEnv("CACHE_HOST"); err != nil {
		return nil, err
	}
	if cacheConfig.Port, err = utils.GetIntFromEnv("CACHE_PORT"); err != nil {
		return nil, err
	}
	if cacheConfig.MaxIdle, err = utils.GetIntFromEnv("CACHE_MAX_IDLE"); err != nil {
		return nil, err
	}
	if cacheConfig.MaxActive, err = utils.GetIntFromEnv("CACHE_MAX_ACTIVE"); err != nil {
		return nil, err
	}
	if cacheConfig.IdleTimeout, err = utils.GetSecondFromEnv("CACHE_IDLE_TIMEOUT"); err != nil {
		return nil, err
	}
	return cacheConfig, nil
}
