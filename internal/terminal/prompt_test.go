package terminal

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockPrompt struct {
	str string
	err error
}

func (m *mockPrompt) ReadPassword() ([]byte, error) {
	defer m.reset()
	return []byte(m.str), m.err
}

func (m *mockPrompt) Read() (string, error) {
	defer m.reset()
	return m.str, m.err
}

func (m *mockPrompt) reset() {
	m.str = ""
	m.err = nil
}

func TestPrompt(t *testing.T) {
	oldReader := pwReader
	defer func() {
		pwReader = oldReader
	}()

	mockReader := &mockPrompt{}

	t.Run("TestSetReader", func(t *testing.T) {
		SetReader(mockReader)

		assert.Equal(t, mockReader, pwReader, "SetReader must match")
	})

	t.Run("TestYesNo", func(t *testing.T) {
		tests := map[string]bool{
			"y":   true,
			"Y":   true,
			"yes": true,
			"Yes": true,
			"yEs": true,
			"yeS": true,
			"YeS": true,
			"YEs": true,
			"yES": true,
			"YES": true,
			"n":   false,
			"N":   false,
			"no":  false,
			"No":  false,
			"nO":  false,
			"NO":  false,
		}

		for entry, expected := range tests {
			mockReader.str = entry
			actual, err := YesNo()

			if assert.NoError(t, err, "YesNo prompt must not error"); err != nil {
				assert.Equal(t, expected, actual, "YesNo prompt must match")
			}
		}

		mockReader.str = "Yes"
		mockReader.err = errors.New("mock")

		actual, err := YesNo()
		if assert.Error(t, err, "YesNo must return an error when unable to read a prompt") {
			assert.False(t, actual, "YesNo must be false when an error is encountered")
		}
	})

	t.Run("TestReadPassword", func(t *testing.T) {
		mockReader.str = "password"
		mockReader.err = nil
		expected := []byte(mockReader.str)

		actual, err := GetPassword()
		if assert.NoError(t, err, "ReadPassword must not error") {
			assert.Equal(t, expected, actual, "ReadPassword results must match")
		}
	})
}
