package word

import (
	"strings"
	"unicode"
)

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase 下划线转大写驼峰
// 主体逻辑是将下划线替换为空格字符，然后将其所有字符修改为其对应的首字母大写的格式，
// 最后将先前的空格字符替换为空，就可以确保各个部分所返回的首字母是大写并且是完整的一个字符串了
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 下线线转小写驼峰
// 主体逻辑可以直接复用大写驼峰的转换方法，然后只需要对其首字母进行处理就好了，在该方法中我们直接将字符串的第一位取出来，然后利用 unicode.ToLower 转换就可以了
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 驼峰转下划线
// 直接使用 go validator 库所提供的转换方法，主体逻辑为将字符串转换为小写的同时添加下划线，
// 比较特殊的一点在于，当前字符若不为小写、下划线、数字，那么进行处理的同时将对 segment 置空，保证其每一段的区间转换是正确的
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		// if is Upper append _ , all r 2 lower
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}
