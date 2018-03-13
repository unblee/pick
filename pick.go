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
