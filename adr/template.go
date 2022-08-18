package adr

type Template struct {
	Variables []*Variables `yaml:"variables"`
	Contents  string       `yaml:"contents"`
	data      map[string]interface{}
}
