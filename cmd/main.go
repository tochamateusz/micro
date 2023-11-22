package main

import (
	"fmt"

	"github.com/tochamateusz/micro/modules/database"
	"github.com/tochamateusz/micro/modules/sandbox"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		sandbox.Module(),
		database.Module(),
		// amqp.Module(),
		fx.Invoke(
			func(s *sandbox.Sandbox) {
				fmt.Printf("%+v\n", s)
			}),
	).Run()
}
