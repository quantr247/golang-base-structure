package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// InitZap inits zap logging
func InitZap(app, env string) error {
	logLevel := configLogLevel(env)
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(logLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.ISO8601TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,

			NameKey:    "app",
			EncodeName: zapcore.FullNameEncoder,
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		return err
	}
	logger = logger.Named(app)
	zap.ReplaceGlobals(logger)
	return nil
}

func configLogLevel(defaultEnv string) zapcore.Level {
	// get from os env
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = defaultEnv
	}
	if env == "" {
		env = "P"
	}

	level := zapcore.ErrorLevel
	switch env {
	case "D", "d", "dev", "DEV":
		level = zapcore.DebugLevel
	case "P", "p", "PROD", "prod":
		level = zapcore.WarnLevel
	}

	logLevelEnv := os.Getenv("LOG_LEVEL")
	switch logLevelEnv {
	case "WARN", "warn":
		level = zapcore.WarnLevel
	case "DEBUG", "debug":
		level = zapcore.DebugLevel
	case "INFO", "info":
		level = zapcore.InfoLevel
	}
	return level
}
