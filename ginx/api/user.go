package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/common"
	"github.com/thorraythorray/go-proj/ginx/common/response"
	"github.com/thorraythorray/go-proj/ginx/dao"
	"github.com/thorraythorray/go-proj/ginx/schema/form"
)

type userApi struct{}

func (u *userApi) GetUsers(c *gin.Context) {
	users, err := dao.UserDao.List()
	if err != nil {
		response.Failed(c, err.Error())
	}
	response.SuccessWithData(c, users)
}

func (u *userApi) CreateUser(c *gin.Context) {
	var user form.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Failed(c, err.Error())
	}

	ok, err := common.Ctx.RequestValidate(&user)
	// fmt.Println(ok, err)
	if ok {
		response.Success(c)
	} else {
		response.Failed(c, err.Error())
	}
}

var UserApiImpl = new(userApi)
