package termeneutics

type Arg struct {
	Name     string
	Flags    []rune
	Function func()
}

func (a Arg) HasFlag(r rune) bool {
	for _, f := range a.Flags {
		if f == r {
			return true
		}
	}

	return false
}

func BuildCmds(args []string) []*Arg {

	argSlice := []*Arg{}

	for _, arg := range args {
		a := NewArg(arg)
		argSlice = append(argSlice, &a)
	}

	return argSlice
}

func NewArg(cmdName string, flags ...interface{}) Arg {

	setFlags := []rune{}
	if len(flags) != 0 {
		for _, flag := range flags {
			f, ok := flag.([]rune)
			if ok {
				setFlags = append(setFlags, f...)
			}
		}
	}

	return Arg{
		Name:     cmdName,
		Flags:    setFlags,
		Function: nil,
	}
}

func (a *Arg) AssignFunction(f func()) {
	*&a.Function = f
}
