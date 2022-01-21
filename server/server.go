package server

import (
	"appcenter-wechat/conf"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewEngine() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	return r
}

func Run() error {
	srv := NewEngine()
	return srv.Run(fmt.Sprintf("%s:%d", conf.Appconf.Server.Addr, conf.Appconf.Server.Port))
}
