package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	CONFIG_FILE = "env.yaml"
)

type Config struct {
	ConfigInterface
	Env             *Env
}

type (
	// add method you need
	ConfigInterface interface {
		GetEnvironment() string
	}

	DBConfig struct {
		Driver   string `yaml:"driver"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		DBPath   string `yaml:"db_path"`
	}

	// modify here for your requirements
	Env struct {
		Environment string   `yaml:"environment"`
		Host        string   `yaml:"host"`
		Port        int      `yaml:"port"`
		Salt        string   `yaml:"salt"`
		Database    DBConfig `yaml:"database"`
	}
)

func Load(filename string) (*Config, error) {
	var e Env
	var c Config

	yml, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yml, &e)
	if err != nil {
		return nil, err
	}

	c.Env = &e

	return &c, nil
}

func (c *Config) GetEnvironment() string {
	return c.Env.Environment
}
