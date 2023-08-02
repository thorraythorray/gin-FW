package dao

import (
	"errors"

	"github.com/thorraythorray/go-proj/ginx/schema"
	"github.com/thorraythorray/go-proj/global"
	"gorm.io/gorm"
)

type userDao struct{}

func (dao *userDao) Exist(username, phone, email string) bool {
	var user schema.User
	return !errors.Is(
		global.DB.Where("username = ?", username).Or("phone = ?", phone).Or("email = ?", email).First(&user).Error,
		gorm.ErrRecordNotFound,
	)
}

func (dao *userDao) UpdateByID(id int, opts schema.UserModel) (schema.UserModel, error) {
	var upUser schema.UserModel
	global.DB.Find(&upUser, id)
	err := global.DB.Model(&upUser).Omit("ID").Updates(opts).Error
	return upUser, err
}

func (dao *userDao) DeleteByID(id int) error {
	return global.DB.Find(&schema.UserModel{}, id).Error
}

var UserDao = new(userDao)
