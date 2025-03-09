package sql

import (
	"fmt"
	"gormodel/internal/utils"
	"strings"
)

type Column struct {
	name             string
	typ              string
	size             string
	precision        string
	unsigned         bool
	isNull           bool
	defaultValue     string
	autoIncrement    bool
	comment          string
	isPrimaryKey     bool
	indexNames       []string
	uniqueIndexNames []string
	isDeletedAt      bool
}

func NewColumn(columnLineStr string) *Column {
	// fmt.Printf("column line str: %s\n", columnLineStr)
	var column Column
	column.name = matchColumnLineName(columnLineStr)
	column.typ, column.size = matchColumnLineType(columnLineStr)
	column.size, column.precision = matchColumnTypeSize(column.size)
	column.unsigned = matchColumnUnsigned(columnLineStr)
	column.isNull = !matchColumnNotNull(columnLineStr)
	column.autoIncrement = matchColumnAutoIncrement(columnLineStr)
	column.comment = matchColumnComment(columnLineStr)
	column.defaultValue = matchColumnDefaultValue(columnLineStr)
	// fmt.Printf("column: %#v\n", column)
	return &column
}

const columnMapTemplate = `

`

const (
	columnTemplate       = "\nfunc %s_%s() Column[%s] {\n	return \"%s\"\n}\n"
	schemaColumnTemplate = "\nfunc (s *%s) Column%s() Column[%s] {\n	return %s_%s() }\n"
)

func (c *Column) column(schema string) string {
	var columns []string
	columns = append(columns,
		fmt.Sprintf(columnTemplate,
			schema,
			strings.ToLower(string([]byte{c.name[0]}))+c.name[1:],
			c.GoType(),
			utils.Snake(c.name)))

	columns = append(columns,
		fmt.Sprintf(schemaColumnTemplate,
			utils.Camel(schema),
			utils.Camel(c.name),
			c.GoType(),
			schema,
			c.name,
		),
	)
	return strings.Join(columns, "\n")
}

func (c *Column) Write() string {
	columnFormat := "%s %s `gorm:\"%s\" json:\"%s\"`"
	return fmt.Sprintf(columnFormat, utils.Camel(c.name), c.GoType(), c.gorm(), utils.Snake(c.name))
}

func (c *Column) GoType() string {
	if c.isDeletedAt {
		return "gorm.DeletedAt"
	}

	goType := convertType(c.typ, c.size)
	if c.isNull {
		return "*" + goType
	}
	return goType
}

func (c *Column) gorm() string {
	// fmt.Printf("column: %#v\n", c)
	var collects []string
	collects = append(collects, "column:"+utils.Snake(c.name))
	if c.unsigned {
		c.typ += " unsigned"
	}
	collects = append(collects, "type:"+strings.ToUpper(c.typ))

	if c.size != "" {
		collects = append(collects, "size:"+c.size)
	}
	if c.precision != "" {
		collects = append(collects, "precision:"+c.precision)
	}

	if c.isPrimaryKey {
		collects = append(collects, "primaryKey")
	}

	if len(c.indexNames) > 0 {
		collects = append(collects, "index:"+strings.Join(c.indexNames, ","))
	}

	if len(c.uniqueIndexNames) > 0 {
		collects = append(collects, "uniqueIndex:"+strings.Join(c.uniqueIndexNames, ","))
	}
	if !c.isNull {
		collects = append(collects, "not null")
	}

	if c.defaultValue != "" {
		collects = append(collects, "default:"+c.defaultValue)
	}

	if c.autoIncrement {
		collects = append(collects, "autoIncrement")
	}

	return strings.Join(collects, ";")
}
