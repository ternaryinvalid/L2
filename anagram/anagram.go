package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {
	s := []string{"пятак", "пятак", "пЯтка", "тяпка", "слиток", "листок", "хрясь", "столик"}
	lowerSlice := make([]string, 0)
	for _, v := range s {
		lowerSlice = append(lowerSlice, strings.Map(unicode.ToLower, v))
	}
	m := anagramma(lowerSlice)
	for k := range m {
		fmt.Println(k+":", m[k])
	}
}

func anagramma(s []string) map[string][]string {
	m := make(map[string][]string)

	for _, v := range s {
		c := 0
		for k := range m {
			if isAnagram(k, v) {
				if !contains(m[k], v) {
					m[k] = append(m[k], v)
					c++
				}

			}
		}
		if c == 0 {
			found := false
			for _, s := range m {
				if contains(s, v) {
					found = true
				}
			}
			if !found {
				m[v] = append(m[v], v)
			}
		}
	}

	for k := range m {
		if len(m[k]) < 2 {
			delete(m, k)
		}
	}
	return m
}

func isAnagram(s1, s2 string) bool {
	r1, r2 := []rune(s1), []rune(s2)
	sort.Slice(r1, func(i, j int) bool {
		return r1[i] < r1[j]
	})
	sort.Slice(r2, func(i, j int) bool {
		return r2[i] < r2[j]
	})

	return string(r1) == string(r2)
}

func contains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}
