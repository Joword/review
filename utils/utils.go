package utils

import (
	"fmt"
	"niftyreview/models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func CheckAuth(username string, auth int) bool {
	var info models.UserMessage
	err := db.Table("user_info").First(&info).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		fmt.Println("The record can not found.")
	}
	if info.IsSuperuser > 0 {
		return true
	}
	return false
}
