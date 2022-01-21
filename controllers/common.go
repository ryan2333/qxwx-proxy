package controllers

import (
	"appcenter-wechat/conf"
	"appcenter-wechat/tools"
	"fmt"
)

func CheckChatServerIp(ip string) bool {

	return false
}

func CheckAuth(sysname, token string) bool {
	fullToken := fmt.Sprint(999, token, 2333)
	return tools.GeneratorMd5(fullToken) == conf.Appconf.Token[sysname].Token
}
