package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/api/response"
	"github.com/thorraythorray/go-proj/ginx/dao"
	"github.com/thorraythorray/go-proj/ginx/form"
	"github.com/thorraythorray/go-proj/ginx/internal"
	"github.com/thorraythorray/go-proj/ginx/model"
	"github.com/thorraythorray/go-proj/pkg/validator"
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
	}
	err := dao.UserDao.Create(&newUser)
	if err != nil {
		response.ServerFailed(c, err.Error())
	} else {
		response.SuccessWithData(c, newUser)
	}
}

// func (u *adminApi) Login(c *gin.Context) {
// 	var loginForm form.Login
// 	if err := c.ShouldBindJSON(&loginForm); err != nil {
// 		response.RequestFailed(c, err.Error())
// 		return
// 	}
// 	errMsg := validator.Validate(loginForm)
// 	if errMsg != "" {
// 		response.RequestFailed(c, errMsg)
// 		return
// 	}
// 	isExist := !errors.Is(
// 		global.DB.Where("username = ? AND password = ?", loginForm.Username, loginForm.Password).First(&model.User),
// 		gorm.ErrRecordNotFound,
// 	)
// 	if !isExist {
// 		response.AuthForbidden(c, errors.New("未注册"))
// 	} else {
// 		jwt := auth.JWT{
// 			SigningKey: internal.SignKey,
// 			ExpireHour: internal.ExpireHour,
// 		}
// 		token := auth.AuthorizeImpl.Obtain()
// 	}
// }

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
