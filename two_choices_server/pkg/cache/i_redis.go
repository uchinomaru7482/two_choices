package cache

// IRedis - IFキャッシュ
type IRedis interface {
	SetBytes(key string, value interface{}, ttl int) error
	GetBytes(key string) ([]byte, error)
	Delete(key string) error
}

var Redis IRedis
