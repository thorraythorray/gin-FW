package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/pkg/gin/response"
)

type UserApi struct{}

func (u *UserApi) GetUsers(c *gin.Context) {
	response.Success(c)
}
