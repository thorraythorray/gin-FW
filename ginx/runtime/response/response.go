package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type resData struct {
	Data interface{}
}

type resMsg struct {
	Msg string
}

func responseWithMsg(c *gin.Context, s int, msg string) {
	res := resMsg{Msg: msg}
	c.JSON(s, res)
}

func Success(c *gin.Context) {
	responseWithMsg(c, http.StatusOK, "ok")
}

func SuccessWithData(c *gin.Context, data interface{}) {
	res := resData{}
	res.Data = data
	c.JSON(http.StatusOK, res)
}

func Failed(c *gin.Context, errMsg string) {
	responseWithMsg(c, http.StatusBadRequest, errMsg)
}

func NotFound(c *gin.Context) {
	responseWithMsg(c, http.StatusNotFound, "目标已存在")
}

func AuthForbidden(c *gin.Context) {
	responseWithMsg(c, http.StatusNotFound, "权限禁止")
}

func UnAuthorized(c *gin.Context) {
	responseWithMsg(c, http.StatusNotFound, "登陆过期，请重新登陆")
}

func Conflict(c *gin.Context) {
	responseWithMsg(c, http.StatusNotFound, "目标冲突")
}
