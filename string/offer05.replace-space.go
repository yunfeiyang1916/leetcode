package main

import (
	"bytes"
)

// 替换空格
func replaceSpace(s string) string {
	buf := bytes.Buffer{}
	for _, v := range s {
		if v == ' ' {
			buf.Write([]byte{'%', '2', '0'})
		} else {
			buf.WriteRune(v)
		}
	}
	return buf.String()
}
