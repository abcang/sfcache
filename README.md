sfcache
===

singleflight + in memory cache

## Example Usage

```go
import (
	"github.com/abcang/sfcache"
)

type Hoge struct {
	ID   int64
	Name string
}

var cacheSizeBytes = 10 * 1024 * 1024
var expireSeconds = 30
var hogeSfc = sfcache.New[Hoge](cacheSizeBytes, expireSeconds)

func getHogeWithCache(id int64) (*Hoge, error) {
	key := strconv.Itoa(int(id))
	return hogeSfc.Do(key, func() (*Hoge, error) {
		return getHogeCore(id)
	})
}

func getHogeCore(id int64) (*Hoge, error) {
	return &Hoge{
		ID:   id,
		Name: "hoge",
	}, nil
}
```
