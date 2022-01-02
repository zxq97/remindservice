package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"remindservice/conf"
	"remindservice/rpc/remind/pb"
	"remindservice/server"
)

var (
	remindConf *conf.Conf
	err        error
)

func main() {
	remindConf, err = conf.LoadYaml(conf.RemindConfPath)
	if err != nil {
		panic(err)
	}

	err = server.InitService(remindConf)
	if err != nil {
		panic(err)
	}

	etcdRegistry := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = remindConf.Etcd.Addr
	})

	service := micro.NewService(
		micro.Name(remindConf.Grpc.Name),
		micro.Address(remindConf.Grpc.Addr),
		micro.Registry(etcdRegistry),
	)
	service.Init()
	err = remind_service.RegisterRemindServerHandler(
		service.Server(),
		new(server.RemindService),
	)
	if err != nil {
		panic(err)
	}
	err = service.Run()
	if err != nil {
		panic(err)
	}
}
