package middleware

import (
	"appcenter-wechat/controllers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	var (
		sysName string
		token   string
	)

	sysName = c.GetHeader("sysname")
	token = c.GetHeader("token")

	if !controllers.CheckAuth(sysName, token) {
		c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("invalid token"))
	}
	c.Next()

}
