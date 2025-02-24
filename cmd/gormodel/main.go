package main

import (
	"fmt"
	"gormodel/gormodel/sql"
	"gormodel/gormodel/template"
	"gormodel/pkg"
	"os"
	"os/exec"
	"strings"
	"time"
)

// gorm AutoMigrate tips:
// 如果要更新一个 column，光修改它的 type 属性还不够
// 还要判断它在修改之后的 DEFAULT value 是否发生了变化
// 如果修改前的 column 有default value，且此 default value
// 值在修改后没有发生变化，那么此 column 也不会触发 alter column 更新
func main() {
	root := pkg.GetRootPath()
	fmt.Printf("root: %v\n", root)
	var sqlFiles = sql.ListAllSqlFiles(root)
	var schemas []*sql.Schema
	for _, file := range sqlFiles {
		fmt.Printf("Read %s\n", file)
		schema := sql.ReadSqlFile(file)
		schemas = append(schemas, schema)
	}

	for _, schema := range schemas {
		fileName := strings.ReplaceAll(schema.Path, ".sql", ".model.go")
		fmt.Printf("Write %s\n", fileName)
		template.CheckDefaultSchema(fileName)

		packageName := sql.ReadPackage(fileName)
		content := fmt.Sprintf(
			strings.Join(template.SchemaFormat, "\n"),
			packageName,
			time.Now().Format(time.RFC3339),
			schema.Write(),
			pkg.Camel(schema.Schema),
			schema.Schema,
		)

		os.Rename(fileName, fileName+".bak")

		file, err := os.Create(fileName)
		if err != nil {
			panic(err)
		}

		defer file.Close()
		file.WriteString(content)
		cmd := exec.Command("goimports", "-w", fileName)
		cmd.Run()
	}
}
