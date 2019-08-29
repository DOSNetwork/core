package log

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	//errors "golang.org/x/xerrors"
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
	TimeTrack(time.Time, string, map[string]interface{})
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
	ss := strings.FieldsFunc(fmt.Sprintf("%+v", err), func(r rune) bool {
		if r == '-' {
			return true
		}
		return false
	})
	errCause := ""
	if len(ss) >= 1 {
		c := strings.Split(ss[len(ss)-1], ":")
		errCause = c[0]
	}
	//
	l.entry.WithFields(logrus.Fields{"errCause": errCause, "errDetail": fmt.Sprintf("%+v", err), "errType": reflect.TypeOf(err).String(), "Time": time.Now()}).Error()

}

func (l *logger) Fatal(err error) {
	if l.entry == nil {
		os.Exit(1)
		return
	}
	l.entry.Fatal(err)
	os.Exit(1)
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
