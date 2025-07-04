package parser

import (
	"github.com/hashicorp/golang-lru/v2/expirable"
)

type UACache struct {
	cache *expirable.LRU[string, *Result]
}

func NewUACache(size int) *UACache {
	c := expirable.NewLRU[string, *Result](size, nil, 0)
	return &UACache{cache: c}
}

func (c *UACache) Get(ua string) (*Result, bool) {
	return c.cache.Get(ua)
}

func (c *UACache) Add(ua string, res *Result) {
	c.cache.Add(ua, res)
}
