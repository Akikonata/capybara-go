package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Qianfan struct {
		ApiKey    string `yaml:"apiKey"`
		SecretKey string `yaml:"secretKey"`
		TokenUrl  string `yaml:"tokenUrl"`
	} `yaml:"qianfan"`
}

var GlobalConfig Config

func init() {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("无法读取YAML文件: %v", err)
	}

	err = yaml.Unmarshal(data, &GlobalConfig)
	if err != nil {
		log.Fatalf("无法解析YAML文件: %v", err)
	}
}
