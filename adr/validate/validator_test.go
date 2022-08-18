package validate

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type noop struct {
	msg   string
	value string
	args  []string
}

func (n *noop) SetMessage(msg string) {
	n.msg = msg
}

func (n *noop) Validate() error {
	if n.value == "-1" {
		return errors.New(n.msg)
	}

	return nil
}

func noopFn(v string, args ...string) (Validator, error) {
	return &noop{
		msg:   "",
		value: v,
		args:  args,
	}, nil
}

func TestValidator(t *testing.T) {
	oldFunctions := allFunctions
	defer func() {
		allFunctions = oldFunctions
	}()

	allFunctions = map[string]Func{
		"noop": noopFn,
	}

	t.Run("TestGetFunction", func(t *testing.T) {
		assert.NotNil(t, GetFunction("noop"), "Existing function must not error")
		assert.Nil(t, GetFunction("fail"), "Non-existent function must error")
	})

	tests := map[string]struct {
		value string
		valid bool
	}{
		"Valid": {
			value: "Test",
			valid: true,
		},
		"Invalid": {
			value: "-1",
			valid: false,
		},
	}

	for name, test := range tests {
		t.Run("TestAll"+name, func(t *testing.T) {
			v, err := noopFn(test.value)

			if assert.NoError(t, err, "Getting validator must not error") {
				if test.valid {
					assert.NoError(t, All(v), "Valid use case must not error")
				} else {
					assert.Error(t, All(v), "Invalid use case must not error")
				}
			}
		})
	}
}
