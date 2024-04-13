package object

type Config struct {
	DBHost string `env:"DB_HOST"`
	DBPort string `env:"DB_PORT"`
	DBName string `env:"DB_NAME"`
	DBUser string `env:"DB_USER"`
	DBPass string `env:"DB_PASSWORD"`
}
