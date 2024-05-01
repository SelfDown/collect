package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgraph-io/ristretto"
)

type Test struct {
	Name string
}

func TestCache(t *testing.T) {
	fmt.Println(time.Now().Unix())
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
	}
	ta := Test{
		Name: "test",
	}
	cache.Set("test", ta, 100000)
	cache.Set("test1", "1", 100000)
	cache.Wait()
	value, t1 := cache.Get("test")
	value1, t1 := cache.Get("test1")
	fmt.Println(value)
	fmt.Println(value1)
	fmt.Println(t1)
}
