package adr

import (
	"github.com/roberts126/go-adr/adr/validate"
	"github.com/roberts126/go-adr/internal/terminal"
)

type Variables struct {
	Name         string        `yaml:"name"`
	Prompt       string        `yaml:"prompt"`
	Default      *string       `yaml:"default,omitempty"`
	Validation   []*Validation `yaml:"validation,omitempty"`
	Optional     *bool         `yaml:"optional,omitempty"`
	Repeats      *bool         `yaml:"repeats,omitempty"`
	RepeatPrompt string        `yaml:"repeatPrompt,omitempty"`
	MaxItems     *int          `yaml:"maxItems"`
	MinItems     *int          `yaml:"minItems"`
	ExitValue    string        `yaml:"exitValue,omitempty"`
	value        interface{}
}

// Input prompts the user for the value(s) for a variable.
func (v *Variables) Input() error {
	if v.Repeats != nil && *v.Repeats {
		return v.repeatInput()
	}

	terminal.Standard(v.Prompt)
	s, err := terminal.Read()
	if err != nil {
		return err
	}

	if s == v.ExitValue {
		v.value = v.Default
	} else {
		v.value = s
	}

	return v.validate(v.value.(string))
}

func (v *Variables) repeatInput() error {
	if v.RepeatPrompt == "" {
		v.RepeatPrompt = v.Prompt
	}

	list := make([]string, 0)

	i := 0
	for {
		if i == 0 {
			terminal.Standard(v.Prompt)
		} else {
			terminal.Standard(v.RepeatPrompt)
		}

		s, err := terminal.Read()
		if err != nil {
			return err
		}

		if s == v.ExitValue {
			v.value = list
			return nil
		}

		if err = v.validate(s); err != nil {
			return err
		}

		list = append(list, s)

		i++
	}
}

func (v *Variables) validate(s string) (err error) {
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
