package termeneutics

import (
	"testing"
)

var c Cmd = Cmd{
	Name: "plop",
}

func TestAddFlags(t *testing.T) {
	testFlags := []rune{'a', 'b', 'c'}
	c.AddFlags(testFlags)
	for i := range c.Flags {
		if c.Flags[i] != testFlags[i] {
			t.Errorf("TestAddFlags failed to load flag %v at index %v", testFlags[i], i)
		}
	}
}

func TestHasFlagFunc(t *testing.T) {
	testFlag := 'b'
	if !c.HasFlag(testFlag) {
		t.Errorf("TestHasFlagFunc failed to load, want flag '%v'", string(testFlag))
	}
}

func TestNumFlags(t *testing.T) {
	i := len(c.Flags)
	if i != 3 {
		t.Errorf("TestNumFlags func has '%v' number of flags, want 3", i)
	}
}
