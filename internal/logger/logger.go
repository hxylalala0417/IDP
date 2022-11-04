package logger

import (
	"IDP/internal/utils"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.SugaredLogger

//Init a logger, logLevel stands for the lowest error level that going to appear inside the log
func Init(logFilePath string, logLevel zapcore.Level) error {
	const LOG_FOLDER string = "./logs"
	if Logger != nil {
		Sync()
	}
	if !utils.Exists(LOG_FOLDER) {
		err := os.Mkdir(LOG_FOLDER, 0777)
		if err != nil {
			return errors.New("create log folder error")
		}
	}
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{LOG_FOLDER + "/" + logFilePath}
	cfg.Level = zap.NewAtomicLevelAt(logLevel)
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	if logger, err := cfg.Build(); err != nil {
		return err
	} else {
		Logger = logger.Sugar()
	}
	return nil
}

// Sync flushes any buffered log entries.
func Sync() {
	if Logger != nil {
		err := Logger.Sync()
		if err != nil {
			fmt.Println("log flush error", err)
		}
	}
}
