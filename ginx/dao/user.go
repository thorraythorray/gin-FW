package dao

import (
	"errors"

	"github.com/thorraythorray/go-proj/ginx/model"
	"github.com/thorraythorray/go-proj/global"
	"gorm.io/gorm"
)

type userDao struct{}

func (dao *userDao) Exist(username, phone, email string) bool {
	var user model.User
	return !errors.Is(global.DB.Where("username = ?", username).Or("phone = ?", phone).Or("email = ?", email).First(&user).Error, gorm.ErrRecordNotFound)
}

func (dao *userDao) Create(u *model.User) error {
	isExist := dao.Exist(u.Username, u.Phone, u.Email)
	if isExist {
		return errors.New("记录已存在")
	}
	// 修改user成员内容格式
	result := global.DB.Create(u)
	return result.Error
}

func (dao *userDao) UpdateByID(id int, opts model.User) (model.User, error) {
	var upUser model.User
	global.DB.Find(&upUser, id)
	err := global.DB.Model(&upUser).Omit("ID").Updates(opts).Error
	return upUser, err
}

func (dao *userDao) DeleteByID(id int) error {
	return global.DB.Find(&model.User{}, id).Error
}

func (dao *userDao) List(offset, limit int) ([]model.User, uint64, error) {
	var users []model.User
	err := global.DB.Offset(offset).Limit(limit).Order("CreatedAt DESC").Find(&users).Error
	count := uint64(global.DB.RowsAffected)
	return users, count, err
}

var UserDao = new(userDao)
