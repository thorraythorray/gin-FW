package api

import (
	"github.com/gin-gonic/gin"
	"github.com/thorraythorray/go-proj/ginx/internal"
	"github.com/thorraythorray/go-proj/ginx/util/response"
	"github.com/thorraythorray/go-proj/pkg/auth"
)

func ObtainToken(c *gin.Context) {
	jwt := auth.JWT{
		SigningKey: internal.SignKey,
		CheckUser:  c.Request.Header.Get("X-User"),
		ExpireHour: internal.ExpireHour,
	}
	tokenstring, err := auth.AuthorizeImpl.Obtain(&jwt)
	if err != nil {
		response.ServerFailed(c, err.Error())
	}
	response.SuccessWithData(c, tokenstring)
}
