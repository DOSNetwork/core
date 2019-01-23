package log

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
)

var (
	root = &logger{}
)

func New(key string, value interface{}) Logger {
	return root.New(key, value)
}

type UTCFormatter struct {
	logrus.Formatter
}

func (u UTCFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = time.Now()

	return u.Formatter.Format(e)
}

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(ioutil.Discard)
	logrus.SetFormatter(UTCFormatter{&logrus.JSONFormatter{}})
	//IP,Subject and appName should read from environment variables
	hook, err := logrustash.NewHook("tcp", "163.172.36.173:9500", "peer")
	if err != nil {
		logrus.Error(err)
	}
	logrus.AddHook(hook)
	appSubject := os.Getenv("appSubject")
	appName := os.Getenv("appName")
	root.entry = logrus.WithFields(logrus.Fields{
		"appSubject": appSubject,
		"appName":    appName,
	})
}

func AddField(key string, value interface{}) {
	root.AddField(key, value)
}

func Debug(msg string) {
	root.Debug(msg)
}

func Info(msg string) {
	root.Info(msg)
}

func Warn(msg string) {
	root.Warn(msg)
}

func Error(err error) {
	root.Error(err)
}

func Fatal(msg string) {
	root.Fatal(msg)
}

// Info is a convenient alias for Root().Info
func Metrics(value interface{}) {
	root.Metrics(value)
}

func Progress(progress string) {
	root.Progress(progress)
}

func Event(e string) {
	root.Event(e)
}
func Fields(f map[string]interface{}) {
	root.Fields(f)
}
