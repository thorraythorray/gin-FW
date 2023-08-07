package schema

import "github.com/thorraythorray/go-proj/pkg/admin/rbac"

// 用户数据结构
type User struct {
	Username string `json:"username" binding:"required" validate:"min=6,max=20"`
	Password string `json:"password" binding:"required" validate:"min=8,max=20"`
}
type UserProfile struct {
	User
	Phone  string `json:"phone" binding:"required" validate:"phone"`
	Email  string `json:"email" binding:"required" validate:"email"`
	RoleID uint64 `json:"role_id" binding:"required"`
}

type UserModel struct {
	BaseModel
	UserProfile
	Status   uint8  `binding:"-"`
	Identity string `binding:"-"`
}

func (UserModel) TableName() string {
	return "users"
}

type UserRegister struct {
	UserProfile
}

func (form *UserRegister) Validate() error {
	return nil
}

type UserLogin struct {
	User
}

func (form *UserLogin) Validate() error {
	return nil
}

// 角色数据结构
type Role struct {
	Name string `json:"name" binding:"required"`
}

type RoleModel struct {
	BaseModel
	Role
}

func (RoleModel) TableName() string {
	return "roles"
}

type NewRole struct {
	rbac.CasbinRules
}

/**
 * @description: middleware用户认证数据结构
 * @return {*}
 */
type AuthorizedUserInfo struct {
	UserID   uint64 `gorm:"column:user_id"`
	UserName string `gorm:"column:user_name"`
	RoleID   uint64 `gorm:"column:role_id"`
	RoleName string `gorm:"column:role_name"`
}
