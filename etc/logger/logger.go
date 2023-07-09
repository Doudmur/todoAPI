package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"os"
)

var logger *logrus.Logger

func init() {
	logger = &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}
}

func Infof(ctx context.Context, str string, args ...interface{}) {
	logger.Infof(str, args)
}

func Errorf(ctx context.Context, str string, err error, args ...interface{}) {
	logger.WithError(err).Errorf(str, args)
}

func SetErrorLevel() {

}

func GetErrorLevel() {

}
