package terminal

import (
	"bytes"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
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

	removeNewline := regexp.MustCompile(`[\r\n]+$`)

	t.Run("TestNonFormats", func(t *testing.T) {
		tests := map[string]func(args ...interface{}){
			"Error":         Error,
			"ErrorZebra":    ErrorZebra,
			"Info":          Info,
			"InfoAlt":       InfoAlt,
			"InfoZebra":     InfoZebra,
			"Panic":         Panic,
			"Standard":      Standard,
			"StandardZebra": StandardZebra,
			"Success":       Success,
			"SuccessZebra":  SuccessZebra,
			"Warn":          Warn,
			"WarnZebra":     WarnZebra,
		}

		for name, fn := range tests {

			if strings.Contains(name, "Zebra") {
				fn(msg, msg)
				actual := removeNewline.ReplaceAllString(buf.String(), "")
				assert.Equalf(t, msg+"\n"+msg, actual, "%s output must match", name)
			} else {
				fn(msg)
				actual := removeNewline.ReplaceAllString(buf.String(), "")
				assert.Equalf(t, msg, actual, "%s output must match", name)
			}

			buf.Reset()
		}
	})

	t.Run("TestFormats", func(t *testing.T) {
		tests := map[string]func(format string, args ...interface{}){
			"Errorf":    Errorf,
			"Infof":     Infof,
			"InfoAltf":  InfoAltf,
			"Panicf":    Panicf,
			"Standardf": Standardf,
			"Successf":  Successf,
			"Warnf":     Warnf,
		}

		for name, fn := range tests {
			fn("%s", msg)
			actual := removeNewline.ReplaceAllString(buf.String(), "")
			assert.Equalf(t, msg, actual, "%s output must match", name)

			buf.Reset()
		}
	})
}
