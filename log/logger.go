package log

import (
	"github.com/go-stack/stack"
	"github.com/sirupsen/logrus"
)

/*
Example:

import (
	"errors"
	"time"
	log "github.com/DOSNetwork/core/log"
)

func myfunc() error {
	return errors.New("error message")
}
func main() {
	startTime := time.Now()
	log.Info("test")

	time.Sleep(500 * time.Millisecond)
	log.Metrics("time-cost", time.Since(startTime).Seconds())

	srvlog := log.New("module", "service")
	srvlog.Error(myfunc())
}
*/

// A Logger writes key/value pairs to a Handler
type Logger interface {
	// New returns a new Logger that has this logger's context plus the given context
	New(key string, value interface{}) Logger

	Metrics(key string, value interface{})
	Info(msg ...interface{})
	Warn(msg ...interface{})
	Error(err ...interface{})
	Fatal(msg ...interface{})
}

type logger struct {
	entry *logrus.Entry
}

func (l *logger) New(key string, value interface{}) Logger {

	child := &logger{l.entry.WithFields(logrus.Fields{key: value})}
	return child
}

func (l *logger) Info(msg ...interface{}) {
	l.entry.Info(msg)
}

func (l *logger) Warn(msg ...interface{}) {
	l.entry.Warn(msg)
}

func (l *logger) Error(err ...interface{}) {
	s := stack.Trace().TrimRuntime()
	l.entry.WithFields(logrus.Fields{"stack": s}).Error(err)
}

func (l *logger) Fatal(msg ...interface{}) {
	l.entry.Fatal(msg)
}

func (l *logger) Metrics(key string, value interface{}) {
	l.entry.WithFields(logrus.Fields{key: value}).Info("")
}
