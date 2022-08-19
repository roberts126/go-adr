package adr

import (
	"fmt"

	"github.com/roberts126/go-adr/adr/validate"
)

// Validation represents a function used to validate an input variable
type Validation struct {
	// Operation is the name of the function to call
	Operation string `yaml:"operation"`

	// Args represent the additional args to pass to the function
	Args []string `yaml:"args"`

	// Message is the string returned if the validation fails
	Message string `yaml:"message"`
}

// GetFunction returns an existing validator function or an error if it doesn't exist.
func (v *Validation) GetFunction(s string) (validate.Validator, error) {
	fn := validate.GetValidatorFunction(v.Operation)

	if fn == nil {
		return nil, fmt.Errorf("validator function with name %s does not exist", v.Operation)
	}

	return fn(s, v.Args...)
}
