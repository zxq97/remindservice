package server

import "context"

func addUnread(ctx context.Context, uid int64, rType int32) error {
	return cacheAddUnread(ctx, uid, rType)
}

func deleteUnread(ctx context.Context, uid int64, rType int32) error {
	return cacheDeleteUnread(ctx, uid, rType)
}

func checkUnread(ctx context.Context, uid int64, rType int32) (bool, error) {
	return cacheCheckUnread(ctx, uid, rType)
}
