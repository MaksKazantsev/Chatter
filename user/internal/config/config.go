package config

import (
	"flag"
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port   string   `yaml:"port"`
	Env    string   `yaml:"env"`
	DB     Postgres `yaml:"db"`
	Broker Broker   `yaml:"broker"`
}

type Postgres struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	User     string `yaml:"user"`
	Name     string `yaml:"name"`
}

type Broker struct {
	Addr  string `yaml:"addr"`
	Topic string `yaml:"topic"`
}

func (p *Postgres) GetAddr() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", p.User, p.Password, p.Host, p.Port, p.Name)
}

func MustInit() *Config {
	path := fetchPath()

	_, err := os.Stat(path)
	if err != nil {
		panic("cfg file not found: " + err.Error())
	}
	b, err := os.ReadFile(path)
	if err != nil {
		panic("failed to read cfg file: " + err.Error())
	}

	var cfg Config
	if err = yaml.Unmarshal(b, &cfg); err != nil {
		panic("failed to unmarshal cfg file: " + err.Error())
	}
	return &cfg
}

func fetchPath() string {
	var path string

	flag.StringVar(&path, "c", "", "path to cfg file")
	flag.Parse()

	err := ifEmpty(path)
	if err != nil {
		panic("failed to load env file")
	}

	return path
}

func ifEmpty(field string) error {
	if field == "" {
		err := godotenv.Load("env")
		if err != nil {
			return err
		}
		os.Getenv("CONFIG_PATH")
	}
	return nil
}
