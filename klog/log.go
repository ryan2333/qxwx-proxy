package klog

import (
	conf "appcenter-wechat/conf"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is the log printer
var Logger *zap.SugaredLogger

var LogLevels = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
}

// InitLogs init the logging system.
// In development mode,
func InitLogs() (err error) {
	logLevel := zapcore.InfoLevel
	if v, ok := LogLevels[conf.Appconf.Log.LogDir]; ok {
		logLevel = v
	}

	// detect log output levels
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= logLevel
	})
	errorLogLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})

	var infoLogWriter io.Writer
	var errorLogWriter io.Writer
	var logEncoder zapcore.Encoder

	// create log writers
	infoLogWriter, err = getRotateWriter("info")
	if err != nil {
		err = fmt.Errorf("create info log writer error, %s", err.Error())
		return
	}
	errorLogWriter, err = getRotateWriter("error")
	if err != nil {
		err = fmt.Errorf("create error log writer error, %s", err.Error())
		return
	}

	// setup the encoder config and options
	logEncoderConfig := zapcore.EncoderConfig{
		MessageKey: "msg",
		LevelKey:   "level",
		TimeKey:    "time",
		CallerKey:  "file",
		EncodeTime: func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	}
	logCoreOptions := make([]zap.Option, 0)
	if conf.Appconf.Log.JsonEncode {
		logEncoderConfig.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(level.String())
		}
		logEncoder = zapcore.NewJSONEncoder(logEncoderConfig)
	} else {
		logCoreOptions = append(logCoreOptions, zap.AddCaller())
		logEncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		logEncoder = zapcore.NewConsoleEncoder(logEncoderConfig)
	}

	// pack up output cores
	logOutputList := []zapcore.Core{
		zapcore.NewCore(logEncoder, zapcore.AddSync(infoLogWriter), infoLevel),
		zapcore.NewCore(logEncoder, zapcore.AddSync(errorLogWriter), errorLogLevel)}
	if conf.Appconf.Log.Console {
		logOutputList = append(logOutputList, zapcore.NewCore(logEncoder, zapcore.AddSync(os.Stdout), infoLevel))
	}
	logCore := zapcore.NewTee(logOutputList...)
	Logger = zap.New(logCore, logCoreOptions...).Sugar()
	return
}

func getRotateWriter(fileName string) (w io.Writer, err error) {
	logFilePath := filepath.Join(conf.Appconf.Log.LogDir, fmt.Sprint(fileName, "-", "%Y-%m-%d", ".log"))
	//logLinkPath := filepath.Join(conf.AppConfig.Logging.OutputLocalDir, fileName)
	w, err = rotatelogs.New(logFilePath,
		//rotatelogs.WithLinkName(logLinkPath),
		rotatelogs.WithMaxAge(time.Hour*24*14),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	return
}
