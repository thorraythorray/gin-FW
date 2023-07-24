package dao

import (
	"errors"

	"github.com/thorraythorray/go-proj/ginx/schema/model"
	"github.com/thorraythorray/go-proj/global"
	"gorm.io/gorm"
)

var db = global.DB

type userDao struct{}

func (dao *userDao) Exist(username, phone, email string) (bool, error) {
	var user model.User
	if !errors.Is(
		db.Where("name = ?", username).
			Or("phone = ?", phone).
			Or("email = ?", email).
			First(&user).Error, gorm.ErrRecordNotFound) {
		return true, errors.New("record exist")
	}
	return false, nil
}

func (dao *userDao) Create(u model.User) (model.User, error) {
	isExist, err := dao.Exist(u.Username, u.Phone, u.Email)
	if isExist {
		return u, err
	}
	// 修改user成员内容格式
	result := db.Create(&u)
	return u, result.Error
}

func (dao *userDao) UpdateByID(id int, opts model.User) (model.User, error) {
	var upUser model.User
	db.Find(&upUser, id)
	err := db.Model(&upUser).Omit("ID").Updates(opts).Error
	return upUser, err
}

func (dao *userDao) DeleteByID(id int) error {
	return db.Find(&model.User{}, id).Error
}

func (dao *userDao) List() ([]model.User, error) {
	var users []model.User
	err := db.Order("CreatedAt DESC").Find(&users).Error
	return users, err
}

var UserDao = new(userDao)
