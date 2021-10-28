package termeneutics

type Arg struct {
	Name     string
	Flags    []rune
	Function func()
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

	// not finished, need a poll / some func for
	// different types of flags

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
	a.Function = f
}
