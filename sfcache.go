package sfcache

import (
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
	"golang.org/x/sync/singleflight"
)

type Sfcache[T any] struct {
	group singleflight.Group
	cache *expirable.LRU[string, T]
}

func New[T any](itemSize int, ttl time.Duration) *Sfcache[T] {
	return &Sfcache[T]{
		cache: expirable.NewLRU[string, T](itemSize, nil, ttl),
	}
}

func (sfc Sfcache[T]) Do(key string, fn func() (T, error)) (T, error) {
	v, err, _ := sfc.group.Do(key, func() (interface{}, error) {
		cached, ok := sfc.cache.Get(key)
		if ok {
			return cached, nil
		}

		data, err := fn()
		if err != nil {
			return nil, err
		}

		sfc.cache.Add(key, data)

		return data, nil
	})
	if err != nil {
		var dummy T
		return dummy, err
	}
	return v.(T), nil
}

func (sfc Sfcache[T]) Delete(key string) {
	sfc.cache.Remove(key)
}

func (sfc Sfcache[T]) Clear() {
	sfc.cache.Purge()
}
