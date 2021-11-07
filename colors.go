package tics

import "fmt"

// Returns blued ASCII
func Blue(s string) string {
	return fmt.Sprintf("\x1b[32;1;4m%v\x1b[0m", s)
}

// Returns dark-blued ASCII
func DarkBlue(s string) string {
	return fmt.Sprintf("\x1b[34;1;4m%v\x1b[0m", s)
}

// Returns pink ASCII str
func Pink(s string) string {
	return fmt.Sprintf("\x1b[35;1;1m%v\x1b[0m", s)
}

// Returns red ASCII
func Red(s string) string {
	return fmt.Sprintf("\x1b[31;1;4m%v\x1b[0m", s)
}

// Returns bolded ASCII
func Emph(s string) string {
	return fmt.Sprintf("\x1b[;1;1m%v\x1b[0m", s)
}
