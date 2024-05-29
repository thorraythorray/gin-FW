package dao

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/thorraythorray/go-Jarvis/admin/rbac"
	"github.com/thorraythorray/go-proj/ginx/schema"
	"github.com/thorraythorray/go-proj/global"
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
	err := global.DB.Create(u).Error
	if err == nil {
		// 更新casbin
		e := rbac.NewCasbin(global.DB)
		e.AddRoleForUser(strconv.Itoa(int(u.ID)), strconv.Itoa(int(u.RoleID)))
	}
	return err
}

func (dao *adminDao) UpdateByID(id uint64, opts schema.UserProfile) (schema.UserModel, error) {
	var upUser schema.UserModel
	global.DB.Find(&upUser, id)
	everRole := upUser.RoleID
	err := global.DB.Model(&upUser).Updates(opts).Error
	if err == nil {
		// 更新casbin
		if upUser.RoleID != everRole {
			e := rbac.NewCasbin(global.DB)
			e.DeleteRoleForUser(strconv.Itoa(int(upUser.ID)), strconv.Itoa(int(everRole)))
			e.AddRoleForUser(strconv.Itoa(int(upUser.ID)), strconv.Itoa(int(upUser.RoleID)))
		}
	}
	return upUser, err
}

func (dao *adminDao) DeleteByID(id string) error {
	_id, _ := strconv.Atoi(id)
	var u schema.UserModel
	fmt.Println("xsssss", _id)
	global.DB.First(&u, _id)
	fmt.Println("xxxxxxx", _id)
	roleid := strconv.Itoa(int(u.RoleID))
	// 删除用户
	err := global.DB.Delete(&u).Error
	if err == nil {
		// 删除casbin权限
		e := rbac.NewCasbin(global.DB)
		e.DeleteRoleForUser(id, roleid)
	}
	return err
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
	var new schema.RoleModel
	new.Role.Name = r.Role
	err := global.DB.Create(&new).Error
	if err == nil {
		// 添加casbin权限
		e := rbac.NewCasbin(global.DB)
		for _, v := range r.CasbinInfos {
			e.AddPolicy(strconv.Itoa(int(new.ID)), v.Path, v.Method)
		}
	}
	return err
}

var AdminDao = new(adminDao)
