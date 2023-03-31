package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Host           string `json:"host"`
	Path           string `json:"path"`
	UserName       string `json:"username"`
	Password       string `json:"password"`
	ConnsPerHost   int    `json:"connsPerHost"`
	TimeoutSeconds int    `json:"timeoutSeconds"`

	RabbitConnAddr string `json:"rabbitConnAddr"`
}

func getDefault() *Config {
	return &Config{
		ConnsPerHost:   10,
		TimeoutSeconds: 10,
	}
}

func Init(name string) (*Config, error) {
	c := getDefault()

	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("open file: %w", err)
	}

	dec := json.NewDecoder(f)

	err = dec.Decode(c)
	if err != nil {
		return nil, fmt.Errorf("json decode: %w", err)
	}

	return c, nil
}
