package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Foo string
	Bar []byte
}

func TestCache(t *testing.T) {
	c := NewInMemoryCache(time.Hour)
	var obj testStruct
	err := c.Get("key", &obj)
	assert.Equal(t, err, ErrCacheMiss)
	cacheObj := testStruct{
		Foo: "foo",
		Bar: []byte("bar"),
	}
	c.Set(&Item{
		Key:    "key",
		Object: &cacheObj,
	})
	cacheObj.Foo = "baz"
	err = c.Get("key", &obj)
	assert.Nil(t, err)
	assert.EqualValues(t, string(obj.Foo), "foo")
	assert.EqualValues(t, string(obj.Bar), "bar")

}
