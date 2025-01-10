package config

const (
	PRODUCTION = "production"
	DEV        = "dev"
	STAGING    = "staging"
)

var (
	ENV             = getEnvString("ENV", DEV)
	SERVICE_NAME    = "videoverse"
	BACKEND_VERSION = "2.0.0"
	LOGGING_FILE    = getEnvString("LOGGING_FILE", "logs/api.log")
	APP_PORT        = getEnvInt("APP_PORT", 9091)
	JWT_SECRET      = getEnvString("JWT_SECRET", "secret")
)
