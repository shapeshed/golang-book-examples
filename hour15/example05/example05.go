package example05

import (
	"bytes"
	"strings"
)

func StringFromAssignment(j int) string {
	var s string
	for i := 0; i < j; i++ {
		s += "a"
	}
	return s
}

func StringFromAppendJoin(j int) string {
	s := []string{}
	for i := 0; i < j; i++ {
		s = append(s, "a")
	}
	return strings.Join(s, "")
}

func StringFromBuffer(j int) string {
	var buffer bytes.Buffer
	for i := 0; i < j; i++ {
		buffer.WriteString("a")
	}
	return buffer.String()
}
