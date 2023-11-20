package main

import (
	"fmt"
	"tochamateusz/micro/modules/message-brokes/amqp"
	"tochamateusz/micro/modules/sandbox"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		sandbox.Module(),
		amqp.Module(),
		fx.Invoke(
			func(s *sandbox.Sandbox) {
				fmt.Printf("%+v\n", s)
			}),
	).Run()
}
