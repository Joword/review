package main

import (
	"github.com/gin-gonic/gin"

	setting "niftyreview/config"
	models "niftyreview/models"
)

func init() {
	setting.Setup()
	models.Setup()
}

func main() {
	gin.SetMode("")
}
