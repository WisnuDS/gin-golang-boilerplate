package configs

import (
	"os"
)

type Env struct {
	ServerPort    string
	Environment   string
	SentryDSN     string
	Schema        string
	DBHost        string
	DBName        string
	GinMode       string
	AMQPServerurl string
	ApiKey        string
	PublicRoute   string
}

func InitEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

func (env *Env) LoadEnv() {
	env.GinMode = os.Getenv("GIN_MODE")
	env.ServerPort = os.Getenv("SERVER_PORT")
	env.Environment = os.Getenv("ENV")
	env.Schema = os.Getenv("SCHEMA")
	env.SentryDSN = os.Getenv("SENTRY_DSN")
	env.DBHost = os.Getenv("DB_HOST")
	env.DBName = os.Getenv("DB_NAME")
	env.AMQPServerurl = os.Getenv("AMQP_SERVER_URL")
	env.ApiKey = os.Getenv("CRM_API_KEY")
	env.PublicRoute = os.Getenv("PUBLIC_ROUTE")
}
