// config/config.go
package config

import (
	"os"
)

type Config struct {
	AWS struct {
		Region      string
		DynamoTable string
	}
}

func LoadConfigFromEnv() Config {
	var cfg Config

	cfg.AWS.Region = os.Getenv("AWS_REGION")
	cfg.AWS.DynamoTable = os.Getenv("DYNAMO_TABLE_NAME")

	return cfg
}
