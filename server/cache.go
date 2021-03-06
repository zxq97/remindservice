package server

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"remindservice/global"
	"remindservice/util/constant"
	"time"
)

const (
	RedisKeyUnread    = "remind_service_unread_%v_%v" // uid remind_type
	RedisKeyUnreadTTl = 30 * 24 * time.Hour
)

func cacheAddUnread(ctx context.Context, uid int64, rType int32) error {
	key := fmt.Sprintf(RedisKeyUnread, uid, rType)
	err := redisCli.Set(ctx, key, 1, RedisKeyUnreadTTl).Err()
	if err != nil {
		global.ExcLog.Printf("ctx %v cacheAddUnread uid %v rtype %v err %v", ctx, uid, rType, err)
	}
	return err
}

func cacheAddBatchUnread(ctx context.Context, uids []int64, rType int32) error {
	keys := make([]string, 0, constant.BatchSize<<1)
	var (
		key string
		err error
	)
	for k, v := range uids {
		key = fmt.Sprintf(RedisKeyUnread, v, rType)
		keys = append(keys, key)
		keys = append(keys, "1")
		if len(keys) == constant.BatchSize || k == len(uids)-1 {
			err = redisCli.MSet(ctx, keys).Err()
			if err != nil {
				global.ExcLog.Printf("ctx %v mset keys %v err %v", ctx, keys, err)
			}
			pipe := redisCli.Pipeline()
			for _, x := range keys {
				pipe.Expire(ctx, x, RedisKeyUnreadTTl)
			}
			_, err = pipe.Exec(ctx)
			if err != nil {
				global.ExcLog.Printf("ctx %v set expire keys %v err %v", ctx, keys, err)
			}
			keys = keys[:0]
		}
	}
	return nil
}

func cacheDeleteUnread(ctx context.Context, uid int64, rType int32) error {
	key := fmt.Sprintf(RedisKeyUnread, uid, rType)
	err := redisCli.Del(ctx, key).Err()
	if err != nil {
		global.ExcLog.Printf("ctx %v cacheDeleteUnread uid %v rtype %v err %v", ctx, uid, rType, err)
	}
	return err
}

func cacheCheckUnread(ctx context.Context, uid int64, rType int32) (bool, error) {
	key := fmt.Sprintf(RedisKeyUnread, uid, rType)
	ok, err := redisCli.Exists(ctx, key).Result()
	if err != nil && err != redis.Nil {
		global.ExcLog.Printf("ctx %v cacheCheckUnread uid %v rtype %v err %v", ctx, uid, rType, err)
		return false, err
	}
	return ok == 1, nil
}
