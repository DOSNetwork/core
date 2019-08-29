package log

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
)

var (
	root = &logger{}
)

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
	appName := os.Getenv("APPNAME")
	nodId := byteTohex(id)
	logIp := os.Getenv("LOGIP")
	clientIP := os.Getenv("NODEIP")
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(ioutil.Discard)
	//IP,Subject and appName should read from environment variables

	if logIp != "" {
		hook, err := logrustash.NewHook("tcp", logIp, appSession)
		if err != nil {
			fmt.Println(err)
			logrus.Error(err)
		}
		logrus.AddHook(hook)
	}
	root.entry = logrus.WithFields(logrus.Fields{
		"appSession": appSession,
		"appName":    appName,
		"clientip":   clientIP,
		"nodeID":     nodId,
	})
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
