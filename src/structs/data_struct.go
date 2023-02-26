package structs

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

type User struct {
	Id        *int64     `bson:"id"`
	Name      string     `bson:"name"`
	Password  string     `bson:"password"`
	CreatedAt *time.Time `bson:"createdAt"`
	DelFlag   int        `bson:"delFlag"`
}

type Article struct {
	Id        *int64                 `gorm:"primaryKey"`
	UserId    *string                `form:"column:user_id" json:"userId"`
	Title     *string                `form:"column:title" json:"title"`
	Content   *string                `form:"column:content" json:"content"`
	CreatedAt *time.Time             `gorm:"column:createDate"`
	DelFlag   *soft_delete.DeletedAt `gorm:"softDelete:flag;column:delFlag"`
}
