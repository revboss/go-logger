package logger

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/heroku/rollrus"
	"github.com/polds/logrus-papertrail-hook"
	"github.com/revboss/go-config"
	"github.com/stvp/roll"
	"strings"
)

type Config struct {
	Environment    string
	Host           string
	RollbarKey     string
	PapertrailHost string
	PapertrailPort int
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
		},
	})
	if e != nil {
		log.Fatal(e)
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
			log.Hooks.Add(phook)
		} else {
			panic(fmt.Sprintf("Error: %+v", err))
		}
	}

	return log
}
