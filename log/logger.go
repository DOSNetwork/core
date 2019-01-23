package log

import (
	"runtime"
	"strings"

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
	AddField(key string, value interface{})
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(err error)
	Fatal(msg string)
	Metrics(value interface{})
	Progress(progress string)
	Event(e string)
	Fields(f map[string]interface{})
}

type logger struct {
	entry *logrus.Entry
}

func (l *logger) New(key string, value interface{}) Logger {
	return &logger{l.entry.WithFields(logrus.Fields{key: value})}
}

func (l *logger) AddField(key string, value interface{}) {
	l.entry = l.entry.WithFields(logrus.Fields{key: value})
}

func (l *logger) Debug(msg string) {
	l.entry.Debug(msg)
}

func (l *logger) Info(msg string) {
	l.entry.Info(msg)
}

func (l *logger) Warn(msg string) {
	l.entry.Warn(msg)
}

func (l *logger) Error(err error) {
	s := stack.Trace().TrimRuntime()
	l.entry.WithFields(logrus.Fields{"stack": s}).Error(err)
}

func (l *logger) Fatal(msg string) {
	l.entry.Fatal(msg)
}

func (l *logger) Metrics(value interface{}) {

	fpcs := make([]uintptr, 1)
	// Skip 2 levels to get the caller
	n := runtime.Callers(2, fpcs)
	if n == 0 {
		l.entry.WithFields(logrus.Fields{"M_": value}).Info("")
	} else {
		caller := runtime.FuncForPC(fpcs[0] - 1)
		s := strings.Split(caller.Name(), ".")
		l.entry.WithFields(logrus.Fields{"M_" + s[len(s)-1]: value}).Debug("")
	}
}

func (l *logger) Progress(progress string) {
	l.entry.WithFields(logrus.Fields{"Progress": progress}).Debug("")
}

func (l *logger) Event(e string) {
	fpcs := make([]uintptr, 1)
	// Skip 3 levels to get the caller
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		l.entry.WithFields(logrus.Fields{"EVENT": e}).Info("")
	} else {
		caller := runtime.FuncForPC(fpcs[0] - 1)
		s := strings.Split(caller.Name(), ".")
		l.entry.WithFields(logrus.Fields{"EVENT": e + s[len(s)-1]}).Debug("")
	}
}

func (l *logger) Fields(f map[string]interface{}) {
	var a logrus.Fields
	a = f
	l.entry.WithFields(a).Debug("")
}
