package sql

import (
	"strconv"
	"strings"
)
func convertType(typ string, size string) string {
	typ = strings.ToLower(typ)
	switch {
	case strings.Contains(typ, "boolean"):
		return "bool"

	case strings.Contains(typ, "decimal"):
		sizeInt, _ := strconv.Atoi(size)
		if sizeInt > 20 {
			return "string"
		}
		return "float64"

	case strings.Contains(typ, "date"), strings.Contains(typ, "time"):
		return "time.Time"

	case strings.Contains(typ, "int"):
		return "int64"

	default:
		return "string"
	}
}
