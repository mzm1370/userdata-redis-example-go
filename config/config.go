package config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func NewConfig() *Config {
	return &Config{
		DBHost:     "localhost",
		DBPort:     "3306",
		DBUser:     "root",
		DBPassword: "password",
		DBName:     "moochin",
		JWTSecret:  "MySecretKeyForJWTToken123",
	}
}
