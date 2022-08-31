package kv

import (
	"context"
	"fmt"
	"time"

	"github.com/domenetwork/dome-lib/pkg/cfg"
	"github.com/domenetwork/dome-lib/pkg/log"
	"github.com/go-redis/redis/v8"
)

var _ KV = &Redis{}

// Redis interfaces with the famous KV storage engine Redis.
type Redis struct {
	cli *redis.Client
}

// Close the connection.
func (kv *Redis) Close() (err error) {
	log.D("redis", "close")
	return
}

// Delete the item from the store.
func (kv *Redis) Delete(key string) (err error) {
	log.D("redis", "delete", key)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = kv.cli.Del(ctx, key).Err()
	return
}

// Get the value that matches the provided key from Redis.
func (kv *Redis) Get(key string) (value interface{}, err error) {
	log.D("redis", "get", key)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	value, err = kv.cli.Get(ctx, key).Result()
	if err == redis.Nil || value == "" {
		log.E("redis.get", "nil or empty")
		return
	} else if err != nil {
		log.E("redis.get", "unknown error", err)
		return
	}
	return
}

// Open the connection to Redis.
func (kv *Redis) Open() (err error) {
	log.D("redis", "open")
	opts := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Str("redis.host"), cfg.Int("redis.port")),
		DB:       0,
		Password: cfg.Str("redis.pass"),
	}
	kv.cli = redis.NewClient(opts)
	return
}

// Set the provided value to the provided key in Redis.
func (kv *Redis) Set(key string, value interface{}, ttl time.Duration) (err error) {
	log.D("redis", "set", key, value)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err = kv.cli.Set(ctx, key, value, ttl).Err(); err != nil {
		log.D("redis.set", key, value, ttl)
	}
	return
}
