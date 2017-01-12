package logger

import (
	"github.com/polds/logrus-papertrail-hook"
	"github.com/sirupsen/logrus"
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
