package sandbox

import (
	"github.com/tochamateusz/micro/modules/micro"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return micro.Module("Sandbox", fx.Provide(New))
}
