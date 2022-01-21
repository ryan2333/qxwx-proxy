package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type appConf struct {
	Server serverConf           `yaml:"server"`
	Log    logConf              `yaml:"log"`
	Token  map[string]tokenInfo `yaml:"token"`
}

type serverConf struct {
	Addr string `yaml:"addr"`
	Port int32  `yaml:"port"`
}

type logConf struct {
	LogDir     string `yaml:"logDir"`
	MaxSize    int    `yaml:"maxSize"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAge     int    `yaml:"maxAge"`
	Compress   bool   `yaml:"compress"`

	// logger options
	LogLevel        string `yaml:"logLevel"`
	JsonEncode      bool   `yaml:"jsonEncode"`
	StacktraceLevel string `yaml:"stackTraceLevel"`
	Console         bool   `yaml:"console"`
}

type tokenInfo struct {
	Token    string `yaml:"token"`
	CallBack string `yaml:"callback"`
}

var Appconf appConf

func InitConf(file string) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return fmt.Errorf("open config file failed, err: %v", err.Error())
	}

	if err := yaml.Unmarshal(data, &Appconf); err != nil {
		return fmt.Errorf("yaml unmarshal failed, err: %v", err.Error())
	}

	fmt.Println("111: ", Appconf.Log)
	return nil
}
