package db

import (
	"context"
	"easy_file/server/common"
	"easy_file/server/config"
	"easy_file/server/tool"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type KV struct {
	Key    string
	Value  string
	Entity map[string][]byte
}

func (kv *KV) Insert() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := config.Cli.Put(ctx, kv.Key, kv.Value, clientv3.WithLease(GetLease(common.LeaseTimeDefault)))
	cancel()
	return err
}

func (kv *KV) Select() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	res, err := config.Cli.Get(ctx, kv.Value)
	cancel()
	for _, r := range res.Kvs {
		kv.Entity[string(r.Key)] = r.Value
	}
	return err
}

func (kv *KV) Delete() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, _ = config.Cli.Delete(ctx, kv.Key)
	cancel()
}

func (kv *KV) Watch() {
	rch := config.Cli.Watch(context.Background(), "whoami") // <-chan WatchResponse
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("\nWatch!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			fmt.Printf("Type: %s Key:%s Value:%s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func GetLease(ttl int64) clientv3.LeaseID {
	lease, err := clientv3.NewLease(config.Cli).Grant(context.TODO(), ttl)
	//返回租约的id
	//TODO 租约时间问题
	config.CatchWarn(err)
	return lease.ID
}

func LoginUser(user *common.User) (string, error) {
	kv := KV{Key: "username", Value: common.UserPrefix + user.Username, Entity: make(map[string][]byte)}
	_ = kv.Select()
	userCheck := common.User{}
	_ = json.Unmarshal(kv.Entity[kv.Value], &userCheck)
	if !tool.ComparePwd(userCheck.Password, user.Password) {
		return "", errors.New("password error")
	}
	uuidNew := uuid.New()
	kv.Key = common.TokenPrefix + user.Username
	kv.Value = uuidNew.String()
	err := kv.Insert()
	if err != nil {
		return "", err
	}
	user = &userCheck
	return kv.Value, nil
}
