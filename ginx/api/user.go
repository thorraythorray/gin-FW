package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/dao"
	"github.com/thorraythorray/go-proj/ginx/request"
	"github.com/thorraythorray/go-proj/ginx/response"
)

type userApi struct{}

func (u *userApi) GetUsers(c *gin.Context) {
	users, err := dao.UserDao.List()
	if err != nil {
		response.Failed(c, err.Error())
	}
	response.SuccessWithContent(c, users, "")
}

func (u *userApi) CreateUser(c *gin.Context) {
	var user request.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Failed(c, err.Error())
	}
	// if hasField(user, "Name")

}

var UserApiImpl = new(userApi)
