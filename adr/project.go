package adr

type Project struct {
	Name      string `yaml:"name"`
	Directory string `yaml:"directory"`
	Template  string `yaml:"template"`
}
