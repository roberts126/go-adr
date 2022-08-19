package adr

import (
	"github.com/roberts126/go-adr/adr/validate"
	"github.com/roberts126/go-adr/internal/terminal"
)

// Variable represents a structure used to collect input from a user
type Variable struct {
	// Name is the name of the variable. Also becomes the template key
	Name string `yaml:"name"`

	// Prompt is the string displayed to the user describing the input.
	Prompt string `yaml:"prompt"`

	// Default is the value if an empty string is entered
	Default *string `yaml:"default,omitempty"`

	// Validation is a list of Validation pointers used to validate the input
	Validation []*Validation `yaml:"validation,omitempty"`

	//Optional marks the variable as option or not
	Optional *bool `yaml:"optional,omitempty"`

	// Repeat describes the repeating behavior of the variable.
	Repeat *Repeat `yaml:"repeat,omitempty"`
	value  interface{}
}

// Repeat represents the components used in building a list of values for a variable
type Repeat struct {
	// ExitValue is the value used to indicate the end of the input
	ExitValue string `yaml:"exitValue,omitempty"`

	// MaxItems is the maximum number of items in a list
	MaxItems *int `yaml:"maxItems,omitempty"`

	// MinItems is the minimum number of items in a list
	MinItems *int `yaml:"minItems,omitempty"`

	// Prompt is the prompt, shown after the initial variable prompt.
	// If this value is an empty string then the initial variable prompt
	// is repeated.
	Prompt string `yaml:"prompt,omitempty"`
}

// Input prompts the user for the value(s) for a variable.
func (v *Variable) Input() error {
	if v.Repeat != nil {
		return v.repeatInput(0)
	}

	s, err := v.read(v.Prompt)
	if err != nil {
		return err
	}

	if s == "" {
		v.value = v.Default
	} else {
		v.value = s
	}

	return v.validate(v.value.(string))
}

func (v *Variable) repeatInput(i int) error {
	if v.Repeat.Prompt == "" {
		v.Repeat.Prompt = v.Prompt
	}

	list := make([]string, 0)

	for {
		var prompt string
		if i == 0 {
			prompt = v.Prompt
		} else {
			prompt = v.Repeat.Prompt
		}

		s, err := v.read(prompt)
		if err != nil {
			return err
		}

		if s == v.Repeat.ExitValue {
			v.value = list
			return nil
		}

		if valid := v.validate(s); valid != nil {
			terminal.Warn(err)

			if err = v.repeatInput(i); err != nil {
				return err
			}
		} else {
			list = append(list, s)

			i++
		}
	}
}

func (v *Variable) read(prompt string) (string, error) {
	terminal.Standard(prompt)

	s, err := terminal.Read()
	if err != nil {
		return "", err
	}

	return s, nil
}

func (v *Variable) validate(s string) (err error) {
	validators := make([]validate.Validator, len(v.Validation))

	for i := range v.Validation {
		validators[i], err = v.Validation[i].GetFunction(s)
		if err != nil {
			return err
		}

		validators[i].SetMessage(v.Validation[i].Message)
	}

	return validate.All(validators...)
}
