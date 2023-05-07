package sandbox

import (
	"tosiek88/micro/modules/logger"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Module("Sandbox",
		fx.Provide(New),
		logger.Module("Sandbox"),
	)

}
