package dao

import (
	"errors"

	"github.com/thorraythorray/go-proj/ginx/schema"
	"github.com/thorraythorray/go-proj/global"
	"github.com/thorraythorray/go-proj/pkg/admin/rbac"
	"gorm.io/gorm"
)

type adminDao struct{}

// --------------------------USER------------------------------------
func (dao *adminDao) UserExist(username, phone, email string) bool {
	var user schema.User
	return errors.Is(
		global.DB.Where("username = ?", username).Or("phone = ?", phone).Or("email = ?", email).First(&user).Error,
		gorm.ErrRecordNotFound,
	)
}

func (dao *adminDao) CreateUser(u *schema.UserModel) error {
	var role *schema.RoleModel
	err := global.DB.First(&role, u.RoleID).Error
	if err == nil {
		err = global.DB.Create(u).Error
		if err == nil {
			// 更新casbin
			e := rbac.NewCasbin(global.DB)
			e.AddRoleForUser(u.Username, role.Name)
		}
	}
	return err
}

func (dao *adminDao) UpdateByID(id uint64, opts schema.UserModel) (schema.UserModel, error) {
	var upUser schema.UserModel
	global.DB.Find(&upUser, id)
	err := global.DB.Model(&upUser).Omit("ID").Updates(opts).Error
	return upUser, err
}

func (dao *adminDao) DeleteByID(id uint64) error {
	return global.DB.Find(&schema.UserModel{}, id).Error
}

// --------------------------ROLE------------------------------------
func (dao *adminDao) RoleExist(rolename string) bool {
	var role schema.RoleModel
	return errors.Is(
		global.DB.Where("name = ?", rolename).First(&role).Error,
		gorm.ErrRecordNotFound,
	)
}

func (dao *adminDao) CreateRole(r *schema.NewRole) error {
	e := rbac.NewCasbin(global.DB)
	for _, v := range r.CasbinInfos {
		e.AddPolicy(r.Role, v.Path, v.Method)
	}
	var new schema.RoleModel
	new.Role.Name = r.Role
	err := global.DB.Create(&new).Error
	return err
}

var AdminDao = new(adminDao)
