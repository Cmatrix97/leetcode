package interview

import (
	"fmt"
	"regexp"
	"strings"
)

// 正则表达式 A(.B?)?([Ee]C)? | .B([Ee]C)?
func isNumber1(s string) bool {
	s = strings.TrimSpace(s)
	pattern := `^(([+-]?\d+(\.(\d+)?)?([Ee]([+-]?\d+))?)|([+-]?\.(\d+)([Ee]([+-]?\d+))?))$`
	matched, _ := regexp.MatchString(pattern, s)
	return matched
}

func SolveOffer20() {
	strsTrue := []string{
		"1 ",
		"+.8",
		"+100",
		"5e2",
		"-123",
		"3.1416",
		"0123",
		".123",
		"-1E-16",
	}
	strsFalse := []string{
		"",
		".",
		"12e",
		"e9",
		"1a3.14",
		"1.2.3",
		"+-5",
		"12e+5.4",
	}
	for _, v := range strsTrue {
		fmt.Println(isNumber1(v))
	}
	fmt.Println("=====")
	for _, v := range strsFalse {
		fmt.Println(isNumber1(v))
	}
}
