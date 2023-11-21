package main

import (
	"fmt"

	"github.com/tochamateusz/micro/modules/database"
	"github.com/tochamateusz/micro/modules/sandbox"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ProvideDatabase() (*gorm.DB, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		"0.0.0.0",
		"root",
		"password",
		"test",
		5432)

	return database.ProvideDb(postgres.Open(dns), &database.Config{})
}

func main() {
	fx.New(
		sandbox.Module(),
		database.Module(fx.Decorate(ProvideDatabase)),
		// amqp.Module(),
		fx.Invoke(
			func(s *sandbox.Sandbox) {
				fmt.Printf("%+v\n", s)
			}),
	).Run()
}
