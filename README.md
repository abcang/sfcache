sfcache
===

singleflight + on memory cache

## Example Usage

```go
type Hoge struct {
	ID   int64
	Name string
}

var cacheSizeBytes = 10 * 1024 * 1024
var expireSeconds = 30
var hogeSfc = sfcache.New[Hoge](cacheSizeBytes, expireSeconds)

func getHogeWithCache() (*Hoge, error) {
  key := "key"
  return hogeSfc.Do(key, func() (*Hoge, error) {
    return getHoge()
  })
}

func getHoge() (*Hoge, error) {
  return &Hoge{
		ID:   1,
		Name: "hoge",
	}, nil
}
```
