package config

import "fmt"

type Es struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (es Es) Url() string {
	return fmt.Sprintf("%s:%d", es.Host, es.Port)
}
