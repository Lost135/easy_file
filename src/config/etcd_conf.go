package config

import (
	"encoding/json"
	clientv3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/net/context"
	"time"
)

var Cli *clientv3.Client

// {username:"root",password:"password",bucket:{default:7},createdAt:"",delFlag:0}
const (
	key = "/sys/user/root"

	Username  = "root"
	Password  = "password"
	Bucket    = "default"
	CreatedAt = ""
	Deleted   = 0
)

type SysUserKv struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Bucket    string `json:"bucket"`
	CreatedAt string `json:"createdAt"`
	Deleted   int8   `json:"deleted"`
}

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
	kv := SysUserKv{Username: Username, Password: Password, Bucket: Bucket, CreatedAt: CreatedAt, Deleted: Deleted}
	if Yml.RootPassword != "" {
		kv.Password = Yml.RootPassword
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
	res, err := Cli.Get(ctx, key)
	if err != nil {
		return err
	}
	if res.Kvs == nil {
		_, err := Cli.Put(ctx, key, string(kv))
		if err != nil {
			return err
		}
	}
	cancel()
	return err
}
