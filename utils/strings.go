package utils

import (
	"os"
	"strings"
	"unicode"
)

// SnakePath 路径蛇形化 /Users/bill/gopath/src/web-clt/cmd/commands => cmd_commands
func SnakePath(path string) string {
	currentpath, _ := os.Getwd()
	path = strings.Replace(path, currentpath+"/", "", -1)
	path = strings.Replace(path, "-", "_", -1)
	path = strings.Replace(path, "/", "_", -1)
	path = strings.Replace(path, ".", "_", -1)
	return strings.ToLower(path)
}

// BarPath 路径横杠形式 /Users/bill/gopath/src/web-clt/cmd/commands => cmd-commands
func BarPath(path string) string {
	currentpath, _ := os.Getwd()
	path = strings.Replace(path, currentpath+"/", "", -1)
	path = strings.Replace(path, "_", "-", -1)
	path = strings.Replace(path, "/", "-", -1)
	path = strings.Replace(path, ".", "-", -1)
	return strings.ToLower(path)
}

// CamelCasePath 路径驼峰化 /Users/bill/gopath/src/web-clt/cmd/commands => CmdCommands
func CamelCasePath(path string) string {
	snake := SnakePath(path)
	splits := strings.Split(snake, "_")
	for idx, s := range splits {
		splits[idx] = Ucfirst(s)
	}
	return strings.Join(splits, "")
}

// LCamelCasePath 路径驼峰化 /Users/bill/gopath/src/web-clt/cmd/commands => cmdCommands
func LCamelCasePath(path string) string {
	snake := SnakePath(path)
	splits := strings.Split(snake, "_")
	for idx, s := range splits {
		if idx == 0 {
			splits[idx] = Lcfirst(s)
		} else {
			splits[idx] = Ucfirst(s)
		}
	}
	return strings.Join(splits, "")
}

// Ucfirst 首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// Lcfirst 首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
