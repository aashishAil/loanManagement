package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Log *logger

type Logger interface {
	ShutDownLogger() error
	Info(message string, args ...Field)
	Infof(message string, args ...interface{})
	Error(message string, args ...Field)
	Errorf(message string, args ...interface{})
	Warn(message string, args ...Field)
	Warnf(message string, args ...interface{})
	Debug(message string, args ...Field)
}

type logger struct {
	zapLogger     *zap.Logger
	sugaredLogger *zap.SugaredLogger
}

func Init(isDevelopment bool) {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	logLevel := zap.InfoLevel
	if isDevelopment {
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
	Log = &logger{
		zapLogger:     zapLogger,
		sugaredLogger: zapLogger.Sugar(),
	}
}

func (l *logger) ShutDownLogger() error {
	return Log.zapLogger.Sync()
}

func (l *logger) Info(message string, args ...Field) {
	Log.sugaredLogger.Info(message, args)
}

func (l *logger) Infof(message string, args ...interface{}) {
	Log.sugaredLogger.Infof(message, args)
}

func (l *logger) Error(message string, args ...Field) {
	Log.sugaredLogger.Error(message, args)
}

func (l *logger) Errorf(message string, args ...interface{}) {
	Log.sugaredLogger.Errorf(message, args)
}

func (l *logger) Warn(message string, args ...Field) {
	Log.sugaredLogger.Warn(message, args)
}

func (l *logger) Warnf(message string, args ...interface{}) {
	Log.sugaredLogger.Warnf(message, args)
}

func (l *logger) Debug(message string, args ...Field) {
	Log.sugaredLogger.Debug(message, args)
}

func (l *logger) Debugf(message string, args ...interface{}) {
	Log.sugaredLogger.Debugf(message, args)
}
