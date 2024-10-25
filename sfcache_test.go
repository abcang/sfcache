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
	hogeSfc := New[Hoge](1000, 3)

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

	_, err = hogeSfc.cache.Get([]byte(key))
	assert.NoError(t, err)

	time.Sleep(5 * time.Second)

	_, err = hogeSfc.cache.Get([]byte(key))
	assert.EqualError(t, err, "Entry not found")
}
