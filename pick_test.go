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
		{
			input: "\n\n\nA\nB\n\n\nC\n\n",
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
			t.Errorf("input: %v, got: %v, want: %v", test.input, got, test.want)
		}
		buf.Reset()
	}
}
