package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type Config struct {
    Temporal struct {
        HostPort    string `yaml:"hostPort"`
        Namespace   string `yaml:"namespace"`
        Identity    string `yaml:"identity"`
    } `yaml:"temporal"`
    RetryPolicy struct {
        InitialInterval    int `yaml:"initialInterval"`
        BackoffCoefficient float64 `yaml:"backoffCoefficient"`
        MaximumInterval    int `yaml:"maximumInterval"`
        MaximumAttempts    int `yaml:"maximumAttempts"`
    } `yaml:"retryPolicy"`
    Logging struct {
        Level string `yaml:"level"`
    } `yaml:"logging"`
    Monitoring struct {
        Enabled bool `yaml:"enabled"`
        Port    int  `yaml:"port"`
    } `yaml:"monitoring"`
}

// LoadConfig loads configuration from a YAML file
func LoadConfig(configFile string) (*Config, error) {
    data, err := ioutil.ReadFile(configFile)
    if err != nil {
        log.Fatalf("Failed to read config file: %v", err)
        return nil, err
    }
    var config Config
    err = yaml.Unmarshal(data, &config)
    if err != nil {
        log.Fatalf("Failed to unmarshal config data: %v", err)
        return nil, err
    }
    return &config, nil
}
