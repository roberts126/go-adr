package adr

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

const errTemplateLoad = "unable to load template from %s; error: %v"

// Project represents the settings for a specific project
type Project struct {
	// Name is the name of the project
	Name string `yaml:"name"`

	// Directory is the location of the ADRs in the project
	Directory string `yaml:"directory"`

	// Template is the path to the template file
	Template string `yaml:"template"`
}

// LoadTemplate loads the template from the specified path. If the file
// does not exist or cannot be read an error is returned.
func (p *Project) LoadTemplate() (*Template, error) {
	_, err := os.Stat(p.Template)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf(errTemplateLoad, p.Template, "file does not exist")
	}

	data, err := ioutil.ReadFile(p.Template)
	if err != nil {
		return nil, fmt.Errorf(errTemplateLoad, p.Template, err)
	}

	var t Template
	if err = yaml.Unmarshal(data, &t); err != nil {
		return nil, err
	}

	return &t, nil
}
