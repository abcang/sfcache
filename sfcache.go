package sfcache

import (
	"encoding/json"

	"github.com/coocood/freecache"
	"golang.org/x/sync/singleflight"
)

type Sfcache[T any] struct {
	group         singleflight.Group
	cache         *freecache.Cache
	expireSeconds int
}

func New[T any](cacheSizeBytes int, expireSeconds int) *Sfcache[T] {
	return &Sfcache[T]{
		cache:         freecache.NewCache(cacheSizeBytes),
		expireSeconds: expireSeconds,
	}
}

func (sfc Sfcache[T]) Do(key string, fn func() (*T, error)) (*T, error) {
	v, err, _ := sfc.group.Do(key, func() (interface{}, error) {
		got, err := sfc.cache.Get([]byte(key))
		if err == nil {
			var data T
			err = json.Unmarshal(got, &data)
			if err != nil {
				return nil, err
			}

			return &data, nil
		}

		data, err := fn()
		if err != nil {
			return nil, err
		}

		jsonData, err := json.Marshal(*data)
		if err != nil {
			return nil, err
		}

		sfc.cache.Set([]byte(key), jsonData, sfc.expireSeconds)

		return data, nil
	})
	if err != nil {
		return nil, err
	}
	return v.(*T), nil
}

func (sfc Sfcache[T]) Delete(key string) {
	sfc.cache.Del([]byte(key))
}

func (sfc Sfcache[T]) Clear() {
	sfc.cache.Clear()
}
