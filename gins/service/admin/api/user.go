package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/gins/response"
	"github.com/thorraythorray/go-proj/gins/service/admin/dao"
)

type userApi struct{}

func (u *userApi) GetUsers(c *gin.Context) {
	users, err := dao.UserDao.List()
	if err != nil {
		response.Failed(c, err.Error())
	}
	response.SuccessWithContent(c, users, "")
}

var UserApiImpl = new(userApi)
