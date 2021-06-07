package redisclient

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/mk1010/idustry/config"
)

var Rdb *redis.Client

func InitRedisClient() error {
	if len(config.ConfInstance.RedisHosts) != 0 {
		rdb := redis.NewClient(&redis.Options{
			Addr:     config.ConfInstance.RedisHosts[0],
			Password: config.ConfInstance.RedisClusterName,
			DB:       0, // use default DB
		})
		err := rdb.Ping(context.Background()).Err()
		if err != nil {
			return err
		}
		Rdb = rdb
	}
	return nil
}
