package sandbox

import (
	"tochamateusz/micro/modules/micro"

	"go.uber.org/fx"
)

func Module() fx.Option {
	return micro.Module("Sandbox", fx.Provide(New))
}
