package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// IsExist returns whether a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// MkdirAll create all directory
func MkdirAll(path string) {
	if IsExist(path) {
		return
	}
	os.MkdirAll(path, 0755)
	fmt.Printf("\t%s%screate dir%s\t %s%s\n", "\x1b[32m", "\x1b[1m", "\x1b[21m", path+string(filepath.Separator), "\x1b[0m")
}

// WriteToFile creates a file and writes content to it
func WriteToFile(filename, content string) {
	f, err := os.Create(filename)
	MustCheck(err)
	defer CloseFile(f)
	_, err = f.WriteString(content)
	MustCheck(err)
}

// ReadFile 读取文件
func ReadFile(filename string) string {
	b, err := ioutil.ReadFile(filename)
	MustCheck(err)
	return string(b)
}

// CloseFile attempts to close the passed file
// or panics with the actual error
func CloseFile(f *os.File) {
	err := f.Close()
	MustCheck(err)
}

// MustCheck panics when the error is not nil
func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// Walk 遍历目录 并添加过滤目录
func Walk(walkPath string, do func(path string, info os.FileInfo, err error) error) {
	filepath.Walk(walkPath, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		if info.IsDir() {
			if info.Name() == ".git" || info.Name() == "node_modules" || info.Name() == "vendor" || info.Name() == "cmd" {
				return filepath.SkipDir
			}

			return nil
		}
		return do(path, info, err)
	})
}

// Relative2AbsolutePath 相对路劲转绝对路径
func Relative2AbsolutePath(relative string) string {
	currpath, _ := os.Getwd()
	return filepath.Join(currpath, relative)
}
