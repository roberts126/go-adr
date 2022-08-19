package adr

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config stores all the project configurations for the tool.
type Config struct {
	// Projects is a list of project the tool is aware of.
	Projects map[string]*Project `yaml:"projects"`
}

// NewConfig creates a new config. If the supplied byte slice is empty
// a default configuration is returned. If the supplied byte slice
// cannot be processed an error is returned.
func NewConfig(data []byte) (*Config, error) {
	var c Config
	if data == nil {
		c = Config{
			Projects: make(map[string]*Project, 0),
		}
	} else {
		if err := yaml.Unmarshal(data, &c); err != nil {
			return nil, err
		}

	}

	return &c, nil
}

// LoadProjectFromCwd loads a project from using the current working
// directory to identify the project. If a project doesn't exist
// in the current directory then an error is returned.
func (c *Config) LoadProjectFromCwd() (*Project, error) {
	cwd, _ := os.Getwd()

	project, ok := c.Projects[cwd]
	if !ok {
		return nil, fmt.Errorf("unable to find an ADR project for the current directory")
	}

	return project, nil
}
