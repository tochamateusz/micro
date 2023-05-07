package logger

import (
	"context"
	"os"

	"go.uber.org/fx"
)

func RegisterOnStop(lifecycle fx.Lifecycle, cfg Config) {
	lifecycle.Append(
		fx.Hook{
			OnStop: func(ctx context.Context) error {
				return os.Rename(cfg.OutputPath, cfg.OutputPath+"_old")
			},
		},
	)
}
