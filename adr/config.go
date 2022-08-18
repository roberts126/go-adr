package adr

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Projects []*Project `yaml:"projects"`
}

func NewConfig(data []byte) (*Config, error) {
	var c Config
	if data == nil {
		c = Config{
			Projects: make([]*Project, 0),
		}
	} else {
		if err := yaml.Unmarshal(data, &c); err != nil {
			return nil, err
		}

	}

	return &c, nil
}

func (c *Config) LoadProjectFromCwd() (*Project, error) {
	cwd, _ := os.Getwd()

	for _, project := range c.Projects {
		if cwd == project.Directory {
			return project, nil
		}
	}

	return nil, fmt.Errorf("unable to find an ADR project for the current directory")
}
