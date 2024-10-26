package sfcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Hoge struct {
	ID   int64
	Name string
}

func TestDo(t *testing.T) {
	itemSize := 100
	ttl := 3 * time.Second
	hogeSfc := New[*Hoge](itemSize, ttl)

	key := "key"
	exec := func() (*Hoge, error) {
		return hogeSfc.Do(key, func() (*Hoge, error) {
			return &Hoge{
				ID:   1,
				Name: "hoge",
			}, nil
		})
	}

	expected := Hoge{
		ID:   1,
		Name: "hoge",
	}

	hoge, err := exec()
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Exactly(t, expected, *hoge)

	hoge, err = exec()
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Exactly(t, expected, *hoge)

	_, ok := hogeSfc.cache.Get(key)
	assert.True(t, ok)

	time.Sleep(5 * time.Second)

	_, ok = hogeSfc.cache.Get(key)
	assert.False(t, ok)
}
