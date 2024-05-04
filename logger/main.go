package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"loanManagement/config"
	"os"
)

var Logger *internalLogger

type InternalLogger interface {
	ShutDownLogger() error
	Info(message string, args ...Field)
	Infof(message string, args ...interface{})
	Error(message string, args ...Field)
	Errorf(message string, args ...interface{})
	Warn(message string, args ...Field)
	Warnf(message string, args ...interface{})
	Debug(message string, args ...Field)
}

type internalLogger struct {
	zapLogger     *zap.Logger
	sugaredLogger *zap.SugaredLogger
}

func init() {
	envConfig := config.Env
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	logLevel := zap.InfoLevel
	if envConfig.IsDevelopment() {
		logLevel = zap.DebugLevel
	}

	loggerConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(logLevel),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling:          nil,
		Encoding:          "json",
		EncoderConfig:     encoderCfg,
		OutputPaths: []string{
			"stderr",
		},
		ErrorOutputPaths: []string{
			"stderr",
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}
	zapLogger := zap.Must(loggerConfig.Build())
	Logger = &internalLogger{
		zapLogger:     zapLogger,
		sugaredLogger: zapLogger.Sugar(),
	}
}

func (l *internalLogger) ShutDownLogger() error {
	return Logger.zapLogger.Sync()
}

func (l *internalLogger) Info(message string, args ...Field) {
	Logger.sugaredLogger.Info(message, args)
}

func (l *internalLogger) Infof(message string, args ...interface{}) {
	Logger.sugaredLogger.Infof(message, args)
}

func (l *internalLogger) Error(message string, args ...Field) {
	Logger.sugaredLogger.Error(message, args)
}

func (l *internalLogger) Errorf(message string, args ...interface{}) {
	Logger.sugaredLogger.Errorf(message, args)
}

func (l *internalLogger) Warn(message string, args ...Field) {
	Logger.sugaredLogger.Warn(message, args)
}

func (l *internalLogger) Warnf(message string, args ...interface{}) {
	Logger.sugaredLogger.Warnf(message, args)
}

func (l *internalLogger) Debug(message string, args ...Field) {
	Logger.sugaredLogger.Debug(message, args)
}

func (l *internalLogger) Debugf(message string, args ...interface{}) {
	Logger.sugaredLogger.Debugf(message, args)
}
