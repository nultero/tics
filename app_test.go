package termeneutics

import (
	"testing"
)

var a = App{
	Cmds:   []Cmd{},
	Consts: map[string]interface{}{},
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

func TestAddConst(t *testing.T) {
	s := "testKey"
	a.AddConst(s, 5)
	if v, ok := a.Consts[s]; !ok {
		t.Errorf("TestAddConst failed with const: %v, want `%v`", v, s)
	}
}

func TestSetAppPath(t *testing.T) {
	path := "~/junkPath"
	a.SetAppPath(path)
	if a.AppPath != path {
		t.Errorf("TestSetAppPath failed with path: %v, want `%v`", a.AppPath, path)
	}
}
