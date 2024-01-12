package my_cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var c *cache.Cache

func Cache() *cache.Cache {
	c = cache.New(15*time.Minute, 20*time.Minute)
	return c
}
