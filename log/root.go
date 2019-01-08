package log

import (
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

var (
	root = &logger{}
)

func New(key string, value interface{}) Logger {
	return root.New(key, value)
}

func init() {
	//IP,Subject and appName should read from environment variables
	hook, err := logrustash.NewHook("tcp", "163.172.36.173:9500", "peer")
	if err != nil {
		logrus.Error(err)
	}
	logrus.AddHook(hook)
	root.entry = logrus.WithFields(logrus.Fields{
		"subject": "subject",
		"appName": "appName",
	})
}

// Info is a convenient alias for Root().Info
func Info(msg ...interface{}) {
	root.Info(msg)
}

func Warn(msg ...interface{}) {
	root.Warn(msg)
}

func Error(err ...interface{}) {
	root.Error(err)
}

func Fatal(msg ...interface{}) {
	root.Fatal(msg)
}

// Info is a convenient alias for Root().Info
func Metrics(key string, value interface{}) {
	root.Metrics(key, value)
}
