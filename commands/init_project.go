package commands

import (
	"errors"
	"fmt"
	"os"
	"strings"

	u "pecker/utils"
)

const (
	// ProjectName 模板项目名
	projectName = "peckergo"
)

// InitProject 初始化项目
func InitProject() error {
	newProjectName := u.CheckProject()
	if newProjectName == "" {
		return errors.New("只能在 peckergo 内使用该命令")
	}
	currentPath, _ := os.Getwd()
	u.Walk(currentPath, func(path string, info os.FileInfo, err error) error {
		content := u.ReadFile(path)

		if strings.Contains(content, projectName) {
			content = strings.Replace(content, projectName, newProjectName, -1)

			u.WriteToFile(path, content)
			fmt.Printf("\t%s%supdate file%s\t %s%s\n", "\x1b[36m", "\x1b[1m", "\x1b[21m", path, "\x1b[0m")
		}
		return nil
	})
	return nil
}
