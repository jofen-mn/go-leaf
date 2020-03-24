package serilaizer

import (
	"github.com/gin-gonic/gin"
	"go-leaf/base/utils"
)

type ResponseData struct {
	RequestId string      `json:"request_id"`
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

func Response(c *gin.Context, error error, data interface{}) {
	var message string
	if error == nil {
		message = ""
	} else {
		message = error.Error()
	}
	c.JSON(0, ResponseData{
		RequestId: utils.Uuid().String(),
		Code:      0,
		Message:   message,
		Data:      data,
	})
}
