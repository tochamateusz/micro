package logger

import (
	"fmt"
	"log"
	"os"

	"go.uber.org/fx"
)

func TestSpread(ints ...int) {
	fmt.Printf("%+v\n", ints)
}

var Module2 = fx.Provide(func() *log.Logger {
	return log.New(os.Stdout, "", 0)
})

// TODO: maybe this should return function
func Module(loggerName string) fx.Option {

	var options = []fx.Option{Module2, fx.Module("Logger "+loggerName,
		fx.Provide(
			func() Config {
				return Config{
					OutputPath: "/tmp/log",
				}
			}, New),
		fx.Invoke(
			RegisterOnStop,
		),
	)}

	return fx.Options(options...)
}
