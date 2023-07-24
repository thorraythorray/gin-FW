package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api/request"
	"github.com/thorraythorray/go-proj/ginx/api/response"
	"github.com/thorraythorray/go-proj/ginx/dao"
	"github.com/thorraythorray/go-proj/ginx/schema/form"
)

type userApi struct{}

func (u *userApi) GetUsers(c *gin.Context) {
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

func (u *userApi) CreateUser(c *gin.Context) {
	var reqForm form.User
	if err := c.ShouldBindJSON(&reqForm); err != nil {
		response.RequestFailed(c, err.Error())
	}

	err := request.Ctx.RequestValidate(&reqForm)
	if err != nil {
		response.ServerFailed(c, err.Error())
	}
	response.Success(c)
}

var UserApiImpl = new(userApi)
