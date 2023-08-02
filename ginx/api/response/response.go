package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseWithMsg(c *gin.Context, s int, msg string) {
	res := map[string]string{"msg": msg}
	c.JSON(s, res)
}

func Success(c *gin.Context) {
	ResponseWithMsg(c, http.StatusOK, "ok")
}

func SuccessWithData(c *gin.Context, data interface{}) {
	res := map[string]interface{}{"data": data}
	c.JSON(http.StatusOK, res)
}

func RequestFailed(c *gin.Context, err error) {
	ResponseWithMsg(c, http.StatusBadRequest, err.Error())
}

func ServerFailed(c *gin.Context, err error) {
	ResponseWithMsg(c, http.StatusInternalServerError, err.Error())
}

func NotFound(c *gin.Context) {
	ResponseWithMsg(c, http.StatusNotFound, "请求的资源不存在")
}

func AuthForbidden(c *gin.Context) {
	ResponseWithMsg(c, http.StatusForbidden, "权限禁止")
}

func UnAuthorized(c *gin.Context) {
	ResponseWithMsg(c, http.StatusUnauthorized, "未经授权，需要身份验证")
}

func Conflict(c *gin.Context) {
	ResponseWithMsg(c, http.StatusConflict, "资源存在，冲突")
}
