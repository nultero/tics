package tics

import (
	"fmt"
	"os"
)

const perm = 0755

// Runs the interactive prompts that create the default
// config and data dirs where this given CLI lives.
func RunConfPrompts(toolName string, confMap map[string]string, defaults []string) {

	// referencing dirs variable in paths.go
	confDir := dirs[0]
	if path, ok := confMap[confDir]; ok {
		d := coalesceDefaults(defaults)
		makeConfigPrompt(path, toolName, d)
	}

	dataDir := dirs[1]
	if path, ok := confMap[dataDir]; ok {
		makeDirPrompt(path, toolName)
	}

	outro(toolName)
}

// Take a slice of strs and run a set of replacements over them,
// eventually returning a single writeable string.
func coalesceDefaults(sl []string) string {
	fullStr := ""
	for _, s := range sl {
		fullStr += ReplaceHomeDir(s, true) + "\n"
	}

	//shave off last \n
	return fullStr[:len(fullStr)-1]
}

// Gets confirmation to create config, and attempts to write conf file.
// Stops executing and throws error outright on failure to write.
func makeConfigPrompt(path, toolName, defaults string) {
	fmt.Printf("> config file `%v` not found\n", path)
	PrintPrompt("config", toolName)

	inp := GetInput()
	for !IsValidInput(inp) {
		PrintInvalid(inp)
		PrintPrompt("config", toolName)
		inp = GetInput()
	}

	if inp == "y" {
		err := os.WriteFile(path, []byte(defaults), perm)
		if err != nil {
			ThrowSysDescriptor(BlameFunc(makeConfigPrompt), err)
		}

		fmt.Printf("created config for %v: '%v' \n", toolName, path)

	} else {
		fmt.Println("config not created")
	}
}

func makeDirPrompt(dirPath, toolName string) {
	fmt.Printf("\n> directory path `%v` not found\n", dirPath)
	PrintPrompt("dir", toolName)

	inp := GetInput()
	for !IsValidInput(inp) {
		fmt.Printf("> directory path `%v` not found\n", dirPath)
		PrintPrompt("dir", toolName)
		inp = GetInput()
	}

	if inp == "y" {
		err := os.Mkdir(dirPath, perm)
		if err != nil {
			ThrowSysDescriptor(BlameFunc(makeDirPrompt), err)
		}
		fmt.Printf("dir created: `%v`\n", dirPath)

	} else {
		fmt.Printf("dir not created; `%v` lib may have a way to rerun this func\n", toolName)
	}
}

func PrintPrompt(kind, toolName string) {
	fmt.Printf("create this %v with %v defaults? [ y / N ] >", kind, Blue(toolName))
}

func outro(toolName string) {
	f := MakeT("all setup for").Blue().Str()
	n := Pink(toolName)
	l := MakeT("has been completed").Blue().Str()
	fmt.Println(f, n, l)
}
