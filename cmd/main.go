package main

import (
	"fmt"
	"tosiek88/micro/modules/sandbox"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		sandbox.Module(),
		fx.Invoke(
			func(s *sandbox.Sandbox) {
				fmt.Printf("%+v\n", s)
			}),
	).Run()
}
