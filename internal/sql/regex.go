package sql

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/CrazyThursdayV50/gotils/pkg/collector"
)

const (
	columnLineStr = "\n\\s+`\\w+`.*"
)

func MatchAllColumns(raw string) []string {
	columnLineReg := regexp.MustCompile(columnLineStr)
	lines := columnLineReg.FindAllString(raw, -1)
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, ",")
	}
	return lines
}

const columnLineNameStr = "\n\\s+`(\\w+)`.*"

var columnLineNameReg = regexp.MustCompile(columnLineNameStr)

func matchColumnLineName(columnLineStr string) string {
	columnLineNames := columnLineNameReg.FindStringSubmatch(columnLineStr)
	if len(columnLineNames) < 2 {
		panic("invalid column line name")
	}
	return columnLineNames[1]
}

const (
	columnLineTypeStr      = "(?i)\n\\s+`\\w+`(?:\\s+.*)?\\s+(TINYINT|SMALLINT|MEDIUMINT|BIGINT|FLOAT|DOUBLE|BOOLEAN|DECIMAL|DATETIME|TIMESTAMP|YEAR|VARCHAR|TINYBLOB|TINYTEXT|MEDUMBLOB|MEDIUMTEXT|LONGBLOB|LONGTEXT|ENUM)+(\\([\\d,\\s]+\\))?\\s*"
	columnLineShortTypeStr = "(?i)\n\\s+`\\w+`(?:\\s+.*)?\\s+(INT|DOUBLE|BOOLEAN|DECIMAL|DATE|TIME|YEAR|CHAR|BLOB|TEXT|ENUM)+(\\([\\d,\\s]+\\))?\\s*"
)

var (
	columnLineTypeReg      = regexp.MustCompile(columnLineTypeStr)
	columnLineShortTypeReg = regexp.MustCompile(columnLineShortTypeStr)
)

func matchColumnLineType(columnLineStr string) (string, string) {
	columnsLineTypes := columnLineTypeReg.FindStringSubmatch(columnLineStr)
	// fmt.Printf("column: %s, column types: %#v\n", columnLineStr, columnsLineTypes)
	if len(columnsLineTypes) < 3 {
		columnsLineTypes = columnLineShortTypeReg.FindStringSubmatch(columnLineStr)
		if len(columnsLineTypes) < 3 {
			panic("invalid column types")
		}
	}

	return columnsLineTypes[1] + columnsLineTypes[2], columnsLineTypes[2]
}

const columnTypeSizeStr = "\\((\\d+)(?:,\\s?(\\d+))?\\)"

var columnTypeSizeReg = regexp.MustCompile(columnTypeSizeStr)

func matchColumnTypeSize(columnTypeSize string) (string, string) {
	size := columnTypeSizeReg.FindStringSubmatch(columnTypeSize)
	if len(size) > 2 {
		return size[1], size[2]
	}

	if len(size) > 1 {
		return size[1], ""
	}

	return "", ""
}

const columnUnsignedStr = `(?i)unsigned`

var columnUnsignedReg = regexp.MustCompile(columnUnsignedStr)

func matchColumnUnsigned(columnLineStr string) bool {
	return columnUnsignedReg.MatchString(columnLineStr)
}

const columnNotNullStr = `(?i)not null`

var columnNotNullReg = regexp.MustCompile(columnNotNullStr)

func matchColumnNotNull(columnLineStr string) bool {
	return columnNotNullReg.MatchString(columnLineStr)
}

const columnAutoIncrementStr = "(?i)auto(:?[\\s_]?)?increment"

var columnAutoIncrementReg = regexp.MustCompile(columnAutoIncrementStr)

func matchColumnAutoIncrement(columnLineStr string) bool {
	return columnAutoIncrementReg.MatchString(columnLineStr)
}

const columnDefaultValueStr = "(?i)\n\\s+`\\w+`\\s+.*DEFAULT\\s+['\"]?([^'\"]+)?['\"]?.*"

var columnDefaultValueReg = regexp.MustCompile(columnDefaultValueStr)

func matchColumnDefaultValue(columnLineStr string) string {
	defaultValues := columnDefaultValueReg.FindStringSubmatch(columnLineStr)
	if len(defaultValues) > 1 {
		defaultValue := defaultValues[1]
		if defaultValue == "" {
			defaultValue = "''"
		}
		return defaultValue
	}
	return ""
}

const columnCommentStr = "(?i)\n\\s+`\\w+`\\s+.*COMMENT\\s+['\"]?([^\"']+)?['\"]?.*,?"

var columnCommentReg = regexp.MustCompile(columnCommentStr)

func matchColumnComment(columnLineStr string) string {
	comments := columnCommentReg.FindStringSubmatch(columnLineStr)
	if len(comments) > 1 {
		return comments[1]
	}
	return ""
}

const schemaName = "(?i)CREATE\\s+TABLE\\s+`?\\w?`?.?`(\\w+)`"

var (
	schemaNameReg = regexp.MustCompile(schemaName)
)

func MatchSchemaName(str string) string {
	sub := schemaNameReg.FindStringSubmatch(str)
	if len(sub) < 2 {
		panic("invalid create table schema")
	}
	return sub[1]
}

const (
	keyRegStr    = "\n\\s*KEY\\s+.*"
	keyNameStr   = "\n\\s*KEY\\s+`(\\w+)`\\(.*\\)"
	keyColumnStr = "\\((:?`\\w+`)(:?\\s*,*\\s*`\\w+`)*\\)"
)

const primaryKeyLineStr = "(?i)\n\\s*PRIMARY\\s+KEY\\s+(?:`\\w+`\\s+)?\\(`\\w+`(?:,\\s?\\w+)*\\)"

var primaryKeyLineReg = regexp.MustCompile(primaryKeyLineStr)

const keyLinePrimaryStr = "(?i)\n\\s*PRIMARY\\s+KEY\\s+\\((.*)\\)"

var keyLinePrimaryReg = regexp.MustCompile(keyLinePrimaryStr)

func matchPrimaryKeyLine(raw string) []string {
	primaryKeyLine := primaryKeyLineReg.FindString(raw)
	primaryKeys := keyLinePrimaryReg.FindStringSubmatch(primaryKeyLine)
	if len(primaryKeys) < 2 {
		panic("invalid primary keys")
	}
	primaryKeyNames := columnNameReg.FindAllStringSubmatch(primaryKeys[1], -1)

	return collector.Slice(primaryKeyNames, func(e []string) string {
		if len(e) < 2 {
			panic(fmt.Sprintf("invalid primary keys %s", e))
		}
		return e[1]
	})
}

const columnNameStr = "`(\\w+)`"

var columnNameReg = regexp.MustCompile(columnNameStr)

const uniqueKeyLineStr = "(?i)\n\\s*UNIQUE\\s+KEY\\s+(?:`\\w+`\\s+)?\\(`\\w+`(?:,\\s?\\w+)*\\)"

var uniqueKeyLineReg = regexp.MustCompile(uniqueKeyLineStr)

const keyLineUniqueKeyNameStr = "(?i)\n\\s*UNIQUE\\s+KEY\\s+`(\\w+)`\\s+\\((.*)\\)"

var keyLineUniqueKeyNameReg = regexp.MustCompile(keyLineUniqueKeyNameStr)

func matchUniqueKeyLine(raw string) map[string][]string {
	uniqueKeyLineResult := uniqueKeyLineReg.FindAllString(raw, -1)
	uniqueKeyMap := make(map[string][]string)
	for _, result := range uniqueKeyLineResult {
		result := keyLineUniqueKeyNameReg.FindStringSubmatch(result)
		if len(result) < 3 {
			panic("invalid unique key line")
		}

		indexName := result[1]
		indexes := result[2]
		indexNameResult := columnNameReg.FindAllStringSubmatch(indexes, -1)
		indexColumns := collector.Slice(indexNameResult, func(e []string) string {
			if len(e) < 2 {
				panic(fmt.Sprintf("invalid index name: %s", e))
			}

			return e[1]
		})

		uniqueKeyMap[indexName] = indexColumns
	}

	return uniqueKeyMap
}

const keyLineKeyNameStr = "(?i)\n\\s*KEY\\s+`(\\w+)`\\s+\\((.*)\\)"

var keyLineKeyNameReg = regexp.MustCompile(keyLineKeyNameStr)

func matchKeyLine(raw string) map[string][]string {
	keyLineResult := keyLineKeyNameReg.FindAllString(raw, -1)
	keyMap := make(map[string][]string)
	for _, r := range keyLineResult {
		result := keyLineKeyNameReg.FindStringSubmatch(r)
		if len(result) < 3 {
			panic(fmt.Sprintf("invalid key line: %s", r))
		}

		indexName := result[1]
		indexes := result[2]
		indexNameResult := columnNameReg.FindAllStringSubmatch(indexes, -1)
		indexColumns := collector.Slice(indexNameResult, func(e []string) string {
			if len(e) < 2 {
				panic(fmt.Sprintf("invalid index name: %s", e))
			}

			return e[1]
		})

		keyMap[indexName] = indexColumns
	}

	return keyMap
}

func matchKey(raw string, columnMap map[string]*Column) {
	columns := matchPrimaryKeyLine(raw)
	for _, column := range columns {
		c := columnMap[column]
		if c == nil {
			panic(fmt.Sprintf("primary column not found: %s", column))
		}

		c.isPrimaryKey = true
	}

	uniqueMap := matchUniqueKeyLine(raw)
	for indexName, columns := range uniqueMap {
		for _, column := range columns {
			c := columnMap[column]
			if c == nil {
				panic(fmt.Sprintf("unique column not found: %s", column))
			}

			c.uniqueIndexNames = append(c.uniqueIndexNames, indexName)
		}
	}

	indexMap := matchKeyLine(raw)
	for indexName, columns := range indexMap {
		for _, column := range columns {
			c := columnMap[column]
			if c == nil {
				panic(fmt.Sprintf("index column not found: %s", column))
			}

			c.indexNames = append(c.indexNames, indexName)
		}
	}
}
