package jwt

import (
	"net/http"
	result "niftyreview/utils/result"

	"github.com/gin-gonic/gin"
)

func handlerFuc(c *gin.Context) {
	var code int
	token := c.Query("token")

	code = result.SUCCESS
	if token == "" {
		code = result.INVALID_PARAMS
	}

	if code != result.SUCCESS {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": code, "msg": "", "data": "",
		})
	}
	c.Next()
}

func JsonWebToken() gin.HandlerFunc {

	return handlerFuc
}
