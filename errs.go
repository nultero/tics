package tics

import (
	"fmt"
	"os"
)

func quit(s string) {
	fmt.Println(s)
	os.Exit(1)
}

func Throw(err string) {
	s := fmt.Sprintf("%s %s", redError(), err)
	quit(s)
}

func ThrowX(err error, desc string) {
	s := fmt.Sprintf("%s %s \n >> %s", redError(), err, desc)
	quit(s)
}

func ThrowSys(err error) {
	s := fmt.Sprintf("%s %s", redError(), err)
	quit(s)
}

func ThrowTooManyArgs(cmd string) {
	s := fmt.Sprintf("%s too many args for '%s' cmd", redError(), cmd)
	quit(s)
}

func redError() string {
	return "\033[31;1;4merror:\033[0m"
}

func ThrowQuiet(s string) {
	fmt.Println(s)
	os.Exit(0)
}
