package termeneutics

import "testing"

var a = App{
	Cmds: []Cmd{},
}

func TestAddCmdCall(t *testing.T) {
	a.AddCmd(c)
	nameBuf := ""
	for _, com := range a.Cmds {
		nameBuf = com.Name
	}

	if nameBuf != c.Name {
		t.Errorf("TestAddCmd failed with cmd: %v, want %v", nameBuf, c)
	}
}
