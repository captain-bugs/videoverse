package config

const (
	PRODUCTION = "production"
	DEV        = "dev"
)

var (
	ENV                = getEnvString("ENV", DEV)
	SERVICE_NAME       = "videoverse"
	BACKEND_VERSION    = "2.0.0"
	LOGGING_FILE       = getEnvString("LOGGING_FILE", "logs/api.log")
	APP_PORT           = getEnvInt("APP_PORT", 9091)
	JWT_SECRET         = getEnvString("JWT_SECRET", "secret")
	DATABASE_PATH      = getEnvString("DATABASE_PATH", "db/videoverse/videoverse.db")
	MIN_VIDEO_DURATION = getEnvFloat("MIN_VIDEO_DURATION", 5.0)
	MAX_VIDEO_DURATION = getEnvFloat("MAX_VIDEO_DURATION", 125.0)
	FILE_UPLOAD_PATH   = getEnvString("FILE_UPLOAD_PATH", "uploads/videos")
	SHARE_SECRET       = getEnvString("JWT_SECRET", "secret2")
	CDN_ENDPOINT       = getEnvString("CDN_ENDPOINT", "http://localhost:9091")
)
