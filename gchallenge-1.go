package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func repeatStr(times int, s string, level int) string {
	result := ""
	for i := 0; i < times; i++ {
		result += s
	}
	return result
}

func getNum(i *int, s string) int {
	num := string(s[*i])
	*i++
	for *i < len(s) {
		if s[*i] == '[' {
			result, _ := strconv.Atoi(num)
			return result
		}
		num += string(s[*i])
		*i++
	}
	return 0
}

func decompressInner(i *int, s string, level *int, closeBracket *int) string {
	result := ""
	for *i < len(s) {
		matched, _ := regexp.MatchString("]", string(s[*i]))
		if matched { //state::CloseBrackets
			if *level != *closeBracket {
				*level--
				*closeBracket++
				return result
			}
			if *level == *closeBracket {
				return result
			}
		}
		matched, _ = regexp.MatchString("[a-z]", string(s[*i]))
		if matched { //state::letter
			result += string(s[*i])
		}
		matched, _ = regexp.MatchString("[0-9]", string(s[*i]))
		if matched { //state::number
			num := getNum(&*i, s)
			if s[*i] == '[' {
				*level++
				result += repeatStr(num, decompressInner(&*i, s, &*level, &*closeBracket), *level)
			}
		}
		*i++
	}
	return result
}

func decompress(s string) string {
	result := ""
	level := 0
	closeBracket := 0
	i := 0
	for i < len(s) {
		matched, _ := regexp.MatchString("[a-z]", string(s[i]))
		if matched { //state::letter
			result += string(s[i])
		}
		matched, _ = regexp.MatchString("[0-9]", string(s[i]))
		if matched { //state::number
			num := getNum(&i, s)
			level = 1
			if s[i] == '[' {
				closeBracket = 0
				result += repeatStr(num, decompressInner(&i, s, &level, &closeBracket), level)
			}
		}
		i++
	}
	return result
}

func main() {

	elements := map[string]string{
		"": "",
		//"aA4[5d]":                      "",
		//"a_4+Y[5Dd]":                   "",@TODO
		"3[abc]4[ab]c":                 "abcabcabcababababc",
		"3[ab2[x3[i]]c]4[ab]c":         "abxiiixiiicabxiiixiiicabxiiixiiicababababc",
		"tr3[u8[c]b9[q]]4[ab]c":        "truccccccccbqqqqqqqqquccccccccbqqqqqqqqquccccccccbqqqqqqqqqababababc",
		"3[ab2[x]c]4[ab]c":             "abxxcabxxcabxxcababababc",
		"tr3[u2[c]b4[q]]3[ab]c":        "truccbqqqquccbqqqquccbqqqqabababc",
		"3[ab2[w2[y]v]]z":              "abwyyvwyyvabwyyvwyyvabwyyvwyyvz",
		"xyv3[a2[x4[w]y]b5[o]c]1[asg]": "xyvaxwwwwyxwwwwybooooocaxwwwwyxwwwwybooooocaxwwwwyxwwwwybooooocasg",
	}
	for k, v := range elements {
		fmt.Printf("\ninput: %s\noutput: %s\nstrings.Compare:%d\n", k, decompress(k), strings.Compare(decompress(k), v))
	}

}
