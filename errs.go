package tics

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
)

func blameFunc(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func quit(s interface{}) {
	fmt.Println(s)
	os.Exit(1)
}

func Throw(err string) {
	s := fmt.Sprintf("%s %s", redError(), err)
	quit(s)
}

func ThrowUnhandled(err error) {
	s := fmt.Sprintf("%s %s", redError(), err)
	fmt.Println(s)
}

func ThrowX(err error, desc string) {
	s := fmt.Sprintf("%s %s \n >> %s", redError(), err, desc)
	quit(s)
}

// An unrecoverable error with a descriptive package pointer to what went wrong.
func ThrowSys(funcCulprit interface{}, err error) {
	bf := blameFunc(funcCulprit)
	tf := Make(bf).Blue().String()
	s := fmt.Sprintf("%s %s has caused -> %s", redError(), tf, err)
	quit(s)
}

func ThrowTooManyArgs(cmd string) {
	s := fmt.Sprintf("%s too many args for '%s' cmd", redError(), cmd)
	quit(s)
}

func ThrowTooFewArgs(cmd string) {
	s := fmt.Sprintf("%s too few args for '%s' cmd", redError(), cmd)
	quit(s)
}

func redError() string {
	return "\033[31;1;4merror:\033[0m"
}

func ThrowQuiet(s string) {
	fmt.Println(s)
	os.Exit(0)
}
