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

func New(key string, value interface{}) Logger {
	if root == nil {
		fmt.Println("root nit")
	}
	return root.New(key, value)
}

type UTCFormatter struct {
	logrus.Formatter
}

func (u UTCFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = time.Now()

	return u.Formatter.Format(e)
}

func Init(id []byte) {
	fmt.Println("lgo Init")
	appSession := os.Getenv("APPSESSION")
	appName := os.Getenv("APPNAME")
	nodId := byteTohex(id)
	logIp := os.Getenv("LOGIP")
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(ioutil.Discard)
	logrus.SetFormatter(UTCFormatter{&logrus.JSONFormatter{}})
	//IP,Subject and appName should read from environment variables
	hook, err := logrustash.NewHook("tcp", logIp, appSession)
	if err != nil {
		fmt.Println(err)
		logrus.Error(err)
	}
	logrus.AddHook(hook)
	root.entry = logrus.WithFields(logrus.Fields{
		"appSession": appSession,
		"appName":    appName,
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
	return "0x" + string(result)
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

func Fatal(err error) {
	root.Fatal(err)
}

// Info is a convenient alias for Root().Info
func Metrics(value interface{}) {
	root.Metrics(value)
}

func Progress(progress string) {
	root.Progress(progress)
}

func Event(e string, f map[string]interface{}) {
	root.Event(e, f)
}
