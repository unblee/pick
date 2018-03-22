package pick

import (
	"errors"
)

var keymap = map[string]rune{
	"tab":     9,
	"enter":   13,
	"ctrl_n":  14,
	"ctrl_p":  16,
	"space":   32,
	"down":    65,
	"up":      66,
	"shift_g": 71,
	"g":       103,
	"j":       106,
	"k":       107,
	"q":       113,
}

type keymapGroup []rune

func newKeymapGroup(keynames ...string) (keymapGroup, error) {
	group := keymapGroup{}
	for _, keyname := range keynames {
		key, existing := keymap[keyname]
		if !existing {
			return keymapGroup{}, errors.New("undefined key")
		}
		group = append(group, key)
	}
	return group, nil
}

func (kg keymapGroup) in(key rune) bool {
	for _, k := range kg {
		if key == k {
			return true
		}
	}
	return false
}
