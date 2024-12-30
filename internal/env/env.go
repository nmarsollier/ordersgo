package env

import (
	"os"
	"strconv"
)

// Configuration properties
type Configuration struct {
	ServerName        string `json:"serverName"`
	Port              int    `json:"port"`
	GqlPort           int    `json:"gqlPort"`
	RabbitURL         string `json:"rabbitUrl"`
	MongoURL          string `json:"mongoUrl"`
	SecurityServerURL string `json:"securityServerUrl"`
	FluentURL         string `json:"fluentUrl"`
}

var config *Configuration

// Get Obtiene las variables de entorno del sistema
func Get() *Configuration {
	if config == nil {
		config = load()
	}

	return config
}

// Load file properties
func load() *Configuration {
	// Default
	result := &Configuration{
		ServerName:        "ordersgo",
		Port:              3004,
		GqlPort:           4004,
		RabbitURL:         "amqp://localhost",
		MongoURL:          "mongodb://localhost:27017",
		SecurityServerURL: "http://localhost:3000",
		FluentURL:         "localhost:24224",
	}

	if value := os.Getenv("SERVER_NAME"); len(value) > 0 {
		result.ServerName = value
	}

	if value := os.Getenv("RABBIT_URL"); len(value) > 0 {
		result.RabbitURL = value
	}

	if value := os.Getenv("MONGO_URL"); len(value) > 0 {
		result.MongoURL = value
	}

	if value := os.Getenv("FLUENT_URL"); len(value) > 0 {
		result.FluentURL = value
	}

	if value := os.Getenv("PORT"); len(value) > 0 {
		if intVal, err := strconv.Atoi(value); err == nil {
			result.Port = intVal
		}
	}

	if value := os.Getenv("GQL_PORT"); len(value) > 0 {
		if intVal, err := strconv.Atoi(value); err == nil {
			result.GqlPort = intVal
		}
	}

	if value := os.Getenv("AUTH_SERVICE_URL"); len(value) > 0 {
		result.SecurityServerURL = value
	}

	return result
}
