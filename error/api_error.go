package error

import "github.com/gin-gonic/gin"

func ErrorHttpResponse(err error, msg string) gin.H {
	return gin.H{
		"error":   err.Error(),
		"message": msg,
	}
}
