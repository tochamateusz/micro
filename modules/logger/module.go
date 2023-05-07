package logger

import "go.uber.org/fx"

//TODO: maybe this should return function
func Module(loggerName string) fx.Option {
	return fx.Module("Logger "+loggerName,
		fx.Provide(New),
	)
}
