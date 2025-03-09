package template

import (
	"bytes"
	"text/template"
)

// ModelTemplate 定义了生成.model.go文件的模板
type ModelTemplate struct {
	PackageName string
	Timestamp   string
	ModelCode   string
	ModelName   string
	TableName   string
}

// SchemaTemplate 定义了生成Schema.model.go文件的模板
type SchemaTemplate struct {
	Timestamp   string
	PackageName string
}

// RenderModelTemplate 渲染.model.go文件的模板
func RenderModelTemplate(data ModelTemplate) (string, error) {
	tmplContent, err := templateFS.ReadFile("templates/model.tmpl")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("model").Parse(string(tmplContent))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// RenderSchemaTemplate 渲染Schema.model.go文件的模板
func RenderSchemaTemplate(data SchemaTemplate) (string, error) {
	tmplContent, err := templateFS.ReadFile("templates/schema.tmpl")
	if err != nil {
		return "", err
	}

	tmpl, err := template.New("schema").Parse(string(tmplContent))
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
} 