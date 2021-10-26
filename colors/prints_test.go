package Colors

import (
	"fmt"
	"testing"
)

const ColoredString = "I probably messed up my terminal colors"

var Fns = map[string]func(s string) string{
	"Blue:":     Blue,
	"Red:":      Red,
	"Pink:":     Pink,
	"DarkBlue:": DarkBlue,
}

func TestColorFuncs(t *testing.T) {
	for k, f := range Fns {
		s := f(k + " " + ColoredString)
		fmt.Println(s)
	}

	t.Errorf(Emph("TestColorsFunc intentionally fails so as to dump colored output"))
}
