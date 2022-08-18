package terminal

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlock(t *testing.T) {
	const msg = "message"

	oldExitFunc := exitFunc
	oldWriter := writer

	defer func() {
		exitFunc = oldExitFunc
		writer = oldWriter
	}()

	exitFunc = func(_ int) {
		return
	}

	var buf bytes.Buffer
	writer = &buf

	block := NewBlock("Test")
	indent := "  "
	block.SetIndent(indent)

	t.Run("TestNonFormats", func(t *testing.T) {
		tests := map[string]func(args ...interface{}) *Block{
			"Error":    block.Error,
			"Info":     block.Info,
			"Standard": block.Standard,
			"Success":  block.Success,
			"Warn":     block.Warn,
		}

		for name, fn := range tests {
			fn(msg)
			block.Render()

			expected := fmt.Sprintf("Test\n  %s\n", msg)
			actual := buf.String()
			assert.Equalf(t, expected, actual, "%s output must match", name)

			buf.Reset()
			block.Reset()
		}
	})

	t.Run("TestFormats", func(t *testing.T) {
		tests := map[string]func(format string, args ...interface{}) *Block{
			"Errorf":    block.Errorf,
			"Infof":     block.Infof,
			"Standardf": block.Standardf,
			"Successf":  block.Successf,
			"Warnf":     block.Warnf,
		}

		for name, fn := range tests {
			fn("%s", msg)
			block.Render()

			expected := fmt.Sprintf("Test\n  %s\n", msg)
			actual := buf.String()
			assert.Equalf(t, expected, actual, "%s output must match", name)

			buf.Reset()
			block.Reset()
		}
	})
}
