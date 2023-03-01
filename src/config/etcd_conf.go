package config

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var Cli *clientv3.Client

func EtcdDb() {
	err := recover()
	Cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{Yml.Etcd.Addr},
		Username:    Yml.Etcd.Name,
		Password:    Yml.Etcd.Pass,
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		CatchFatal(err)
		return
	}
	CatchInfo("connect to etcd success")
	//defer cli.Close()

}
