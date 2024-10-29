package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func unpack(s string) string {
	var result strings.Builder

	if unicode.IsDigit(rune(s[0])) {
		return ("некорректная строка")
	}

	if len(s) == 0 {
		return ""
	}

	for i := 0; i < len(s); i++ {

		if unicode.IsDigit(rune(s[i])) {
			count, _ := strconv.Atoi(string(s[i]))
			prevoius := s[i-1]
			result.WriteString(strings.Repeat(string(prevoius), count))
		} else {
			result.WriteByte(s[i])
		}
	}

	return result.String()

}

func main() {

	var s string
	fmt.Print("Введите строку: ")
	fmt.Scan(&s)
	res := unpack(s)
	fmt.Println(res)

}
