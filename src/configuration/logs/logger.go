package logs

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var(
	log *zap.Logger

	LOG_OUTPUT = "LOG_OUTPUT"
	LOG_LEVEL = "LOG_LEVEL"
)

func init() {
	logConfig := zap.Config{
		OutputPaths: []string{GetOutPutLogs()},
		Level: zap.NewAtomicLevelAt(GetLevelLogs()),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey: "level",
			TimeKey: "time",
			MessageKey: "message",
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	log, _ = logConfig.Build()
}

func GetOutPutLogs() string{
	output := strings.ToLower(strings.TrimSpace(os.Getenv(LOG_OUTPUT)))
	if output == "" {
		return "stdout"
	}
	return output
}

func GetLevelLogs() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(LOG_LEVEL))){
		
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zap.ErrorLevel
	case "debug":
		return zapcore.DebugLevel
	default:
		return zapcore.InfoLevel
	}
}

func Info(message string, tags...zap.Field){
	log.Info(message, tags...)
	log.Sync()
}

func Error(message string, err error, tags...zap.Field){
	tags = append(tags, zap.NamedError("error", err))
	log.Info(message, tags...)
	log.Sync()
}