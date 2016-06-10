package logger

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Environment    string
	Host           string
	RollbarKey     string
	PapertrailHost string
	PapertrailPort int
}

func LoadConfig() (Config, error) {
	c := Config{}

	envs := map[string]interface{}{
		"ENVIRONMENT":     &c.Environment,
		"HOST":            &c.Host,
		"ROLLBAR_KEY":     &c.RollbarKey,
		"PAPERTRAIL_HOST": &c.PapertrailHost,
		"PAPERTRAIL_PORT": &c.PapertrailPort,
	}

	var isProduction bool

	for k, cv := range envs {
		v := os.Getenv(k)
		switch cv.(type) {
		case *string:
			if k == "ENVIRONMENT" && strings.ToLower(v) == "production" {
				isProduction = true
			}

			if len(v) == 0 {
				if k == "HOST" {
					var err error
					v, err = os.Hostname()
					if nil != err {
						return c, fmt.Errorf("Unable to retrieve default hostname: %v", err)
					}
				} else {
					return c, fmt.Errorf("Missing required configuration value %s", k)
				}
			}

			*cv.(*string) = v
		case *int:
			pv, _ := strconv.Atoi(v)
			if pv == 0 && !(k == "LOG_PORT" && isProduction == false) {
				return c, fmt.Errorf("Missing required configuration value %s", k)
			}
			*cv.(*int) = pv
		}
	}

	return c, nil
}
