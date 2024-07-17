package config

type Config struct {
	Server   Server
	Database Database
	Email    Email
	Redis    Redis
	Google   Google
}

type Server struct {
	Port string
	Host string
}

type Database struct {
	Dialect  string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Redis struct {
	Address  string
	Password string
}

type Email struct {
	Host     string
	Port     string
	User     string
	Password string
}

type Google struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	AuthURL      string
}
