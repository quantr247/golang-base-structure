package memcache

import (
	"time"

	goCache "github.com/patrickmn/go-cache"
)

func NewMemCacheHelper() *goCache.Cache {
	// Create a cache with a default expiration time of 5 minutes, and which
	// purges expired items every 10 minutes
	c := goCache.New(5*time.Minute, 10*time.Minute)
	return c
}
