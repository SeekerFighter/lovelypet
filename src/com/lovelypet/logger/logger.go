package logger

import (
	"github.com/sirupsen/logrus"
	"lovelypet/src/com/lovelypet/constant"
	"os"
	"time"
)

const (
	LOGPATH = "../logs/"
)

var (
	Log = logrus.New()
	path = LOGPATH + time.Now().Format(constant.FORMAT) + "/"
)

func init() {
	_ = mkDir(path)
	hook := NewHook(path+"2.log")
	Log.AddHook(hook)
	// 设置日志格式为json格式
	//Log.SetFormatter(&logrus.JSONFormatter{})
	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	//Log.SetOutput(os.Stdout)
	// 设置日志级别为warn以上
	//log.SetLevel(log.WarnLevel)

	//path := "/var/log/go.log"
	//writer, _ := rotatelogs.New(
	//	path+".%Y%m%d%H%M",
	//	rotatelogs.WithLinkName(path),
	//	rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
	//	rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	//)
}

func mkDir(path string) error {
	if IsExist(path) {
		return nil
	}
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	_ = os.Chmod(path, os.ModePerm)
	return nil
}

//IsExist  判断文件夹/文件是否存在  存在返回 true
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func Print(level logrus.Level, fields logrus.Fields, args ...interface{}) {
	if fields == nil {
		fields = logrus.Fields{}
	}
	entry := Log.WithFields(fields)
	switch level {
	case logrus.PanicLevel:
		entry.Panic(args)
	case logrus.FatalLevel:
		entry.Fatal(args)
	case logrus.ErrorLevel:
		entry.Error(args)
	case logrus.WarnLevel:
		entry.Warn(args)
	case logrus.InfoLevel:
		entry.Info(args)
	case logrus.DebugLevel:
		entry.Debug(args)
	default:
	}
}
