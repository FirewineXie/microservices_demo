package jaegerc

import (
	"go.uber.org/zap"
)

type zapLogger struct {
	logger *zap.Logger
}

func (l *zapLogger) Error(msg string) {
	l.logger.Error(msg)
}

// Infof logs a message at info priority
func (l *zapLogger) Infof(msg string, args ...interface{}) {
	l.logger.Sugar().Infof(msg, args...)

}

// Debugf logs a message at debug priority
func (l *zapLogger) Debugf(msg string, args ...interface{}) {
	l.logger.Info("DEBUG :"+msg, zap.Any("args", args))
}
