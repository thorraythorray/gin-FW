package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type msgResponse struct {
	Msg string
}

func ResponseWithMsg(c *gin.Context, s int, msg string) {
	res := msgResponse{Msg: msg}
	c.JSON(s, res)
}

func RequestFailed(c *gin.Context, errMsg string) {
	ResponseWithMsg(c, http.StatusBadRequest, errMsg)
}

func ServerFailed(c *gin.Context, errMsg string) {
	ResponseWithMsg(c, http.StatusInternalServerError, errMsg)
}

func NotFound(c *gin.Context) {
	ResponseWithMsg(c, http.StatusNotFound, "目标已存在")
}

func AuthForbidden(c *gin.Context) {
	ResponseWithMsg(c, http.StatusNotFound, "权限禁止")
}

func UnAuthorized(c *gin.Context) {
	ResponseWithMsg(c, http.StatusNotFound, "用户不存在")
}

func Conflict(c *gin.Context) {
	ResponseWithMsg(c, http.StatusNotFound, "目标冲突")
}

func Success(c *gin.Context) {
	ResponseWithMsg(c, http.StatusOK, "ok")
}

type dataResponse struct {
	Data interface{}
}

func SuccessWithData(c *gin.Context, data interface{}) {
	res := dataResponse{Data: data}
	c.JSON(http.StatusOK, res)
}
