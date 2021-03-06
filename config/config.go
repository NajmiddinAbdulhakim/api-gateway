package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	UserServiceHost string
	UserServicePort int

	PostServiceHost string
	PostServicePort int

	// context timeout in seconds
	CtxTimeout int
	RedisHost string
	RedisPort int

	LogLevel string
	HTTPPort string

	SigninKey string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":1111"))
	c.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "127.0.0.1"))
	c.UserServicePort = cast.ToInt(getOrReturnDefault("USER_SERVICE_PORT", 9999))

	c.RedisHost=cast.ToString(getOrReturnDefault("REDIS_HOST","localhost"))
	c.RedisPort=cast.ToInt(getOrReturnDefault("REDIS_PORT",6379))

	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "127.0.0.1"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 2222))

	c.SigninKey = cast.ToString(getOrReturnDefault("SIGNIN_KEY", "tzkvwyuvywfczegoqclvuegvpkfwqkmqlarlxsscimnwonamoslueiendywvsolsiynsnlobnymlnuwpqmtzvbphhajlebrzlzpezdkhiaepevreisgcrhxkizrhcjrcmqrkkcgpdokbvanpfnocaqhugdfdquha"))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
