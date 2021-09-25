package logrus

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.Out = os.Stdout
	// 设置输出的等级
	Log.SetLevel(logrus.InfoLevel)
}
