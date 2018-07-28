package utils

import (
	"fmt"
	"path"
	"strings"
)

// CreateWithTmpl 通过模板创建文件
// relativeFilename 目标文件名(相对路径)
// template 新建模板
// templateKV 新建模板替换键值对 例如 {"packageName": "app"} 会把 template 里的 {{packageName}} 都替换成 app
func CreateWithTmpl(relativeFilename, template string, templateKV map[string]string) {
	Create(Relative2AbsolutePath(relativeFilename), Template2Content(template, templateKV))
}

// AppendWithTmpl 使用模板更新
// relativeFilename 目标文件名相对路径
// template 新建模板
// templateKV 新建模板替换键值对 例如 {"packageName": "app"} 会把 template 里的 {{packageName}} 都替换成 app
// placehold 占位符 如果文件已存在，会使用 appendTmpl 替换当前文件内的占位符
// appendTmpl 替换占位符的模板
// appendTmplKV 对应 appendTmpl 模板的键值对
func AppendWithTmpl(relativeFilename string, templateKV map[string]string, placehold, appendTmpl string, appendTmplKV map[string]string) {
	Append(Relative2AbsolutePath(relativeFilename), placehold, Template2Content(appendTmpl, appendTmplKV))
}

// Create 创建文件
func Create(filename, content string) {
	dir := path.Dir(filename)
	MkdirAll(dir)
	WriteToFile(filename, content)
	fmt.Printf("\t%s%screate file%s\t %s%s\n", "\x1b[36m", "\x1b[1m", "\x1b[21m", filename, "\x1b[0m")
}

// Append 用content替换占位符
func Append(filename, placehold, appendContent string) {
	if IsExist(filename) {
		content := ReadFile(filename)
		if strings.Contains(content, appendContent) {
			return
		}
		content = strings.Replace(content, placehold, appendContent, -1)
		fmt.Printf("\t%s%supdate file%s\t %s%s\n", "\x1b[34m", "\x1b[1m", "\x1b[21m", filename, "\x1b[0m")
		WriteToFile(filename, content)
	}
}

// Template2Content 模板转文本
func Template2Content(template string, templateKV map[string]string) string {
	for k, v := range templateKV {
		old := fmt.Sprintf("{{%s}}", k)
		template = strings.Replace(template, old, v, -1)
	}
	return template
}
