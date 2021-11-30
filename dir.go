package tics

import (
	"io/fs"
	"io/ioutil"
	"os"
	"strings"
)

func GetDir(path string) []fs.FileInfo {
	f, err := ioutil.ReadDir(path)
	if err != nil {
		ThrowSys(GetDir, err)
	}

	return f
}

func GetFile(path string) []byte {
	b, err := os.ReadFile(path)
	if err != nil {
		ThrowSys(GetFile, err)
	}

	return b
}

func GetDirFiles(path string) []string {
	files := []string{}
	dir := GetDir(path)
	for _, file := range dir {
		files = append(files, file.Name())
	}

	return files
}

func GetDirFilesExclusive(path string, notExcludedByFunc func(string) bool) []string {
	files := []string{}
	dir := GetDir(path)
	for _, file := range dir {
		if notExcludedByFunc(file.Name()) {
			files = append(files, file.Name())
		}
	}

	return files
}

func GetDirContents(dirName string) map[string]interface{} {
	m := map[string]interface{}{}
	dir := GetDir(dirName)

	for _, file := range dir {
		p := dirName + "/" + file.Name()
		m[file.Name()] = GetFile(p)
	}

	return m
}

func GetDirContentsExclusive(dirName string, notExcludedByFunc func(string) bool) map[string]interface{} {
	m := map[string]interface{}{}
	dir := GetDir(dirName)

	for _, file := range dir {
		if notExcludedByFunc(file.Name()) {
			p := dirName + "/" + file.Name()
			m[file.Name()] = GetFile(p)
		}
	}

	return m
}

func NotConfig(name string) bool {
	if len(name) > 4 {
		return name[len(name)-4:] != "yaml"
	}

	return true
}

func SearchInDirNames(dirPath, searchStr string) []string {
	dir := GetDirFilesExclusive(dirPath, NotConfig)
	matches := []string{}
	for _, d := range dir {
		if strings.Contains(d, searchStr) {
			matches = append(matches, d)
		}
	}

	return matches
}
