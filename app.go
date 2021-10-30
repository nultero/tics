package termeneutics

type App struct {
	Cmds []Cmd
}

func (a *App) AddCmd(c Cmd) {
	a.Cmds = append(a.Cmds, c)
}
