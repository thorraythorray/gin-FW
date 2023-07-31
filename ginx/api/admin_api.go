package api

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api/response"
	"github.com/thorraythorray/go-proj/ginx/dao"
	"github.com/thorraythorray/go-proj/ginx/form"
	"github.com/thorraythorray/go-proj/ginx/internal"
	"github.com/thorraythorray/go-proj/ginx/model"
	"github.com/thorraythorray/go-proj/global"
	"github.com/thorraythorray/go-proj/pkg/auth"
	"github.com/thorraythorray/go-proj/pkg/helper"
	"github.com/thorraythorray/go-proj/pkg/validator"
	"gorm.io/gorm"
)

type adminApi struct{}

func (u *adminApi) Register(c *gin.Context) {
	var reqForm form.Register
	if err := c.ShouldBindJSON(&reqForm); err != nil {
		response.RequestFailed(c, err.Error())
		return
	}
	errMsg := validator.Validate(reqForm)
	if errMsg != "" {
		response.RequestFailed(c, errMsg)
		return
	}
	newUser := model.User{
		Username: reqForm.Username,
		Password: reqForm.Password,
		Phone:    reqForm.Phone,
		Email:    reqForm.Email,
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
	var loginForm form.Login
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		response.RequestFailed(c, err.Error())
		return
	}
	errMsg := validator.Validate(loginForm)
	if errMsg != "" {
		response.RequestFailed(c, errMsg)
		return
	}
	var user model.User
	isExist := !errors.Is(
		global.DB.Where("username = ? AND password = ?", loginForm.Username, loginForm.Password).First(&user).Error,
		gorm.ErrRecordNotFound,
	)
	if !isExist {
		response.UnAuthorized(c)
	} else {
		jwt := auth.JWT{
			SigningKey: internal.JwtSignKey,
			ExpireHour: internal.JwtExpireHour,
		}
		userIdentify := fmt.Sprintf("%d", user.ID)
		token, err := auth.AuthorizeImpl.Obtain(&jwt, userIdentify)
		if err != nil {
			response.ServerFailed(c, err.Error())
			return
		}
		response.SuccessWithData(c, token)
	}
}

func (u *adminApi) GetUsers(c *gin.Context) {
	var reqForm form.Pagination
	if err := c.ShouldBindJSON(&reqForm); err != nil {
		response.RequestFailed(c, err.Error())
		return
	}

	offset, limit := reqForm.PageInfo()
	users, total, err := dao.UserDao.List(offset, limit)
	if err != nil {
		response.ServerFailed(c, err.Error())
	} else {
		res := reqForm.ResponseInfo(users, total)
		response.SuccessWithData(c, res)
	}
}

var AdminApiImpl = new(adminApi)
