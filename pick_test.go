package pick

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSplitStdin(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{
			input: "A\nB\nC",
			want:  []string{"A", "B", "C"},
		},
		{
			input: "A\nB\nC\n",
			want:  []string{"A", "B", "C"},
		},
		{
			input: "\n\n\nA\nB\n\n\nC\n\n",
			want:  []string{"A", "B", "C"},
		},
	}

	for _, test := range tests {
		buf := bytes.NewBufferString(test.input)
		got, err := splitStdin(buf)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("input: %v, got: %v, want: %v", test.input, got, test.want)
		}
		buf.Reset()
	}
}

func TestPick_drawList(t *testing.T) {
	tests := []struct {
		items     []string
		cursorPos int
		want      string
	}{
		{
			items:     []string{"A", "B", "C"},
			cursorPos: 0,
			want:      color("A", fgcolor["black"], fgcolor["black"], bgcolor["white"], bgcolor["white"]) + "\nB\nC\n",
		},
	}

	for _, test := range tests {
		stdout := new(bytes.Buffer)
		pick := New(nil, stdout, nil)
		pick.drawList(test.items, test.cursorPos)
		got := stdout.String()
		if test.want != got {
			t.Errorf("items: %v, cursorPos: %v, got: %v, want: %v", test.items, test.cursorPos, got, test.want)
		}
	}
}
