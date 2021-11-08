package tics

import (
	"fmt"
	"os"
)

const perm = 0755

// Gets confirmation to create config, and attempts to write conf file.
// Stops executing and throws error outright on failure to write.
func MakeConfigPrompt(path, lib, defaults string) {
	fmt.Printf("> config file `%v` not found\n", path)
	PrintPrompt("config", lib)

	inp := GetInput()
	for !IsValidInput(inp) {
		PrintInvalid(inp)
		PrintPrompt("config", lib)
		inp = GetInput()
	}

	if inp == "y" {
		err := os.WriteFile(path, []byte(defaults), perm)
		if err != nil {
			ThrowSys(err)
		}

		fmt.Printf("created config for %v: '%v' \n", lib, path)

	} else {
		fmt.Println("config not created")
	}
}

func MakeDirPrompt(dirPath, lib string) {
	fmt.Printf("\n> directory path `%v` not found\n", dirPath)
	PrintPrompt("dir", lib)

	inp := GetInput()
	for !IsValidInput(inp) {
		fmt.Printf("> directory path `%v` not found\n", dirPath)
		PrintPrompt("dir", lib)
		inp = GetInput()
	}

	if inp == "y" {
		err := os.Mkdir(dirPath, perm)
		if err != nil {
			ThrowSys(err)
		}
		fmt.Printf("dir created: `%v`\n", dirPath)

	} else {
		fmt.Printf("dir not created; `%v` lib may have a way to rerun this func\n", lib)
	}
}

func PrintPrompt(kind, lib string) {
	fmt.Printf("create this %v with %v defaults? [ y / N ] >", kind, Blue(lib))
}

// func Fix(path string) {

// 	_, err := os.Stat(path)

// 	path = nv.GetHomeDir() + nv.ExpandTilde(path)

// 	if err != nil {
// 		if errors.Is(err, fs.ErrNotExist) {
// 			novemDirNotExists(path)
// 			initNovemFile(path)
// 		}
// 	}
// }

// const prompt = "create new novem dir? "

// func novemDirNotExists(path string) {
// 	s := "'" + Blue(path) + "'"
// 	fmt.Println("novem directory", s, "does not exist / was deleted \n ")

// 	pr := Emph("[ y / n ]")

// 	fmt.Print(prompt, pr, " ")
// 	answer := strings.ToLower(nv.GetInput())

// 	for !isValid(answer) {
// 		fmt.Println("\n", FingerRed(), Red("expecting valid input"))
// 		fmt.Print(prompt, pr, " ")
// 		answer = strings.ToLower(nv.GetInput())
// 	}

// 	if strings.Contains(answer, "y") {
// 		createConf(path)

// 	} else {
// 		fmt.Println(NovemNine(), "> sure thing, pardner")
// 		os.Exit(0)
// 	}

// }

// func createConf(dir string) {
// 	err := os.Mkdir(dir, perm)

// 	if err != nil {
// 		errs.SysErr(err)
// 	}

// 	dir = "'" + DarkBlue(dir) + "'"
// 	fmt.Println("created", dir)
// }

// func initNovemFile(path string) {

// 	ft, err := os.Stat(path)
// 	if err != nil {
// 		errs.SysErr(err)
// 	}

// 	path = path + "/novem.txt"

// 	now := nv.GetTimeStr(ft.ModTime())
// 	b := []byte(path + nv.SEPARATOR + now)

// 	r := os.WriteFile(path, b, perm)
// 	if r != nil {
// 		errs.SysErr(r)
// 	}

// 	fmt.Println("wrote out", DarkBlue(path), "as datafile")
// }
