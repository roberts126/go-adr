package adr

// Template represents the data necessary to render an ADR
type Template struct {
	// Variables is a list of Variable pointers used to populate the template
	Variables []*Variable `yaml:"variables"`

	// Contents is the actual contents of the ADR
	Contents string `yaml:"contents"`

	data map[string]interface{}
}
