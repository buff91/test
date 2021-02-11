package configs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type Config struct {
	AWSConfig AWSConfig `json:"AWSConfig"`
}

type AWSConfig struct {
	AccessKeyId	string `json:"ACCESS_KEY_ID"`
	SecretAccessKey	string `json:"SECRET_ACCESS_KEY"`
}

var config *Config
var once sync.Once
var err error

func newConfig() (*Config, error) {
	file, err := os.Open("configs/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	config := Config{}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	envAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	envSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if envAccessKeyID != "" && envSecretAccessKey != "" {
		config.AWSConfig.AccessKeyId = envAccessKeyID
		config.AWSConfig.SecretAccessKey = envSecretAccessKey
	}

	return &config, nil
}

func GetConfig() (*Config){
	once.Do(func() {
		config, err = newConfig()
	})

	if err != nil {
		log.Println(err)
	}

	return config
}