package templates

import (
	"fmt"
	"gormodel/internal/sql"
	"gormodel/internal/utils"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// CheckDefaultSchema 从已有的文件中读取 schema，备份原来的 .go 文件
func CheckDefaultSchema(fileName string) {
	pkg := sql.ReadPackage(fileName)
	dir := filepath.Dir(fileName)
	schemaFile := filepath.Join(dir, "Schema.model.go")
	os.Rename(schemaFile, schemaFile+".bak")
	file, err := os.Create(schemaFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 使用模板渲染Schema.model.go文件
	data := SchemaTemplate{
		Timestamp:   time.Now().Format(time.RFC3339),
		PackageName: pkg,
	}
	
	content, err := RenderSchemaTemplate(data)
	if err != nil {
		panic(err)
	}
	
	file.WriteString(content)

	cmd := exec.Command("goimports", "-w", schemaFile)
	cmd.Run()
}

// GenerateModelFile 生成模型文件
func GenerateModelFile(schema *sql.Schema, modelName string, tableName string) error {
	fileName := modelName
	fmt.Printf("Write %s\n", fileName)
	
	packageName := sql.ReadPackage(fileName)
	
	// 使用模板渲染.model.go文件
	data := ModelTemplate{
		PackageName: packageName,
		Timestamp:   time.Now().Format(time.RFC3339),
		ModelCode:   schema.Write(),
		ModelName:   utils.Camel(tableName),
		TableName:   schema.Schema,
	}
	
	content, err := RenderModelTemplate(data)
	if err != nil {
		return err
	}

	os.Rename(fileName, fileName+".bak")

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	
	file.WriteString(content)
	cmd := exec.Command("goimports", "-w", fileName)
	cmd.Run()
	
	return nil
} 