package config

import (
	_ "embed"
	"github.com/goccy/go-yaml"
	"os"
	"sync"
)

type Config struct {
	PostgresSQL struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		Port     int    `yaml:"port"`
		Sslmode  string `yaml:"sslmode"`
	} `yaml:"PostgresSQL"`

	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"Redis"`

	Registry struct {
		RegistryAddress string `yaml:"registryAddress"`
		UserName        string `yaml:"userName"`
		Password        string `yaml:"password"`
	} `yaml:"Registry"`
}

var (
	instance *Config
	once     sync.Once
)

//go:embed config.yaml
var data []byte

func GetConfig() *Config {
	once.Do(func() {
		instance = loadConfig()
	})
	return instance
}

func loadConfig() *Config {
	var conf Config

	err := yaml.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}

	etcdHostEnv := os.Getenv("ETCD_HOST")
	if etcdHostEnv != "" {
		conf.Registry.RegistryAddress = etcdHostEnv
	}

	return &conf
}
