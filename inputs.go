package tics

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput() string {
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	return strings.ReplaceAll(s, "\n", "")
}

func IsValidInput(a string) bool {
	a = strings.ToLower(a)
	return strings.Contains(a, "y") || strings.Contains(a, "n")
}

func PrintInvalid(inp string) {
	fmt.Printf("\ninvalid input '%v' \n", inp)
}
