package server

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

const (
	RedisKeyUnread    = "remind_service_unread_%v_%v" // uid remind_type
	RedisKeyUnreadTTl = 30 * 24 * time.Hour
)

func cacheAddUnread(ctx context.Context, uid int64, rType int32) error {
	key := fmt.Sprintf(RedisKeyUnread, uid, rType)
	err := redisCli.Set(key, 1, RedisKeyUnreadTTl).Err()
	if err != nil {
		log.Printf("ctx %v cacheAddUnread uid %v rtype %v err %v", ctx, uid, rType, err)
	}
	return err
}

func cacheDeleteUnread(ctx context.Context, uid int64, rType int32) error {
	key := fmt.Sprintf(RedisKeyUnread, uid, rType)
	err := redisCli.Del(key).Err()
	if err != nil {
		log.Printf("ctx %v cacheDeleteUnread uid %v rtype %v err %v", ctx, uid, rType, err)
	}
	return err
}

func cacheCheckUnread(ctx context.Context, uid int64, rType int32) (bool, error) {
	key := fmt.Sprintf(RedisKeyUnread, uid, rType)
	ok, err := redisCli.Exists(key).Result()
	if err != nil && err != redis.Nil {
		log.Printf("ctx %v cacheCheckUnread uid %v rtype %v err %v", ctx, uid, rType, err)
		return false, err
	}
	return ok == 1, nil
}
