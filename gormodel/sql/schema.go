package sql

import (
	"fmt"
	"gormodel/pkg"
	"os"
	"strings"
)

type Schema struct {
	Path    string
	Raw     string
	Schema  string
	Columns []*Column
}

func (s *Schema) Write() string {
	schemaFormat := "type %s struct {\n%s\n}\n"
	var columnLines []string
	for _, column := range s.Columns {
		columnLines = append(columnLines, column.Write())
	}

	var parts []string
	parts = append(parts,
		fmt.Sprintf(
			schemaFormat,
			pkg.Camel(s.Schema),
			strings.Join(columnLines, "\n"),
		),
	)
	for _, column := range s.Columns {
		parts = append(parts, column.column(s.Schema))
	}
	return strings.Join(parts, "\n")
}

func ReadSqlFile(path string) *Schema {
	var schema Schema
	schema.Path = path
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	schema.Raw = string(data)
	schema.Schema = MatchSchemaName(schema.Raw)
	fmt.Printf("- read columns ...\n")
	columnStrs := MatchAllColumns(schema.Raw)
	columns := make(map[string]*Column)
	for _, columnStr := range columnStrs {
		column := NewColumn(columnStr)
		columns[column.name] = column
		schema.Columns = append(schema.Columns, column)
	}

	fmt.Printf("- read index ...\n")

	matchKey(schema.Raw, columns)

	// fmt.Printf("schema: %+v\n", schema)
	return &schema
}
