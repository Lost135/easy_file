package database

import (
	"easy_file/src/config"
)

func RedisDemo(kv *config.KV) (res config.KV, err error) {
	kv.Value += "_isNew"
	err = config.Rdb.Set(config.Ctx, kv.Key, kv.Value, 0).Err()
	if err != nil {
		return res, err
	}

	val, err := config.Rdb.Get(config.Ctx, kv.Key).Result()
	if err != nil {
		return res, err
	}
	res.Value = val
	return res, err
}
