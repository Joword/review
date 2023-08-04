package app

import (
	"niftyreview/utils/log"

	"github.com/astaxie/beego/validation"
)

func MarkException(errors []*validation.Error) {
	for _, err := range errors {
		log.Info(err.Key, err.Message)
	}
	return
}
