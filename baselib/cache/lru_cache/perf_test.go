package lru_cache

import (
	"fmt"
	"testing"
)

type MyValue []byte

func (mv MyValue) Size() int {
	return cap(mv)
}

func BenchmarkGet(b *testing.B) {
	cache := NewLRUCache(64 * 1024 * 1024)
	value := make(MyValue, 1000)
	cache.Set("stuff", value)
	for i := 0; i < b.N; i++ {
		val, ok := cache.Get("stuff")
		if !ok {
			err := fmt.Errorf("error")
			panic(err)
		}
		_ = val
	}
}
