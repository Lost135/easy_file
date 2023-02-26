package sys

import (
	"easy_file/src/config"
	"easy_file/src/structs"
	"easy_file/src/utils"
	"errors"
)

func CheckUserName(user structs.User) (count int64, res string, err error) {

	result := config.Db.Raw("select count(name) from user where name = '" + user.Name + "' and delFlag = 0").Scan(&count)
	config.CatchErr(result.Error)
	if count != 0 {
		err = errors.New("用户名已存在")
		return count, "用户名已存在", err
	}
	return count, "用户名不存在", nil
}

func Register(user structs.User) (res string, err error) {

	err = utils.EncodePwd(&user.Password)
	config.CatchFatal(err)
	result := config.Db.Create(&user)
	config.CatchWarn(result.Error)
	if result.RowsAffected != 1 {
		err = errors.New("创建错误")
		return "创建错误", err
	}
	return "创建成功", nil
}
