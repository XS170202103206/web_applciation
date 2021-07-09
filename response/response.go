package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SuccessRes(c *gin.Context, msg interface{}, data interface{}) {
	c.JSON(http.StatusOK,gin.H{
		"msg" : msg,
		"data" : data,
	})
}

func ErrorRes(c *gin.Context, httpCode int, msg interface{}, err interface{}) {
	c.JSON(httpCode,gin.H{
		"msg" : msg,
		"error" : err,
	})
}