package termeneutics

import (
	"testing"
)

var k Cmd = Cmd{
	Name: "plop",
}

func TestAddDesc(t *testing.T) {
	s := "makes a test call"
	k.AddDesc(s)
	if k.DescStr != s {
		t.Errorf("TestAddDesc fails with desc: %v, want %v", k.DescStr, s)
	}
}

func TestSetNumArgs(t *testing.T) {
	n := 2
	k.SetNumPossArgs(n)
	if k.NumPossArgs != n {
		t.Errorf("TestSetNumArgs fails with num: %v, want %v", k.NumPossArgs, n)
	}
}
