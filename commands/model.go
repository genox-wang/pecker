package commands

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	u "pecker/utils"
)

// AddModel 添加模型
func AddModel(name string) error {
	newProjectName := u.CheckProject()
	if newProjectName == "" {
		return errors.New("只能在 peckergo 内使用该命令")
	}
	currentPath, _ := os.Getwd()
	templatesModelRoot := path.Join(currentPath, "templates", "model")
	templatesModelAppendRoot := path.Join(currentPath, "templates", "model_append")

	u.SetProjectName(newProjectName)

	u.SetTemplateKV(map[string]string{
		"projectName": newProjectName,
		"modelName":   u.LCamelCasePath(name),
		"model_name":  u.SnakePath(name),
		"model-name":  u.BarPath(name),
		"ModelName":   u.CamelCasePath(name),
	})
	u.Walk(templatesModelRoot, func(pathName string, info os.FileInfo, err error) error {
		template := u.ReadFile(pathName)
		modelPath := path.Join(strings.Split(u.Template2Content(info.Name(), u.TemplateKV()), "^")...)
		modelPath = modelPath[:len(modelPath)-4]

		u.CreateWithTemplate(modelPath, template)

		return nil
	})

	u.Walk(templatesModelAppendRoot, func(pathName string, info os.FileInfo, err error) error {
		appTemplate := u.ReadFile(pathName)

		reg := regexp.MustCompile("//.*don't remove this line")

		placeholder := reg.FindString(appTemplate)

		fmt.Printf("\t%s%splaceholder%s\t %s%s\n", "\x1b[36m", "\x1b[1m", "\x1b[21m", placeholder, "\x1b[0m")

		appendPath := path.Join(strings.Split(u.Template2Content(info.Name(), u.TemplateKV()), "^")...)
		appendPath = appendPath[:len(appendPath)-4]

		u.AppendWithTemplate(appendPath, placeholder, appTemplate)

		return nil
	})

	return nil
}
