package adr

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("TestNewConfigDefault", func(t *testing.T) {
		tests := map[string]struct {
			expected  *Config
			source    []byte
			willError bool
		}{
			"Default": {
				expected: &Config{
					Projects: make([]*Project, 0),
				},
				source:    nil,
				willError: false,
			},
			"FromBytes": {
				expected: &Config{
					Projects: []*Project{
						{
							Name:      "Example",
							Directory: "/some/path/to/repo/docs/adr",
							Template:  "~/.config/adr/templates/default.tpl",
						},
					},
				},
				source:    []byte(strings.ReplaceAll(ExampleConfiguration, "# ", "")),
				willError: false,
			},
			"BadYaml": {
				expected:  nil,
				source:    []byte("---\nkey: value\n badIndent: value"),
				willError: true,
			},
		}

		for name, test := range tests {
			actual, err := NewConfig(test.source)

			if test.willError {
				assert.Errorf(t, err, "NewConfig%s must error", name)
			} else {
				assert.NoErrorf(t, err, "NewConfig%s must not error", name)
			}
			assert.Equalf(t, test.expected, actual, "NewConfig%s must match", name)
		}
	})

	t.Run("TestLoadProjectFromCwd", func(t *testing.T) {
		cwd, _ := os.Getwd()
		home, _ := os.UserHomeDir()

		defer func() {
			_ = os.Chdir(cwd)
		}()

		cnf := &Config{
			Projects: []*Project{
				{
					Name:      "Valid",
					Directory: cwd,
					Template:  "default.tpl",
				},
			},
		}

		tests := map[string]struct {
			expected *Project
			wd       string
		}{
			"Valid": {
				expected: &Project{
					Name:      "Valid",
					Directory: cwd,
					Template:  "default.tpl",
				},

				wd: cwd,
			},
			"Invalid": {
				expected: nil,
				wd:       home,
			},
		}

		for name, test := range tests {
			_ = os.Chdir(test.wd)

			actual, err := cnf.LoadProjectFromCwd()

			if test.expected == nil {
				assert.Errorf(t, err, "Project %s must error", name)
			} else {
				assert.NoErrorf(t, err, "Project %s must not error", name)
			}

			assert.Equalf(t, test.expected, actual, "Project %s must match", name)
		}
	})
}
