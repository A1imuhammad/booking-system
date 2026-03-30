package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string     `yaml:"env" env-default:"local"`
	JWT      JWT        `yaml:"jwt"`
	GRPC     GRPCConfig `yaml:"grpc"`
	Postgres Postgres   `yaml:"postgres"`
}

type JWT struct {
	Secret string        `yaml:"secret"`
	TTL    time.Duration `yaml:"ttl" env-default:"24h"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env-default:"44041"`
	Timeout time.Duration `yaml:"timeout" env-default:"10s"`
}

type Postgres struct {
	PostgresUser     string `yaml:"postgres_user" env-default:"root"`
	PostgresPassword string `yaml:"postgres_password" env-default:"1234"`
	PostgresPort     uint16 `yaml:"postgres_port" env-default:"5432"`
	PostgresDB       string `yaml:"postgres_db" env-default:"auth_db"`
	PostgresHost     string `yaml:"postgres_host" env-default:"localhost"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("fail to read config" + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
