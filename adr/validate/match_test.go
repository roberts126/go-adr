package validate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatch(t *testing.T) {
	pattern := `^[A-Z][A-Za-z0-9]+$`
	invalidPattern := `[a-z`
	tests := map[string]struct {
		Value       string
		Pattern     *string
		ShouldError bool
		Valid       bool
	}{
		"Test": {
			Value:       "Test",
			Pattern:     &pattern,
			ShouldError: false,
			Valid:       true,
		},
		"InvalidMatch": {
			Value:       "10-Test",
			Pattern:     &pattern,
			ShouldError: false,
			Valid:       false,
		},
		"Error": {
			Value:       "Test",
			Pattern:     nil,
			ShouldError: true,
			Valid:       false,
		},
		"InvalidRegex": {
			Value:       "Test",
			Pattern:     &invalidPattern,
			ShouldError: false,
			Valid:       false,
		},
	}

	for name, test := range tests {
		t.Run("TestMatch"+name, func(t *testing.T) {
			var matcher Validator
			var err error

			if test.Pattern != nil {
				matcher, err = Match(test.Value, *test.Pattern)
			} else {
				matcher, err = Match(test.Value)
			}

			var msg string
			if test.ShouldError {
				assert.Error(t, err, msg)
			} else {
				if assert.NoErrorf(t, err, msg) {
					valid := matcher.Validate()

					if test.Valid {
						assert.NoErrorf(t, valid, "Test %s must validate", name)
					} else {
						assert.Errorf(t, valid, "Test %s must not validate", name)
					}
				}
			}
		})
	}

	t.Run("TestSetMessage", func(t *testing.T) {
		match := MatchValidator{
			value:   "",
			pattern: "",
			msg:     "",
		}

		match.SetMessage("test")

		assert.Equal(t, match.msg, "test", "Message must be updated after SetMessage")
	})
}
