package terminal

import (
	"bufio"
	"os"
	"regexp"
	"syscall"

	"golang.org/x/term"
)

var pwReader Prompt

type terminalReader struct {
	buf int
}

func (r *terminalReader) ReadPassword() ([]byte, error) {
	return term.ReadPassword(r.buf)
}

func (r *terminalReader) Read() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

func SetReader(r Prompt) {
	pwReader = r
}

type Prompt interface {
	ReadPassword() ([]byte, error)
	Read() (string, error)
}

func YesNo() (bool, error) {
	yn, err := Read()
	if err != nil {
		return false, err
	}

	r, _ := regexp.Compile(`(?mi)y(?:es)?`)
	return r.MatchString(yn), nil
}

func GetPassword() ([]byte, error) {
	if pwReader == nil {
		pwReader = &terminalReader{
			buf: int(syscall.Stdin),
		}
	}

	return pwReader.ReadPassword()
}

func Read() (string, error) {
	if pwReader == nil {
		pwReader = &terminalReader{
			buf: int(syscall.Stdin),
		}
	}

	return pwReader.Read()
}
