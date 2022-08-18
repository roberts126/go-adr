package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnv(t *testing.T) {
	t.Run("TestGetExistingString", func(t *testing.T) {
		key := "ADR_UNIT_TEST"
		defaultStr := "default"
		expected := "test"

		_ = os.Unsetenv(key)
		assert.Equal(t, defaultStr, GetString(key, defaultStr), "Non-existent key must return the default value")

		if assert.NoError(t, os.Setenv(key, expected), "Setting the test variable must not error") {
			assert.Equal(t, expected, GetString(key, defaultStr), "Existing key value must match")
		}
	})
}
