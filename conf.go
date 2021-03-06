package tics

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const Perm = 0755

// Runs the interactive prompts that create the default
// config and data dirs where this given CLI lives.
func RunConfPrompts(toolName string, confMap map[string]string, defaults []string) {

	// referencing dirs str key variable in paths.go
	confFile := dirs[1]
	if path, ok := confMap[confFile]; ok {
		d := coalesceDefaults(defaults)
		if !FileExists(path) {
			makeConfigPrompt(path, toolName, d)
		}
	}

	dataDir := dirs[2]
	if path, ok := confMap[dataDir]; ok {
		if !FileExists(path) {
			makeDirPrompt(path, toolName)
		}
	}

	outro(toolName)
}

// Checks if file exists by tagging os.Stat; came about as tics
// needs a way to check which of the consuming
// CLIs' dependent files tripped the config error.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// Take a slice of strs and run a set of replacements over them,
// eventually returning a single writeable string.
func coalesceDefaults(sl []string) string {
	fullStr := ""
	for _, s := range sl {
		fullStr += ReplaceHomeDir(s, true) + "\n"
	}

	if len(fullStr) == 0 {
		return "" // edge case
	}

	//shave off last \n
	return fullStr[:len(fullStr)-1]
}

// Gets confirmation to create config, and attempts to write conf file.
// Stops executing and throws error outright on failure to write.
func makeConfigPrompt(path, toolName, defaults string) {
	p := Make(path).Bold().String()
	fmt.Printf("> config file `%v` not found\n", p)
	PrintPrompt("config", toolName)

	inp := GetInput()
	for !IsValidInput(inp) {
		PrintInvalid(inp)
		PrintPrompt("config", toolName)
		inp = GetInput()
	}

	if inp == "y" {
		err := os.WriteFile(path, []byte(defaults), Perm)
		if err != nil {

			// tried to create file, parent dir did not exist
			if errors.Is(err, fs.ErrNotExist) {
				dirPath := filepath.Dir(path)
				makeDirPrompt(dirPath, toolName)

				// retry file write
				err = os.WriteFile(path, []byte(defaults), Perm)

				if err != nil { // just dump if another error
					ThrowSys(makeConfigPrompt, err)
				}

			} else {
				ThrowSys(makeConfigPrompt, err)
			}
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
		err := os.Mkdir(dirPath, Perm)
		if err != nil {
			ThrowSys(makeDirPrompt, err)
		}
		fmt.Printf("dir created: `%v`\n", dirPath)

	} else {
		fmt.Printf("dir not created; `%v` lib may have a way to rerun this func\n", toolName)
	}
}

func PrintPrompt(kind, toolName string) {
	tn := Make(toolName).Blue().String()
	fmt.Printf("create this %v with %v defaults? [ y / N ] >", kind, tn)
}

func outro(toolName string) {
	begin := Make("all setup for").Blue().String()
	name := Make(toolName).Pink().String()
	last := Make("has been completed").Blue().String()

	fmt.Println(begin, name, last)
}
