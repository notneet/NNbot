package common

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

var DEFAULT_CONFIG_DIR = "./config/general.yaml"

func LoadConfig(filename string) (Config, error) {
	// read yaml
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read YAML file: %v", err)
	}

	// parse yaml
	var config Config
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read YAML file: %v", err)
	}

	return config, nil
}
