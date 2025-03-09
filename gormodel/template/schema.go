package template

import (
	"gormodel/gormodel/sql"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// 从已有的文件中读取 schema，备份原来的 .go 文件
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
