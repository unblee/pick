package pick

import (
	"errors"
	"fmt"
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

var dupNLReg = regexp.MustCompile(`\n{1,}`)

func splitStdin(stdin io.Reader) (items []string, err error) {
	b, err := ioutil.ReadAll(stdin)
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

func (p *Pick) Exec() {
}


func (p *Pick) drawList(items []string, cursorPos int) {
	for i, item := range items {
		if i == cursorPos {
			fmt.Fprint(p.stdout, color(
				item,
				fgcolor["black"],
				fgcolor["black"],
				bgcolor["white"],
				bgcolor["white"],
			))
		} else {
			fmt.Fprint(p.stdout, item)
		}
		fmt.Fprint(p.stdout, "\n")
	}
}
