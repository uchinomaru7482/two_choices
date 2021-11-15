package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Config - 設定
type Config struct {
	Host        string
	Port        int
	MaxIdle     int           // Idle状態のコネクション数の上限
	MaxActive   int           // activeなコネクション数の上限
	IdleTimeout time.Duration // Idle状態になってからIdleTimeoutの時間が経過したらcloseする
}

// CachePool - キャッシュプール
type CachePool struct {
	config *Config
	Pool   *redis.Pool
}

// NewCachePool - 生成
func NewCachePool(config *Config) *CachePool {
	return &CachePool{config: config}
}

// InitPool - 初期化
func (s *CachePool) InitPool() {
	s.Pool = &redis.Pool{
		MaxIdle:     s.config.MaxIdle,
		MaxActive:   s.config.MaxActive,
		IdleTimeout: s.config.IdleTimeout,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", fmt.Sprintf("%s:%d", s.config.Host, s.config.Port))
		},
	}
}

// SetBytes - Bytesの保存
func (s *CachePool) SetBytes(key string, value interface{}, ttl int) error {
	// 構造体をjson形式に変換する
	serialized, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return s.multiExec("SET", key, serialized, ttl)
}

// GetBytes - Bytes取得
func (s *CachePool) GetBytes(key string) ([]byte, error) {
	conn := s.Pool.Get()
	defer conn.Close()

	b, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		// keyが無かった場合エラーが返ってくる？
		if err == redis.ErrNil {
			return nil, nil
		}
		return nil, err
	}
	return b, nil
}

func (s *CachePool) multiExec(command, key string, value interface{}, ttl int) error {
	conn := s.Pool.Get()
	defer conn.Close()

	if err := conn.Send("MULTI"); err != nil {
		return err
	}
	if err := conn.Send(command, key, value); err != nil {
		return err
	}
	if ttl > 0 {
		if err := conn.Send("EXPIRE", key, ttl); err != nil {
			return err
		}
	}
	if _, err := conn.Do("EXEC"); err != nil {
		return err
	}
	return nil
}

// Delete - 削除
func (s *CachePool) Delete(key string) error {
	conn := s.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}

// FlushAll - 全削除
func (s *CachePool) FlushAll() (string, error) {
	conn := s.Pool.Get()
	defer conn.Close()

	r, err := redis.String(conn.Do("FLUSHALL"))
	if err != nil {
		if err == redis.ErrNil {
			return "", nil
		}
		return "", err
	}
	return r, nil
}
