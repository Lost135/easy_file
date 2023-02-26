package database

import (
	"easy_file/src/config"
	"easy_file/src/structs"
	"errors"
	"gorm.io/gorm"
)

func UserCreate(user structs.User) (err error) {
	//支持根据构造体或map、[]map插入
	//.Create(map[string]interface{}{
	//  "Name": "xx", "Age": 18,
	//})

	result := config.Db.Create(&user)

	//忽略字段
	//db.Omit("Name", "CreatedAt").Create(&user)
	// INSERT INTO `users` (`age`) VALUES ("2020-01-01 00:00:00.000")

	/**使用 SQL 表达式
	db.Model(User{}).Create(map[string]interface{}{
		"Name": "xxx",
		"Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 100)"}},
	})
	INSERT INTO `users` (`name`,`location`) VALUES ("xx",ST_PointFromText("POINT(100 100)"));
	*/

	if result.Error == nil {
		return
	}
	return result.Error
}

/* 全局的钩子 */
func BatchCreate(users []structs.User) (err error) {
	//批量插入 跳过钩子
	result := config.Db.Session(&gorm.Session{SkipHooks: true}).Create(&users)

	return result.Error
}

func SelectDemo(user *structs.User) {
	/**
	只查一个
	*/
	// 获取第一条记录（主键升序）
	result := config.Db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// 根据主键获取记录，如果是非整型主键
	//db.First(&user, "id = ?", "string_primary_key")
	// SELECT * FROM users WHERE id = 'string_primary_key';

	// 获取一条记录，没有指定排序字段
	//db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	//db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func SelectDemo2(users *[]structs.User, name string) (err error) {
	// 拼接where  ?会被转义 %v 会SQL 注入
	//类似的有 Not Or Order Limit & Offset Group By & Having Distinct Joins
	//			Count   Scope
	result := config.Db.Table("user").Select("id", "name").Where("name LIKE ?", "%"+name+"%").Find(&users)

	//所以为什么不直接写sql呢  :(
	//db.Raw("SELECT name, age FROM user WHERE name like ?", "%"+name+"%").Scan(&users)

	// Map
	//db.Where(map[string]interface{}{"name": "xx", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "xx" AND age = 20;

	// 检查 ErrRecordNotFound 错误
	errors.Is(result.Error, gorm.ErrRecordNotFound)
	if result.Error != nil {
		return result.Error
	}
	return
}

func UpdateDemo(user *structs.User) (err error) {

	result := config.Db.Model(structs.User{Id: user.Id}).Where("age = ?", user.Password).Updates(map[string]interface{}{"name": user.Name})
	//db.Raw("update  user set name = 'xxx' WHERE id = ?", 37).Scan(&user)
	// UPDATE `user` SET `name`='user123' WHERE age = 13 AND `id` = 27
	if result.Error != nil {
		return result.Error
	}
	return
}

func DeleteDemo(user *structs.User) (err error) {

	result := config.Db.Where("name LIKE ?", "%"+user.Name+"%").Delete(&structs.User{Id: user.Id})
	// 不在构造体中定义软删除 DELETE from user where name LIKE "%xx%" and id = 12;
	// 定义软删除 UPDATE `user` SET `del_flag`=1 WHERE name LIKE '%test1%' AND `user`.`id` = 32 AND `user`.`del_flag` = 0
	if result.Error != nil {
		return result.Error
	}
	return
}
