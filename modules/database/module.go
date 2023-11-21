package database

import (
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ProvideDatabase() (*gorm.DB, error) {
	dns := "host=0.0.0.0 user=root password=password dbname=test port=5432 sslmode=disable"
	return gorm.Open(postgres.Open(dns), &gorm.Config{})
}

func Module() fx.Option {
	return fx.Provide(ProvideDatabase)
}
