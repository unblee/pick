package pick

import (
	"errors"
	"io"
	"io/ioutil"
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

func (p *Pick) splitItems() (items []string, err error) {
	b, err := ioutil.ReadAll(p.stdin)
	if err != nil {
		return
	}

	if len(b) == 0 {
		err = errors.New("input is empty")
		return
	}

	s := strings.TrimRight(string(b), "\n") // trim trailing "\n"
	items = strings.Split(s, "\n")

	return
}
