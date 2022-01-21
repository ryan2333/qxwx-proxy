package main

import (
	"appcenter-wechat/conf"
	"appcenter-wechat/klog"
	"flag"
	"log"
)

var (
	file string
)

func init() {
	flag.StringVar(&file, "f", "conf/app.yaml", "config file path, default: conf/app.yaml")
}

func main() {
	flag.Parse()
	if err := conf.InitConf("conf/app.yaml"); err != nil {
		log.Fatal(err)
		return
	}

	if err := klog.InitLogs(); err != nil {
		log.Fatal(err)
		return
	}

}
