package termeneutics

import (
	"fmt"
	"os"
)

type App struct {
	Cmds    []Cmd
	Consts  map[string]interface{}
	AppPath string
}

func NewApp() App {
	return App{
		Cmds:   []Cmd{},
		Consts: map[string]interface{}{},
	}
}

func (a *App) AddCmd(c Cmd) {
	a.Cmds = append(a.Cmds, c)
}

func (a *App) AddConst(key string, val interface{}) {
	a.Consts[key] = val
}

func (a *App) ParseArgs() ([]string, error) {
	//
	// boiler, there's a few things to run through and check
	// all cmds /docstrings filled out?
	// paths all valid?
	// metaflags for the term lib to do certain things?
	// metacmds? etc
	// errors?
	//
	//
	// Parse should probably be the entry point
	// for the above overall
	//
	//  what should the returns be, if I can't return a generic func?

	var parsedArgs []string

	args := os.Args[1:]
	for _, r := range args {
		fmt.Println(r)
	}

	return parsedArgs, nil
}

func (a *App) SetAppPath(path string) {
	a.AppPath = path
}
