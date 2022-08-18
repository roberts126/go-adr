package adr

import (
	"fmt"

	"github.com/roberts126/go-adr/adr/validate"
)

type Validation struct {
	Operation string   `yaml:"operation"`
	Args      []string `yaml:"args"`
	Message   string   `yaml:"message"`
}

func (v *Validation) GetFunction(s string) (validate.Validator, error) {
	fn := validate.GetFunction(v.Operation)

	if fn == nil {
		return nil, fmt.Errorf("validator function with name %s does not exist", v.Operation)
	}

	return fn(s, v.Args...)
}
