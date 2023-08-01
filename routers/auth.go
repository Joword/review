package routers

import (
	"fmt"
	"net/http"
	"niftyreview/utils"
	"niftyreview/utils/result"

	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {

	appG := utils.Gin{C: c}

	username := c.PostForm("username")
	password := c.PostForm("password")

	fmt.Println(username + "|" + password)

	appG.Response(http.StatusAccepted, result.INVALID_PARAMS, nil)

}
