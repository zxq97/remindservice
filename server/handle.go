package server

import (
	"context"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"remindservice/conf"
	"remindservice/rpc/remind/pb"
)

type RemindService struct {
}

var (
	mcCli    *memcache.Client
	redisCli redis.Cmdable
	dbCli    *gorm.DB
)

func InitService(config *conf.Conf) error {
	var err error
	log.SetFlags(log.Ldate | log.Lshortfile | log.Ltime)
	mcCli = conf.GetMC(config.MC.Addr)
	redisCli = conf.GetRedisCluster(config.RedisCluster.Addr)
	//dbCli, err = conf.GetGorm(fmt.Sprintf(conf.MysqlAddr, config.Mysql.User, config.Mysql.Password, config.Mysql.Host, config.Mysql.Port, config.Mysql.DB))
	return err
}

func (rs *RemindService) AddUnread(ctx context.Context, req *remind_service.RemindInfo, res *remind_service.EmptyResponse) error {
	return addUnread(ctx, req.Uid, req.RemindType)
}

func (rs *RemindService) AddBatchUnread(ctx context.Context, req *remind_service.RemindBatchRequest, res *remind_service.EmptyResponse) error {
	return addBatchUnread(ctx, req.Uids, req.RemindType)
}

func (rs *RemindService) DeleteUnread(ctx context.Context, req *remind_service.RemindInfo, res *remind_service.EmptyResponse) error {
	return deleteUnread(ctx, req.Uid, req.RemindType)
}

func (rs *RemindService) CheckUnread(ctx context.Context, req *remind_service.RemindInfo, res *remind_service.CheckResponse) error {
	ok, err := checkUnread(ctx, req.Uid, req.RemindType)
	if err != nil {
		return err
	}
	res.Unread = ok
	return nil
}
