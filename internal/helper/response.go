package helper

import "github.com/gin-gonic/gin"

func SuccessHttpResponse(msg string, data interface{}) gin.H {
	return gin.H{
		"message": msg,
		"data":    data,
	}
}
