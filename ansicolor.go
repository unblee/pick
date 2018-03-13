package pick

import (
	"fmt"
)

type ansiColor map[string]string

func (a ansiColor) get(name, fallback string) string {
	if c, ok := a[name]; ok {
		return c
	}
	return a[fallback]
}

var (
	fgcolor = ansiColor{
		"black":   "30",
		"gray":    "30",
		"red":     "31",
		"green":   "32",
		"yellow":  "33",
		"blue":    "34",
		"magenta": "35",
		"cyan":    "36",
		"white":   "37",
	}

	bgcolor = ansiColor{
		"black":   "40",
		"gray":    "40",
		"red":     "41",
		"green":   "42",
		"yellow":  "43",
		"blue":    "44",
		"magenta": "45",
		"cyan":    "46",
		"white":   "47",
	}
)

func color(s, fg, fgfallback, bg, bgfallback string) string {
	fc := fgcolor.get(fg, fgfallback)
	bc := bgcolor.get(bg, bgfallback)
	return fmt.Sprintf("\x1b[%s;%sm%s", fc, bc, s)
}
