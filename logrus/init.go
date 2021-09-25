package logrus

import (
	"fmt"
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Log *logrus.Logger

//
func init() {
	filePath := "logrus/log"
	src, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err :", err)
	}
	Log = logrus.New()
	Log.Out = src
	// 设置输出的等级
	Log.SetLevel(logrus.InfoLevel)
	// 重定向输出文件
	logWriter, _ := rotate.New(
		// 加输出的时间
		filePath+"%Y%m%d.Log",
		rotate.WithMaxAge(7*24*time.Hour),
		rotate.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		// 123456 固定格式
		TimestampFormat: "2006-01-02 15:04:05",
	})
	Log.AddHook(Hook)
}
