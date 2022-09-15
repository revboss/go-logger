package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/polds/logrus-papertrail-hook"
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
