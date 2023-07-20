package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data interface{}
	Msg  string
}

func HttpResponse(c *gin.Context, s int, obj interface{}) {
	c.JSON(s, obj)
}

func Success(c *gin.Context) {
	HttpResponse(c, http.StatusOK, "ok")
}

func SuccessWithContent(c *gin.Context, data interface{}, msg string) {
	res := Response{}
	if data != nil {
		res.Data = data
	}
	if msg == "" {
		res.Msg = "ok"
	} else {
		res.Msg = msg
	}
	HttpResponse(c, http.StatusOK, &res)
}

func NotFound(c *gin.Context) {
	HttpResponse(c, http.StatusNotFound, nil)
}

func AuthForbidden(c *gin.Context) {
	HttpResponse(c, http.StatusForbidden, nil)
}

func NeedAuthorized(c *gin.Context) {
	HttpResponse(c, http.StatusUnauthorized, nil)
}

func Failed(c *gin.Context, msg string) {
	res := Response{}
	res.Msg = msg
	HttpResponse(c, http.StatusBadRequest, &res)
}

func Conflict(c *gin.Context) {
	HttpResponse(c, http.StatusConflict, nil)
}
