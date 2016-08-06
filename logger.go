package logger

import (
	"github.com/Sirupsen/logrus"
	"github.com/heroku/rollrus"
	"github.com/polds/logrus-papertrail-hook"
	"github.com/revboss/go-config"
	"github.com/stvp/roll"
	"strings"
)

var levels = map[string]logrus.Level{
	"debug": logrus.DebugLevel,
	"info":  logrus.InfoLevel,
	"warn":  logrus.WarnLevel,
	"error": logrus.ErrorLevel,
	"fatal": logrus.FatalLevel,
	"panic": logrus.PanicLevel,
}

type Config struct {
	Environment    string
	Host           string
	RollbarKey     string
	PapertrailHost string
	PapertrailPort int
	Level          string
}

func New(app string) *logrus.Logger {
	log := logrus.New()

	conf := &Config{}
	e := config.LoadConfig(config.Config{
		Env: map[string]config.Value{
			"ENVIRONMENT": {
				Default: "development",
				Value:   &conf.Environment,
			},
			"LOG_LEVEL": {
				Default: "info",
				Value:   &conf.Level,
			},
		},
	})
	if e != nil {
		log.Fatal(e)
	}

	conf.Level = strings.ToLower(conf.Level)
	if level, ok := levels[conf.Level]; ok {
		log.Level = level
	} else {
		log.Warnf("Invalid log level %q, using %q", conf.Level, log.Level)
	}

	le := strings.ToLower(conf.Environment)
	if le == "production" || le == "staging" {
		e := config.LoadConfig(config.Config{
			Env: map[string]config.Value{
				"HOST": {
					Value: &conf.Host,
				},
				"ROLLBAR_KEY": {
					Value: &conf.RollbarKey,
				},
				"PAPERTRAIL_HOST": {
					Default: "logs2.papertrailapp.com",
					Value:   &conf.PapertrailHost,
				},
				"PAPERTRAIL_PORT": {
					Default: "33263",
					Value:   &conf.PapertrailPort,
				},
			},
		})
		if e != nil {
			log.Fatal(e)
		}

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
			log.Hooks.Add(PapertrailHook{phook})
		} else {
			log.WithError(err).Panic("Failed to add Papertrail hook")
		}
	}

	return log
}
