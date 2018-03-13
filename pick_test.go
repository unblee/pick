package pick

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSplitItems(t *testing.T) {
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
	}

	for _, test := range tests {
		buf := bytes.NewBufferString(test.input)
		pick := New(buf, nil, nil)
		got, err := pick.splitItems()
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("got: %v, want: %v", got, test.want)
		}
		buf.Reset()
	}
}
