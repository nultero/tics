package tics

import (
	"os"
	"strings"
)

var dirs = []string{
	"confDir",
	"dataDir",
}

// Expects "confDir" and possibly a "dataDir" in the map argument; returns
// these same values with the $USER, $HOME, or '~' strings replaced with the
// actual $HOME value.
//
// Left open to multiple data dirs in the future.
func CobraRootInitBoilerPlate(confMap map[string]string, crashOnErr bool) map[string]string {

	for _, d := range dirs {
		if path, ok := confMap[d]; ok {
			confMap[d] = ReplaceHomeDir(path, crashOnErr)
		}
	}

	return confMap
}

// Replaces the $USER / $HOME / ~ in the string arg and returns
// the str with actual system homeDir. Retains crashOnErr from the parent function, or from arg.
func ReplaceHomeDir(s string, crashOnErr bool) string {

	home, err := os.UserHomeDir()
	if err != nil {
		if crashOnErr {
			ThrowSys(err)
		} else {
			ThrowUnhandled(err)
		}
	}

	replacements := []string{
		"$HOME",
		"$USER",
		"~/",
	}

	for _, v := range replacements {
		s = strings.ReplaceAll(s, v, home)
	}

	return s
}
