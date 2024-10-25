sfcache
===

singleflight + on memory cache

## Example Usage

```go
type Hoge struct {
	ID   int64
	Name string
}

cacheSizeBytes := 1024 * 1024
expireSeconds := 30
hogeSfc := sfcache.New[Hoge](cacheSizeBytes, expireSeconds)

key = "key"
hoge, err = hogeSfc.Do(key, func() (*Hoge, error) {
	return getHoge()
})

func getHoge() (*Hoge, error) {
  return &Hoge{
		ID:   1,
		Name: "hoge",
	}, nil
}
```
