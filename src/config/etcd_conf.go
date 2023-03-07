package config

import (
	"easy_file/src/common"
	"easy_file/src/tool"
	"encoding/json"
	clientv3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/net/context"
	"time"
)

var Cli *clientv3.Client

// {username:"root",password:"password",bucket:{default:7},createdAt:"",delFlag:0}

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

	kv := common.User{
		Username: common.RootUsername,
		Password: common.RootPassword,
		Bucket: common.Bucket{
			Name: common.BucketDefault,
			Role: common.RootRole,
		},
		Status:    common.UserStatusDefault,
		CreatedAt: common.RootCreatedAt,
		Deleted:   common.DeletedDefault,
	}
	if Yml.RootPassword != "" {
		kv.Password = Yml.RootPassword
	}

	err = tool.EncodePwd(&kv.Password)
	if err != nil {
		CatchFatal(err)
		return
	}
	ksM, err2 := json.Marshal(kv)
	if err2 != nil {
		CatchFatal(err)
		return
	}
	err = InitSysUser(ksM)
	if err != nil {
		CatchFatal(err)
		return
	}
	CatchInfo("connect to etcd success")
	//defer cli.Close()
}

func InitSysUser(kv []byte) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	res, err := Cli.Get(ctx, common.RootKey)
	if err != nil {
		return err
	}
	if res.Kvs == nil {
		_, err := Cli.Put(ctx, common.RootKey, string(kv))
		if err != nil {
			return err
		}
	}
	cancel()
	return err
}
