package route

import (
	"appcenter-wechat/controllers"
	"appcenter-wechat/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine) {

	wn := r.Group("/wechat")
	{
		wn.POST("/", controllers.SendMessageController)
		wn.POST("/stark/callback", controllers.ReceiveStarkCallBackController)
		wn.POST("/callback", controllers.ReceiveCallBackController)
	}

	wn.Use(middleware.Auth)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "healthcheck"})
	})
}
