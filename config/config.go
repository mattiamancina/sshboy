package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Servers []Server `yaml:"servers"`
}

type Server struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	User string `yaml:"user"`
}

var cfg *Config

func Init(path string) {
	data, err := os.ReadFile(path)
	if err != nil {

		fmt.Println(err)

	}

	c := &Config{} // allocate a real struct
	if err := yaml.Unmarshal(data, c); err != nil {

		return
	}

	cfg = c
}

func Get() *Config {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg
}

func GetServer(name string) *Server {

	for _, s := range cfg.Servers {
		if s.Name == name {
			return &s
		}
	}

	return nil

}
