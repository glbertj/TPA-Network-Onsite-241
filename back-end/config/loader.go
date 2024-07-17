package config

import (
	"back-end/utils"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		utils.CheckError(err)
	}

	return &Config{
		Server: Server{
			Port: os.Getenv("SERVER_PORT"),
			Host: os.Getenv("SERVER_HOST"),
		},
		Database: Database{
			Dialect:  os.Getenv("DATABASE_DIALECT"),
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     os.Getenv("DATABASE_PORT"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASS"),
			Name:     os.Getenv("DATABASE_NAME"),
		},
		Email: Email{
			Host:     os.Getenv("MAIL_HOST"),
			Port:     os.Getenv("MAIL_PORT"),
			User:     os.Getenv("MAIL_USER"),
			Password: os.Getenv("MAIL_PASS"),
		},
		Redis: Redis{
			Address:  os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASS"),
		},
		Google: Google{
			ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
			ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
			RedirectURL:  os.Getenv("OAUTH_REDIRECT_URL"),
			AuthURL:      os.Getenv("OAUTH_AUTH_URL"),
		},
	}

}
