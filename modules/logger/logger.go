package logger

import (
	"go.uber.org/zap"
)

func New(config Config) (*zap.SugaredLogger, error) {
	cfg := zap.NewProductionConfig()
  //TODO: enviroment variable
	cfg.OutputPaths = append(cfg.OutputPaths, "/tmp/log")
	logger, err := cfg.Build()
	sugar := logger.Sugar()
	return sugar, err
}
