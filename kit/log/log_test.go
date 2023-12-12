package log_test

import (
	"testing"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/hachimi-lab/re/kit/log"
)

func init() {
	logger := log.New(
		log.WithDevelopment(false),
		log.WithLevel(log.InfoLevel),
		log.WithDisableStacktrace(false),
	).Store(
		&lumberjack.Logger{
			Filename:   "./logs/re.log",
			MaxSize:    100,
			MaxAge:     7,
			MaxBackups: 14,
			LocalTime:  true,
			Compress:   true,
		},
	).Extend(
		log.StdoutWriteSyncer(),
	).Build(
		log.BuildWithHooks(func(entry log.Entry) error {
			return nil
		}),
	).Named("TEST")

	log.SetLogger(logger)
}

func TestLog(t *testing.T) {
	logger := log.GetLogger()

	log.Debug("hi, re!")
	log.Debugf("hi, %s!", "re")

	log.Info("hi, re!")
	logger.Info("hi, re!")

	log.Error("hi, re!", log.Int64("timestamp", time.Now().Unix()))
	log.Fatal("hi, re!", log.Int64("timestamp", time.Now().Unix()))

	for {
		log.Error("hi, re!")
	}
}
