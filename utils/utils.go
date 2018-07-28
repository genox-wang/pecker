package utils

import (
	"os"
	"path"
	"path/filepath"
)

var (
	projectName = "app"
	templateKV  = map[string]string{}
	// MustHaveDirs 项目必须拥有的目录
	MustHaveDirs = []string{
		"api/model", "api/controller", "api/router",
	}
)

// CheckProject 验证项目目录是否正确，返回项目名
func CheckProject() string {
	currentPath, _ := os.Getwd()
	for _, dir := range MustHaveDirs {
		if !IsExist(path.Join(currentPath, dir)) {
			return ""
		}
	}
	return filepath.Base(currentPath)
}

// // MustProject 必须是模板项目
// func MustProject() string {
// 	projectName := CheckProject()
// 	if projectName == "" {
// 		panic(errors.New("is not a project of peckergo"))
// 	}
// 	return projectName
// }

// SetProjectName 设置项目名
func SetProjectName(pName string) {
	projectName = pName
}

// SetTemplateKV 配置全局模板映射
func SetTemplateKV(kv map[string]string) {
	templateKV = kv
}

// TemplateKV 获取templateKV
func TemplateKV() map[string]string {
	return templateKV
}

// CreateWithTemplate 根据模板创建
func CreateWithTemplate(relativeFilename, template string) {
	CreateWithTmpl(relativeFilename, template, templateKV)
}

// AppendWithTemplate 根据模板创建或添加
func AppendWithTemplate(relativeFilename string, placehold, appendTmpl string) {
	AppendWithTmpl(relativeFilename, templateKV, placehold, appendTmpl, templateKV)
}
