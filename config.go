package main

import (
	"os"

	"github.com/go-yaml/yaml"
)

type ServerConfig struct {
	Port string `yaml:"port"`
}

type StorageAdapterType string

const (
	RedisStorageAdapterType  StorageAdapterType = "redis"
	MemoryStorageAdapterType                    = "memory"
)

type StorageConfig struct {
	Adapter  StorageAdapterType `yaml:"adapter"`
	Host     string             `yaml:"host,omitempty"`
	Password string             `yaml:"password,omitempty"`
	DB       int                `yaml:"db,omitempty"`
}

type ResponsesConfig struct {
	AskForName    string `yaml:"ask_for_name"`
	AskForPurpose string `yaml:"ask_for_purpose"`
	Complete      string `yaml:"complete"`
	Error         string `yaml:"error"`
	UnknownState  string `yaml:"unknown_state"`
}

type TwilioConfig struct {
	SID             string `yaml:"sid"`
	Token           string `yaml:"token"`
	FromPhoneNumber string `yaml:"from_phone_number"`
	ToPhoneNumber   string `yaml:"to_phone_number"`
}

type Config struct {
	Server    ServerConfig    `yaml:"server"`
	Storage   StorageConfig   `yaml:"storage"`
	Responses ResponsesConfig `yaml:"responses"`
	Twilio    TwilioConfig    `yaml:"twilio"`
}

func ReadConfig() (*Config, error) {
	f, err := os.Open("./config.yml")
	if err != nil {
		return nil, err
	}

	dec := yaml.NewDecoder(f)
	var result Config
	err = dec.Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
