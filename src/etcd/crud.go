package etcd

import (
	"context"
	"easy_file/src/config"
	"fmt"
	"time"
)

func Insert(key string, value string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := config.Cli.Put(ctx, key, value)
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	fmt.Printf("\nInsert====================================")
	fmt.Printf(key, value)
}

func Select(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	res, _ := config.Cli.Get(ctx, key)
	cancel()
	//遍历键值对
	fmt.Printf("\nSelect====================================")
	for _, kv := range res.Kvs {
		fmt.Printf("%s:%s \n", kv.Key, kv.Value)
	}
}

func Delete(key string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, _ = config.Cli.Delete(ctx, key)
	cancel()
	Select(key)
}

func Watch() {
	rch := config.Cli.Watch(context.Background(), "whoami") // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("\nWatch!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
