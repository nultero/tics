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
