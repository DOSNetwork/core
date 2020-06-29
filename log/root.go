package log

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/bshuster-repo/logrus-logstash-hook"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
)

var (
	root = &logger{}
)

type dosFormatter struct {
	logrus.TextFormatter
}

func (f *dosFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// this whole mess of dealing with ansi color codes is required if you want the colored output otherwise you will lose colors in the log levels
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 31 // gray
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	return []byte(fmt.Sprintf("[%s] - \x1b[%dm%s\x1b[0m - %s\n", entry.Time.Format(f.TimestampFormat), levelColor, strings.ToUpper(entry.Level.String()), entry.Message)), nil
}

// New returns a new logger with the given key/value .
// New is a convenient alias for Root().New
func New(key string, value interface{}) Logger {
	if root == nil {
		fmt.Println("no root")
	}
	return root.New(key, value)
}

// Init setups default field and add hook
func Init(id []byte) {
	appSession := os.Getenv("APPSESSION")
	logIp := os.Getenv("LOGIP")
	nodeId := byteTohex(id)
	os.Setenv("NODEID", nodeId)
	logrus.SetOutput(ioutil.Discard)

	if logIp != "" {
		logrus.SetLevel(logrus.DebugLevel)

		hook, err := logrustash.NewHook("tcp", logIp, appSession)
		if err != nil {
			fmt.Println(err)
			logrus.Error(err)
		}
		logrus.AddHook(hook)
	}

	path := "./vault/doslog.txt"
	writer, err := rotatelogs.New(
		path+".%Y-%m-%d-%H-%M-%S",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)
	if err != nil {
		fmt.Println("rotatelogs.New err", err)
	}

	logrus.AddHook(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  writer,
			logrus.DebugLevel: writer,
			logrus.ErrorLevel: writer,
		},
		&dosFormatter{
			logrus.TextFormatter{
				ForceColors:     true,
				TimestampFormat: "2006-01-02 15:04:05",
				FullTimestamp:   true,
			},
		},
	))
	root.entry = logrus.WithFields(logrus.Fields{})
}

func byteTohex(a []byte) string {
	unchecksummed := hex.EncodeToString(a[:])
	sha := sha3.NewLegacyKeccak256()
	sha.Write([]byte(unchecksummed))
	hash := sha.Sum(nil)

	result := []byte(unchecksummed)
	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}
	return "" + string(result)
}

// AddField is a convenient alias for Root().AddField
func AddField(key string, value interface{}) {
	root.AddField(key, value)
}

// Debug is a convenient alias for Root().Debug
func Debug(msg string) {
	root.Debug(msg)
}

// Info is a convenient alias for Root().Info
func Info(msg string) {
	root.Info(msg)
}

// Warn is a convenient alias for Root().Warn
func Warn(msg string) {
	root.Warn(msg)
}

// Error is a convenient alias for Root().Error
func Error(err error) {
	if root != nil {
		root.Error(err)
	}
}

// Fatal is a convenient alias for Root().Fatal
func Fatal(err error) {
	root.Fatal(err)
}

// Event is a convenient alias for Root().Event
func Event(e string, f map[string]interface{}) {
	root.Event(e, f)
}

// TimeTrack is a convenient alias for Root().TimeTrack
func TimeTrack(start time.Time, e string, info map[string]interface{}) {
	root.TimeTrack(start, e, info)
}
