package micro

import (
	"tochamateusz/micro/modules/logger"

	"go.uber.org/fx"
)

func Module(name string, opts ...fx.Option) fx.Option {
	var defaultOptions = []fx.Option{logger.Module(name)}
	defaultOptions = append(defaultOptions, opts...)
	return fx.Options(defaultOptions...)
}
