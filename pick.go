package pick

import (
	"errors"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

type Pick struct {
	stdin          io.Reader
	stdout, stderr io.Writer
}

func New(stdin io.Reader, stdout, stderr io.Writer) *Pick {
	return &Pick{
		stdin:  stdin,
		stdout: stdout,
		stderr: stderr,
	}
}

func (p *Pick) Exec() {
}

var dupNLReg = regexp.MustCompile(`\n{1,}`)

func (p *Pick) splitStdin() (items []string, err error) {
	b, err := ioutil.ReadAll(p.stdin)
	if err != nil {
		return
	}

	if len(b) == 0 {
		err = errors.New("input is empty")
		return
	}

	replaced := dupNLReg.ReplaceAllString(string(b), "\n")
	trimmed := strings.Trim(replaced, "\n") // trim unnecessary "\n"
	items = strings.Split(trimmed, "\n")

	return
}
