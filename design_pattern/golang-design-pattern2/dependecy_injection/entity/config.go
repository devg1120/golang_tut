package entity

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	DSN string `envconfig:"DSN" required:"true"`
}
