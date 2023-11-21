package database

import (
	"fmt"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ProvideDatabase(config *Config) (*gorm.DB, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.Host,
		config.User,
		config.Password,
		config.DbName,
		config.Port)

	return ProvideDb(postgres.Open(dns), config)
}

func ProvideDb(dialer gorm.Dialector, config *Config) (*gorm.DB, error) {
	return gorm.Open(dialer, &gorm.Config{})
}

func Module() fx.Option {
	return fx.Provide(ProvideConfig, ProvideDatabase)
}
