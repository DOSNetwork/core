package log

import (
	"os"
	"runtime"
	"strings"
	"time"

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
	Fatal(err error)
	Metrics(value interface{})
	TimeTrack(time.Time, string, map[string]interface{})
	Progress(progress string)
	Event(e string, f map[string]interface{})
}

type logger struct {
	entry *logrus.Entry
}

func (l *logger) New(key string, value interface{}) Logger {
	if l.entry != nil {
		return &logger{l.entry.WithFields(logrus.Fields{key: value})}
	}
	return nil
}

func (l *logger) AddField(key string, value interface{}) {
	if l.entry == nil {
		return
	}
	l.entry = l.entry.WithFields(logrus.Fields{key: value})
}

func (l *logger) Debug(msg string) {
	if l.entry == nil {
		return
	}
	l.entry.Debug(msg)
}

func (l *logger) Info(msg string) {
	if l.entry == nil {
		return
	}
	l.entry.Info(msg)
}

func (l *logger) Warn(msg string) {
	if l.entry == nil {
		return
	}
	l.entry.Warn(msg)
}

func (l *logger) Error(err error) {
	if l.entry == nil {
		return
	}
	s := stack.Trace().TrimRuntime()
	l.entry.WithFields(logrus.Fields{"Stack": s, "ErrMsg": err}).Error(err)
}

func (l *logger) Fatal(err error) {
	if l.entry == nil {
		os.Exit(1)
		return
	}
	l.entry.Fatal(err)
	os.Exit(1)
}

func (l *logger) Metrics(value interface{}) {
	if l.entry == nil {
		return
	}
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
	if l.entry == nil {
		return
	}
	l.entry.WithFields(logrus.Fields{"Progress": progress}).Debug("")
}

func (l *logger) Event(e string, info map[string]interface{}) {
	if l.entry == nil {
		return
	}
	if info != nil {
		l.entry.WithFields(logrus.Fields{"EVENT": e, "Time": time.Now()}).WithFields(info).Debug("")
	} else {
		l.entry.WithFields(logrus.Fields{"EVENT": e, "Time": time.Now()}).Debug("")

	}
}

func (l *logger) TimeTrack(start time.Time, e string, info map[string]interface{}) {
	elapsed := time.Since(start).Nanoseconds() / 1000

	if l.entry == nil {
		return
	}
	if info != nil {
		l.entry.WithFields(logrus.Fields{"EVENT": e, e: elapsed, "Time": time.Now()}).WithFields(info).Debug("")
	} else {
		l.entry.WithFields(logrus.Fields{"EVENT": e, e: elapsed, "Time": time.Now()}).Debug("")
	}
}
