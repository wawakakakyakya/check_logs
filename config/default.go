package config

func NewDefaultLogConfig() LogConfig {
	return LogConfig{Level: 1, Path: "./check_logs_by_mail.log", MaxSize: 10, MaxBackups: 5, MaxAge: 7, Compress: true}
}
