package logger

import "go.uber.org/zap"

func New() (*zap.SugaredLogger, error) {
	logger, err := zap.NewProduction()
	sugar := logger.Sugar()
	return sugar, err
}
