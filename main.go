package main

import (
	"github.com/gin-gonic/gin"

	setting "niftyreview/config"
)

func init() {
	setting.Setup()
}

func main() {
	gin.SetMode("")
}
