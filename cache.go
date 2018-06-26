package cache

import "errors"

var (
	ErrCacheMiss = errors.New("cache: cache miss")
)

type Cache interface {
	Add(*Item) error
	Decrement(string, uint64) (uint64, error)
	Delete(string) error
	DeleteAll() error
	FlushAll() error
	Get(key string) (*Item, error)
	GetMulti(keys []string) (map[string]*Item, error)
	Increment(string, uint64) (uint64, error)
	Replace(*Item) error
	Set(*Item) error
	Touch(key string, seconds int32) error
}

type Item struct {
	Key        string
	Value      []byte
	Flags      uint32
	Expiration int32
}

func New(host string, port uint32) Cache {
	return newMemClient(host, port)
}

func ResumableErr(err error) bool {
	switch err {
	case ErrCacheMiss:
		return true
	}
	return false
}
