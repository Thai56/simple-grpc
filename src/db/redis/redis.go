package redisDb

import (
	"time"

	redis "github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// DefaultRedisOpts - returns default settings that will be used to create a new redis client
func DefaultRedisOpts() *redis.Options {
	return &redis.Options{
		Addr:         "acd-redis:6379",
		Password:     "", // no password set
		DB:           0,  // use default DB
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	}
}

// NewRedisClient - will return an instance of a new redis client server
func NewRedisClient() *redis.Client {
	log.Info("creating new redis client")

	redisClient := redis.NewClient(DefaultRedisOpts())

	log.Infof("Successfully connected with redis client")

	return redisClient
}
