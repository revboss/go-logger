package logger

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/polds/logrus-papertrail-hook.v3"
)

type PapertrailHook struct {
	*logrus_papertrail.Hook
}

func (ph PapertrailHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
	}
}
