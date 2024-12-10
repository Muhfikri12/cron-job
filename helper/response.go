package helper

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool
	Message string
	Data    any
}

func Responses(c *gin.Context, status int, massage string, data any) {
	responseStatus := data != nil
	c.JSON(status, Response{
		Status:  responseStatus,
		Message: massage,
		Data:    data,
	})
}
