package api

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api/request"
	"github.com/thorraythorray/go-proj/ginx/api/response"
	"github.com/thorraythorray/go-proj/ginx/dao"
	"github.com/thorraythorray/go-proj/ginx/form"
	"github.com/thorraythorray/go-proj/ginx/internal"
	"github.com/thorraythorray/go-proj/ginx/model"
	"github.com/thorraythorray/go-proj/global"
	"github.com/thorraythorray/go-proj/pkg/auth"
	"github.com/thorraythorray/go-proj/pkg/helper"
	"gorm.io/gorm"
)

type adminApi struct{}

func (u *adminApi) Register(c *gin.Context) {
	var register form.Register
	request.Validate(c, &register)

	newUser := model.User{
		Username: register.Username,
		Password: register.Password,
		Phone:    register.Phone,
		Email:    register.Email,
		Status:   uint8(internal.Active),
		Identity: helper.UuidString(),
	}
	err := dao.UserDao.Create(&newUser)
	if err != nil {
		response.ServerFailed(c, err.Error())
	} else {
		response.SuccessWithData(c, newUser)
	}
}

func (u *adminApi) Login(c *gin.Context) {
	var login form.Login
	request.Validate(c, &login)

	var user model.User
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
		userIdentify := fmt.Sprintf("%d", user.ID)
		token, err := auth.AuthorizerImpl.Obtain(&jwt, userIdentify)
		if err == nil {
			response.SuccessWithData(c, token)
		} else {
			response.ServerFailed(c, err.Error())
		}
	}
}

func (u *adminApi) GetUsers(c *gin.Context) {
	var pag form.Pagination
	request.Validate(c, &pag)

	offset, limit := pag.PageInfo()
	users, total, err := dao.UserDao.List(offset, limit)
	if err != nil {
		response.ServerFailed(c, err.Error())
	} else {
		res := pag.ResponseInfo(users, total)
		response.SuccessWithData(c, res)
	}
}

var AdminApiImpl = new(adminApi)
