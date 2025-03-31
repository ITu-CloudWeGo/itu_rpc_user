package config

import (
	_ "embed"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/goccy/go-yaml"
	"os"
	"strconv"
	"sync"
)

type Config struct {
	Kitex struct {
		Service       string `yaml:"Service"`
		Address       string `yaml:"Address"`
		LogLevel      string `yaml:"LogLevel"`
		LogFileName   string `yaml:"LogFileName"`
		LogMaxSize    int    `yaml:"LogMaxSize"`
		LogMaxAge     int    `yaml:"LogMaxAge"`
		LogMaxBackups int    `yaml:"LogMaxBackups"`
	} `yaml:"Kitex"`

	PostgresSQL struct {
		Host     string `yaml:"Host"`
		User     string `yaml:"User"`
		Password string `yaml:"Password"`
		DBName   string `yaml:"Dbname"`
		Port     int    `yaml:"Port"`
		Sslmode  string `yaml:"Sslmode"`
	} `yaml:"PostgresSQL"`

	Redis struct {
		Address  string `yaml:"Address"`
		Username string `yaml:"Username"`
		Password string `yaml:"Password"`
		DB       int    `yaml:"DB"`
	} `yaml:"Redis"`

	Registry struct {
		RegistryAddress []string `yaml:"RegistryAddress"`
		Username        string   `yaml:"Username"`
		Password        string   `yaml:"Password"`
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

func LogLevel() klog.Level {
	level := GetConfig().Kitex.LogLevel
	switch level {
	case "trace":
		return klog.LevelTrace
	case "debug":
		return klog.LevelDebug
	case "info":
		return klog.LevelInfo
	case "notice":
		return klog.LevelNotice
	case "warn":
		return klog.LevelWarn
	case "error":
		return klog.LevelError
	case "fatal":
		return klog.LevelFatal
	default:
		return klog.LevelInfo
	}
}

func loadConfig() *Config {
	var conf Config

	err := yaml.Unmarshal(data, &conf)
	if err != nil {
		panic(err)
	}

	// Redis
	redisAddrEnv := os.Getenv("REDIS_ADDR")
	if redisAddrEnv != "" {
		conf.Redis.Address = redisAddrEnv
	}
	redisUserNameEnv := os.Getenv("REDIS_USERNAME")
	if redisUserNameEnv != "" {
		conf.Redis.Username = redisUserNameEnv
	}
	redisPasswdEnv := os.Getenv("REDIS_PASSWORD")
	if redisPasswdEnv != "" {
		conf.Redis.Password = redisPasswdEnv
	}
	redisDBEnv := os.Getenv("REDIS_DB")
	if redisDBEnv != "" {
		db, err := strconv.Atoi(redisDBEnv)
		if err == nil {
			conf.Redis.DB = db
		}
	}

	// Registry
	registryAddrEnv := os.Getenv("REGISTRY_ADDR")
	if registryAddrEnv != "" {
		conf.Registry.RegistryAddress = []string{registryAddrEnv}
	}
	registryUserNameEnv := os.Getenv("REGISTRY_USERNAME")
	if registryUserNameEnv != "" {
		conf.Registry.Username = ""
	}
	registryPasswordEnv := os.Getenv("REGISTRY_PASSWORD")
	if registryPasswordEnv != "" {
		conf.Registry.Password = registryPasswordEnv
	}

	return &conf
}
