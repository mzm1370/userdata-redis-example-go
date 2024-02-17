package config

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

type RediConfig struct {
	Address  string
	Password string
	DB       int
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

func LoadRedisConfig() *RediConfig {

	return &RediConfig{
		Address:  "localhost:679",
		Password: "",
		DB:       0,
	}
}
