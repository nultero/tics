package tics

import (
	"fmt"
)

type tic struct {
	content string
}

// CANONICAL COLORS
// YOU MAY NOT BE COLOR BLIND HAHA
// Value 	Color
// \e[0;30m 	Black
// \e[0;31m 	Red
// \e[0;32m 	Green
// \e[0;33m 	Yellow
// \e[0;34m 	Blue
// \e[0;35m 	Purple
// \e[0;36m 	Cyan
// \e[0;37m 	White

// Creates an intermediate str 'tic' for chain processing.
// Call tic.Str() to unwrap a str out of the container.
func MakeT(s string) *tic {
	return &tic{content: s}
}

// Transforms *t to a blue ASCII.
func (t *tic) Blue() *tic {
	t.content = fmt.Sprintf("\x1b[32;1;4m%v\x1b[0m", t.content)
	return t
}

// Returns bolded ASCII
func (t *tic) Bold() *tic {
	t.content = fmt.Sprintf("\x1b[;1;1m%v\x1b[0m", t.content)
	return t
}

func (t *tic) Str() string {
	return t.content
}
