package env

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DBHost     string `envconfig:"DB_HOST" required:"true"`
	DBPort     int    `envconfig:"DB_PORT" required:"true"`
	DBDatabase string `envconfig:"DB_DATABASE" required:"true"`
	DBUser     string `envconfig:"DB_USER" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
}

var Conf Config

func init() {
	if err := envconfig.Process("APP", &Conf); err != nil {
		panic(err)
	}
}
