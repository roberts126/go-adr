package adr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	tests := map[string]struct {
		validation Validation
		valid      bool
	}{
		"Existing": {
			valid: true,
			validation: Validation{
				Operation: "match",
				Args:      []string{"^.*$"},
				Message:   "invalid",
			},
		},
		"Non-existing": {
			valid: false,
			validation: Validation{
				Operation: "asdf",
				Args:      []string{"^.*$"},
				Message:   "invalid",
			},
		},
	}

	for name, v := range tests {
		t.Run("TestGetFunction"+name, func(t *testing.T) {
			fn, err := v.validation.GetFunction("test")

			if v.valid {
				assert.NoError(t, err, "Existing validation must not error.")
				assert.NotNil(t, fn, "Validation function must not be nil")
			} else {
				assert.Error(t, err, "Non-existing validation must error.")
				assert.Nil(t, fn, "Invalid validation function must be nil")
			}
		})
	}
}
