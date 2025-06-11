package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	PostgresUser             string
	PostgresPassword         string
	PostgresDB               string
	PostgresHost             string
	PostgresPort             int
	KafkaHost                string
	KafkaPort                string
	KafkaConsumerGroupID     string
	PrometheusPushGatewayURL string
	InstanceID               string
}

func LoadConfig() (*Config, error) {
	postgresUser, ok := os.LookupEnv("POSTGRES_USER")
	if !ok {
		return nil, fmt.Errorf("POSTGRES_USER environment variable is not set")
	}

	postgresPassword, ok := os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		return nil, fmt.Errorf("POSTGRES_PASSWORD environment variable is not set")
	}

	postgresHost, ok := os.LookupEnv("POSTGRES_HOST")
	if !ok {
		return nil, fmt.Errorf("POSTGRES_HOST environment variable is not set")
	}

	postgresPortStr, ok := os.LookupEnv("POSTGRES_PORT")
	if !ok {
		return nil, fmt.Errorf("POSTGRES_PORT environment variable is not set")
	}

	postgresPort, err := strconv.Atoi(postgresPortStr)
	if err != nil {
		return nil, fmt.Errorf("POSTGRES_PORT must be an integer: %w", err)
	}

	postgresDB, ok := os.LookupEnv("POSTGRES_DB")
	if !ok {
		return nil, fmt.Errorf("POSTGRES_DB environment variable is not set")
	}

	kafkaHost, ok := os.LookupEnv("KAFKA_HOST")
	if !ok {
		return nil, fmt.Errorf("KAFKA_HOST environment variable is not set")
	}

	kafkaPort, ok := os.LookupEnv("KAFKA_PORT")
	if !ok {
		return nil, fmt.Errorf("KAFKA_PORT environment variable is not set")
	}

	kafkaConsumerGroupID, ok := os.LookupEnv("KAFKA_CONSUMER_GROUP_ID")
	if !ok {
		return nil, fmt.Errorf("KAFKA_CONSUMER_GROUP_ID environment variable is not set")
	}

	prometheusPushGatewayURL, ok := os.LookupEnv("PROMETHEUS_PUSH_GATEWAY_URL")
	if !ok {
		return nil, fmt.Errorf("PROMETHEUS_PUSH_GATEWAY_URL environment variable is not set")
	}

	instanceID, ok := os.LookupEnv("INSTANCE_ID")
	if !ok {
		return nil, fmt.Errorf("INSTANCE_ID environment variable is not set")
	}

	return &Config{
		PostgresUser:             postgresUser,
		PostgresPassword:         postgresPassword,
		PostgresHost:             postgresHost,
		PostgresPort:             postgresPort,
		PostgresDB:               postgresDB,
		KafkaHost:                kafkaHost,
		KafkaPort:                kafkaPort,
		KafkaConsumerGroupID:     kafkaConsumerGroupID,
		PrometheusPushGatewayURL: prometheusPushGatewayURL,
		InstanceID:               instanceID,
	}, nil
}

func (c *Config) GetPostgresUser() string {
	return c.PostgresUser
}

func (c *Config) GetPostgresPassword() string {
	return c.PostgresPassword
}

func (c *Config) GetPostgresHost() string {
	return c.PostgresHost
}

func (c *Config) GetPostgresPort() int {
	return c.PostgresPort
}

func (c *Config) GetPostgresDB() string {
	return c.PostgresDB
}

func (c *Config) GetKafkaBrokers() []string {
	return []string{c.KafkaHost + ":" + c.KafkaPort}
}

func (c *Config) GetKafkaConsumerGroupID() string {
	return c.KafkaConsumerGroupID
}

func (c *Config) GetPrometheusPushGatewayURL() string {
	return c.PrometheusPushGatewayURL
}
