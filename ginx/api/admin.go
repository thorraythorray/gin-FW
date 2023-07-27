package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/dao"
	"github.com/thorraythorray/go-proj/ginx/form"
	"github.com/thorraythorray/go-proj/ginx/internal"
	"github.com/thorraythorray/go-proj/ginx/model"
	"github.com/thorraythorray/go-proj/ginx/util/request"
	"github.com/thorraythorray/go-proj/ginx/util/response"
	"github.com/thorraythorray/go-proj/pkg/auth"
)

type adminApi struct{}

func (u *adminApi) ObtainToken(c *gin.Context) {
	jwt := auth.JWT{
		SigningKey: internal.SignKey,
		CheckUser:  c.Request.Header.Get("X-User"),
		ExpireHour: internal.ExpireHour,
	}
	tokenstring, err := auth.AuthorizeImpl.Obtain(&jwt)
	if err != nil {
		response.ServerFailed(c, err.Error())
	}
	response.SuccessWithData(c, tokenstring)
}

func (u *adminApi) GetUsers(c *gin.Context) {
	var reqForm form.Pagination
	if err := c.ShouldBindJSON(&reqForm); err != nil {
		response.RequestFailed(c, err.Error())
	}

	offset, limit := reqForm.PageInfo()
	users, total, err := dao.UserDao.List(offset, limit)
	if err != nil {
		response.ServerFailed(c, err.Error())
	}
	res := reqForm.ResponseInfo(users, total)
	response.SuccessWithData(c, res)
}

func (u *adminApi) CreateUser(c *gin.Context) {
	var reqForm form.User
	if err := c.ShouldBindJSON(&reqForm); err != nil {
		response.RequestFailed(c, err.Error())
	}
	err := request.Ctx.ValidateForm(&reqForm)
	if err != nil {
		response.ServerFailed(c, err.Error())
	}
	newUser := model.User{
		Username: reqForm.Username,
		Password: reqForm.Password,
		Phone:    reqForm.Phone,
		Email:    reqForm.Email,
		Status:   uint8(internal.Active),
	}
	err = dao.UserDao.Create(&newUser)
	if err != nil {
		response.ServerFailed(c, err.Error())
	}
	response.SuccessWithData(c, newUser)
}

var AdminApiImpl = new(adminApi)
