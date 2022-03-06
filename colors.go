package tics

import (
	"fmt"
	"strings"
)

type str struct {
	content string
}

// Value 	Color
// \e[0;30m 	Black
// \e[0;31m 	Red
// \e[0;32m 	Green
// \e[0;33m 	Yellow
// \e[0;34m 	Blue
// \e[0;35m 	Purple
// \e[0;36m 	Cyan
// \e[0;37m 	White

// Creates an intermediate 'tic' for string chain processing.
// Call var.String() to unwrap out of the container.
func Make(s string) *str {
	return &str{content: s}
}

func (s *str) String() string {
	return s.content
}

// 33 is yellow / gold on default terminal?
var fmtMap = map[string]string{
	"darkblue": "\x1b[34;1;4m%v\x1b[0m",
	"blue":     "\x1b[34;1m%v\x1b[0m",
	"bold":     "\x1b[;1;1m%v\x1b[0m",
	"green":    "\x1b[32;1m%v\x1b[0m",
	"magenta":  "\x1b[35;1m%v\x1b[0m",
	"pink":     "\x1b[35;1;1m%v\x1b[0m",
	"red":      "\x1b[31;1;1m%v\x1b[0m",
	"uline":    "\x1b[;1;4m%v\x1b[0m",
}

func lookup(lkFmt string) string {
	s, ok := fmtMap[lkFmt]
	if !ok {
		return ""
	}

	return s
}

// Transforms *s into a blue ASCII.
func (s *str) Blue() *str {
	s.content = fmt.Sprintf(lookup("blue"), s.content)
	return s
}

// Bolds the string at *s.
func (s *str) Bold() *str {
	s.content = fmt.Sprintf(lookup("bold"), s.content)
	return s
}

// Transforms *s into a dark blue ASCII.
func (s *str) DarkBlue() *str {
	s.content = fmt.Sprintf(lookup("darkblue"), s.content)
	return s
}

// Transforms *s into a green ASCII.
func (s *str) Green() *str {
	s.content = fmt.Sprintf(lookup("green"), s.content)
	return s
}

// Transforms *s into a magenta ASCII.
func (s *str) Magenta() *str {
	s.content = fmt.Sprintf(lookup("magenta"), s.content)
	return s
}

// Transforms *s into a pink ASCII.
func (s *str) Pink() *str {
	s.content = fmt.Sprintf(lookup("pink"), s.content)
	return s
}

// Transforms *s into a red ASCII.
func (s *str) Red() *str {
	s.content = fmt.Sprintf(lookup("red"), s.content)
	return s
}

// Repeats *s's value a set number of times. Note
// that this should be called after all other *s
// transforms or else values can print incorrectly.
func (s *str) Repeat(count int) *str {
	s.content = strings.Repeat(s.content, count)
	return s
}

// Underlines the string at *s.
func (s *str) Underline() *str {
	s.content = fmt.Sprintf(lookup("uline"), s.content)
	return s
}
