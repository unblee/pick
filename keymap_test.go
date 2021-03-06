package pick

import (
	"testing"
)

func TestKeyGroup_in(t *testing.T) {
	tests := []struct {
		keynames []string
		key      rune
		want     bool
	}{
		{
			keynames: []string{"tab", "j"},
			key:      keymap["tab"],
			want:     true,
		},
		{
			keynames: []string{"tab", "j"},
			key:      keymap["q"],
			want:     false,
		},
	}

	for _, test := range tests {
		kg, err := newKeymapGroup(test.keynames...)
		if err != nil {
			t.Fatal(err)
		}
		got := kg.in(test.key)
		if got != test.want {
			t.Errorf("keynames: %v, key: %v, got: %v, want: %v", test.keynames, test.key, got, test.want)
		}
	}
}

func TestNewKeyGroupErrorComposer_error(t *testing.T) {
	tests := []struct {
		keynames []string
		want     bool
	}{
		{
			keynames: []string{"enter", "space"},
			want:     true,
		},
		{
			keynames: []string{"notExisting", "space"},
			want:     false,
		},
	}

	for _, test := range tests {
		kgec := newKeymapGroupErrorComposer{}
		for _, name := range test.keynames {
			_ = kgec.newKeymapGroup(name)
		}
		got := kgec.error() == nil
		if got != test.want {
			t.Errorf("keynames: %v, got: %v, want: %v", test.keynames, got, test.want)
		}
	}
}
