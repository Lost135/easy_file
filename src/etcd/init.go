package etcd

import (
	"easy_file/src/config"
	"fmt"
	"golang.org/x/net/context"
	"time"
)

const (
	userRoot = "Password"
)

func CheckInit() {

}

func InitEtcd() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := config.Cli.Put(ctx, userRoot, userRoot)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
}
