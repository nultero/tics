package termeneutics

import (
	"fmt"
	"termeneutics/colors"
	"testing"
)

var a Arg = NewArg("plopCmd", []rune{'a', 'b', 'c'})

func TestNewArgsFlagsInterface(t *testing.T) {
	testFlags := []rune{'a', 'b', 'c'}
	for i := range a.Flags {
		if a.Flags[i] != testFlags[i] {
			s := fmt.Sprintf("TestNewArgsFlagsInterface failed to load flag %v at index %v", testFlags[i], i)
			t.Errorf(colors.Emph(s))
		}
	}
}

func TestHasFlagFunc(t *testing.T) {
	testFlag := 'b'

	if !a.HasFlag(testFlag) {
		s := fmt.Sprintf("TestHasFlagFunc failed to load, want flag '%v'", string(testFlag))
		t.Errorf(colors.Emph(s))
	}
}

func TestNumFlags(t *testing.T) {
	var a Arg = NewArg("plopCmd", []rune{'a', 'b', 'c'})
	i := len(a.Flags)

	if i != 3 {
		s := fmt.Sprintf("TestNumFlags func has '%v' number of flags, want 3", i)
		msg := colors.Emph(s)
		t.Errorf(msg)
	}
}
