package config

const (
	Host          = "localhost"
	UserName      = "postgres"
	Password      = "postgres"
	Port          = "5432"
	Database      = "wb"
	NatsClusterId = ""
	NatsHostName  = ""
)

type Config struct {
	Host          string
	UserName      string
	Password      string
	Port          string
	Database      string
	NatsClusterId string
	NatsHostName  string
}
