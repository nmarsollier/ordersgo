package env

import (
	"os"
	"strconv"
)

// Configuration properties
type Configuration struct {
	Port              int    `json:"port"`
	GqlPort           int    `json:"gqlPort"`
	RabbitURL         string `json:"rabbitUrl"`
	SecurityServerURL string `json:"securityServerUrl"`
	FluentUrl         string `json:"fluentUrl"`
	AwsAccessKeyId    string `json:"AwsAccessKeyId"`
	AwsSecret         string `json:"AwsSecret"`
	AwsRegion         string `json:"AwsRegion"`
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
		Port:              3004,
		GqlPort:           4004,
		RabbitURL:         "amqp://localhost",
		AwsAccessKeyId:    "placeholder_value",
		AwsSecret:         "placeholder_value",
		AwsRegion:         "placeholder_value",
		SecurityServerURL: "http://localhost:3000",
		FluentUrl:         "localhost:24224",
	}

	if value := os.Getenv("RABBIT_URL"); len(value) > 0 {
		result.RabbitURL = value
	}

	if value := os.Getenv("AWS_ACCESS_KEY_ID"); len(value) > 0 {
		result.AwsAccessKeyId = value
	}

	if value := os.Getenv("AWS_SECRET_ACCESS_KEY"); len(value) > 0 {
		result.AwsSecret = value
	}

	if value := os.Getenv("AWS_REGION"); len(value) > 0 {
		result.AwsRegion = value
	}

	if value := os.Getenv("FLUENT_URL"); len(value) > 0 {
		result.FluentUrl = value
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
