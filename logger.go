package logger

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/heroku/rollrus"
	"github.com/polds/logrus-papertrail-hook"
	"github.com/stvp/roll"
	"strings"
)

func New(app string) *logrus.Logger {
	log := logrus.New()

	conf, e := LoadConfig()
	if e != nil {
		log.Fatal(e)
	}

	le := strings.ToLower(conf.Environment)
	if le == "production" || le == "staging" {
		log.Hooks.Add(&rollrus.Hook{
			Client: roll.New(conf.RollbarKey, conf.Environment),
		})

		phook, err := logrus_papertrail.NewPapertrailHook(&logrus_papertrail.Hook{
			Host:     conf.PapertrailHost,
			Port:     conf.PapertrailPort,
			Hostname: conf.Host,
			Appname:  app,
		})

		if err == nil {
			log.Hooks.Add(phook)
		} else {
			panic(fmt.Sprintf("Error: %+v", err))
		}
	}

	return log
}
