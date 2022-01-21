package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func SendMessageController(c *gin.Context) {
	var (
		sysName string
		err     error
	)
	fmt.Println(sysName, err)

}

func ReceiveCallBackController(c *gin.Context) {

}

func ReceiveStarkCallBackController(c *gin.Context) {}
