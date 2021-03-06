package helpers

import (
	"os"
)

type ConnectionVar string

const (
	Host     ConnectionVar = "DB_HOST"
	Name     ConnectionVar = "DB_NAME"
	Password ConnectionVar = "DB_PASSWORD"
	Port     ConnectionVar = "DB_PORT"
	Ssl      ConnectionVar = "DB_SSL"
	User     ConnectionVar = "DB_USER"
)

func ConnEnvVars() map[ConnectionVar]string {
	return map[ConnectionVar]string{
		Host:     os.Getenv(string(Host)),
		Name:     os.Getenv(string(Name)),
		Password: os.Getenv(string(Password)),
		Port:     os.Getenv(string(Port)),
		Ssl:      os.Getenv(string(Ssl)),
		User:     os.Getenv(string(User)),
	}
}
