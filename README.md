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

var itemSize = 100
var ttl = 30 * time.Second
var hogeSfc = sfcache.New[*Hoge](itemSize, ttl)

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
