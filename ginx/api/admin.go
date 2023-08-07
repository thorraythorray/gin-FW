package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api/request"
	"github.com/thorraythorray/go-proj/ginx/api/response"
	"github.com/thorraythorray/go-proj/ginx/dao"
	"github.com/thorraythorray/go-proj/ginx/internal"
	"github.com/thorraythorray/go-proj/ginx/schema"
	"github.com/thorraythorray/go-proj/global"
	"github.com/thorraythorray/go-proj/pkg/auth"
	"github.com/thorraythorray/go-proj/pkg/helper"
	"gorm.io/gorm"
)

type adminApi struct{}

func (u *adminApi) Login(c *gin.Context) {
	var login schema.UserLogin
	if ok := request.Validate(c, &login); !ok {
		return
	}

	var user schema.UserModel
	isExist := !errors.Is(
		global.DB.Where("username = ? AND password = ?", login.Username, login.Password).First(&user).Error,
		gorm.ErrRecordNotFound,
	)
	if !isExist {
		response.NotFound(c)
	} else {
		jwt := auth.JWT{
			SigningKey: internal.JwtSignKey,
			ExpireHour: internal.JwtExpireHour,
		}
		token, err := auth.AuthorizerImpl.Obtain(&jwt, user.Username)
		res := map[string]string{"token": token}
		if err == nil {
			response.SuccessWithData(c, res)
		} else {
			response.ServerFailed(c, err)
		}
	}
}

func (u *adminApi) Register(c *gin.Context) {
	var register schema.UserRegister
	if ok := request.Validate(c, &register); !ok {
		return
	}
	notFound := dao.AdminDao.UserExist(register.Username, register.Phone, register.Email)
	if !notFound {
		response.Conflict(c)
		return
	}
	newUser := schema.UserModel{
		UserProfile: schema.UserProfile{
			User: schema.User{
				Username: register.Username,
				Password: register.Password,
			},
			Phone:  register.Phone,
			Email:  register.Email,
			RoleID: register.RoleID,
		},
		Status:   uint8(internal.Active),
		Identity: helper.UuidString(),
	}

	// 修改user成员内容格式
	err := dao.AdminDao.CreateUser(&newUser)
	if err != nil {
		response.ServerFailed(c, err)
	} else {
		response.SuccessWithData(c, newUser)
	}
}

func (u *adminApi) CreateRole(c *gin.Context) {
	var role schema.NewRole
	if ok := request.Validate(c, &role); !ok {
		return
	}
	notFound := dao.AdminDao.RoleExist(role.Role)
	if !notFound {
		response.Conflict(c)
		return
	}
	err := dao.AdminDao.CreateRole(&role)
	if err != nil {
		response.ServerFailed(c, err)
	} else {
		response.Success(c)
	}
}

func (u *adminApi) GetUsers(c *gin.Context) {
	var page schema.Pagination
	if ok := request.Validate(c, &page); !ok {
		return
	}
	offset, limit := page.PageInfo()

	var users []schema.UserModel
	res := global.DB.Offset(offset).Limit(limit).Order("create_at desc").Find(&users)
	if res.Error != nil {
		response.ServerFailed(c, res.Error)
	} else {
		res := page.ResponseInfo(users, res.RowsAffected)
		response.SuccessWithData(c, res)
	}
}

var AdminApiImpl = new(adminApi)