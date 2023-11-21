package database

type Config struct {
	Host     string
	User     string
	Password string
	DbName   string
	Port     uint
}

func ProvideConfig() *Config {
	return &Config{
		Host:     "0.0.0.0",
		User:     "root",
		Password: "password",
		DbName:   "test",
		Port:     5432,
	}
}
