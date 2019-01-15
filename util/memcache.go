package util

import (
	"context"
	"time"

	"google.golang.org/appengine/capability"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/memcache"
)

func MemGet(ctx context.Context, key string) (string, bool) {
	if !capability.Enabled(ctx, "memcache", "*") {
		log.Warningf(ctx, "memcache is currently unavailable.")
		return "", false
	}

	item, err := memcache.Get(ctx, key)
	if err == nil {
		return string(item.Value), true
	}

	if err == memcache.ErrCacheMiss {
		log.Infof(ctx, err.Error())
	} else {
		log.Warningf(ctx, err.Error())
	}

	return "", false
}

func MemPut(ctx context.Context, key, val string, expire time.Duration) {
	if !capability.Enabled(ctx, "memcache", "*") {
		log.Warningf(ctx, "memcache is currently unavailable.")
		return
	}

	log.Infof(ctx, "memput: "+key)
	item := &memcache.Item{
		Key:        key,
		Value:      []byte(val),
		Expiration: expire,
	}
	err := memcache.Set(ctx, item)
	if err != nil {
		log.Warningf(ctx, err.Error())
	}
}
