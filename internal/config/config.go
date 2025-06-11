package config

import (
	"fmt"
	"os"
)

type Config struct {
	KafkaHost                string
	KafkaPort                string
	KafkaConsumerGroupID     string
	PrometheusPushGatewayURL string
	InstanceID               string
}

func LoadConfig() (*Config, error) {
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
		KafkaHost:                kafkaHost,
		KafkaPort:                kafkaPort,
		KafkaConsumerGroupID:     kafkaConsumerGroupID,
		PrometheusPushGatewayURL: prometheusPushGatewayURL,
		InstanceID:               instanceID,
	}, nil
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
