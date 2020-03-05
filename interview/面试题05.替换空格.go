package interview

import (
	"bytes"
	"fmt"
)

// silce
func replaceSpace1(s string) string {
	var res = []byte{}
	for _, v := range s {
		if v == ' ' {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, byte(v))
		}
	}
	return string(res)
}

// bytes.Buffer
func replaceSpace2(s string) string {
	var buf bytes.Buffer
	for _, v := range s {
		if v == ' ' {
			buf.WriteString("%20")
			continue
		}
		buf.WriteRune(v)
	}
	return buf.String()
}

func SolveOffer05() {
	s := "We are happy."
	fmt.Println(replaceSpace1(s))
}
