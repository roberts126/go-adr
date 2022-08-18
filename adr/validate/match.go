package validate

import (
	"errors"
	"regexp"
)

var _ Validator = (*MatchValidator)(nil)

// MatchValidator provides a mechanism to match a value to a regular expression.
type MatchValidator struct {
	value   string
	pattern string
	msg     string
}

// Match returns a MatchValidator that implements the Validator interface.
//
// If there isn't exactly one arg passed then an error is thrown.
func Match(v string, args ...string) (Validator, error) {
	if len(args) != 1 {
		return nil, errors.New("missing pattern to match")
	}

	validator := &MatchValidator{
		value:   v,
		pattern: args[0],
	}

	return validator, nil
}

// SetMessage sets the error message that is returned when validation fails.
func (v *MatchValidator) SetMessage(msg string) {
	v.msg = msg
}

// Validate validates the value matches the regular express and returns an error if it doesn't
func (v *MatchValidator) Validate() error {
	r, err := regexp.Compile(v.pattern)
	if err != nil {
		return err
	}

	if !r.MatchString(v.value) {
		return errors.New(v.msg)
	}

	return nil
}
