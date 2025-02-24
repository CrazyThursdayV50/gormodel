package pkg

import (
	"os"
	"regexp"
	"strings"
	"unicode"
)

func GetRootPath() string {
	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return root
}

func Camel(val string) string {
	words := strings.FieldsFunc(val, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for i := range words {
		words[i] = strings.Title(strings.ToLower(words[i]))
	}

	return strings.Join(words, "")
}

func Snake(val string) string {
	// 正则表达式，用于匹配每个单词的起始位置
	re := regexp.MustCompile("([a-z])([A-Z])")
	// 在小写和大写字符之间加下划线，并全部转为小写
	snake := re.ReplaceAllString(val, "${1}_${2}")
	return strings.ToLower(snake)
}
