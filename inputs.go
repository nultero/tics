package tics

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/manifoldco/promptui"
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

// Literally just a wrapper around promptui, but reflects on errs,
// and is reusable for my libs.
func SelectBetween(explanation string, args []string) string {
	prompt := promptui.Select{
		Label: explanation,
		Items: args,
	}

	_, result, err := prompt.Run()
	if err != nil {
		ThrowSys(SelectBetween, err)
	}

	return result
}
