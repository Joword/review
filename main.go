package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	setting "niftyreview/config"
	models "niftyreview/models"
	"niftyreview/routers"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	gin.SetMode("")

	routerInit := routers.InitRouter()

	server := &http.Server{
		Handler: routerInit,
	}

	server.ListenAndServe()
}
